package cmd

import (
	//"bytes"
	//"fmt"
	//"bytes"
	"machine"
	"os"
	"runtime"

	//"bytes"
	"io"
	"strings"

	"github.com/MelodicTyper/wabi/sysfuncs"
)

func HandleCmdPrompt(inputBuffer []byte) []byte {
	data, err := machine.Serial.ReadByte()

	if err == nil {
		if data == '\r' || data == '\n' {
			if len(inputBuffer) > 0 {
				print("\r\n")
				ProcessCmd(inputBuffer)
				inputBuffer = inputBuffer[:0] // reset buffer
			}
			print("\r\n> ")
			return inputBuffer
		}

		machine.Serial.WriteByte(data) // return characters to screen

		if len(inputBuffer) < cap(inputBuffer) {
			inputBuffer = append(inputBuffer, data)
		}
	}
	return inputBuffer
}

var m runtime.MemStats

var fs *sysfuncs.Filesystem

func SetFS(f *sysfuncs.Filesystem) {
	fs = f
}

func ProcessCmd(cmdBuf []byte) {

	cmd := string(cmdBuf)
	cmdFirst := strings.SplitAfterN(cmd, " ", 2)[0]
	cmdFirst = strings.Trim(cmdFirst, " ")
	
	
	//print(cmdFirst)
	switch cmdFirst {
	case "hi":
		println("Hello!")

		runtime.ReadMemStats(&m)
		// HeapSys = Bytes reserved by the runtime for the heap (Total Heap Size)
		// HeapAlloc = Bytes of allocated heap objects (Used)
		// Free = Total Heap Size - Used
		free := m.HeapSys - m.HeapAlloc
		percent := (m.HeapAlloc * 100) / m.HeapSys
		println("Total Heap:")
		println(m.HeapSys)
		println("Used Heap:")
		println(m.HeapAlloc)
		println("Free Heap:")
		println(free)
		println("---")
		println("Percent used:")
		println(percent)

	case "on":
		sysfuncs.TurnLEDOn()
	case "off":
		sysfuncs.TurnLEDOff()

	case "fs-write-file":
		// fs-write-file [filePath] ...bufToWrite
		s := strings.SplitAfterN(cmd, " ", 3)
		f := fs.OpenFile(strings.Trim("/" + s[1], " "))
		//defer f.Close()

		buf := make([]byte, 64)
		copy(buf, s[2])
		_, err := f.Write(buf)
		if err != nil {
			
			print(err)
		}
		print("Successfully written file ", s[1], " with contents ", s[2])
		erro := f.Close()
		if erro != nil {
			panic("DIDN'T SAVE" + erro.Error())
		}
	case "fs-read-file":
		// fs-read-file [filePath]
		s := strings.SplitAfterN(cmd, " ", 2)
		f, err := fs.Internal.OpenFile("/" + s[1], os.O_RDONLY)
		if err != nil {
			print("ERR on path:"+ "/" + s[1]+ err.Error())
			break;
		}
		defer f.Close()
		buf := make([]byte, 128)
		for {
			n, err := f.Read(buf)
			if err != nil {
				if err == io.EOF {
					//print("eOF")
					break
				}
				panic(err)
			}
			machine.Serial.Write(buf[:n])
		}
		//print(buf)
	case "fs-ls":
		// fs-ls [path]
		s := strings.SplitAfterN(cmd, " ", 2)
		dir,err := fs.Internal.Open(s[1])
		if err != nil {
			print("Could not open directory", s[1], err.Error())
			return
		}
		defer dir.Close()
		infos, err := dir.Readdir(0)
		_ = infos
		if err != nil {
			print("Could not read directory %s: %v\n", s[1], err)
			return
		}
		for _, info := range infos {
			s := "-rwxrwxrwx"
			if info.IsDir() {
				s = "drwxrwxrwx"
			}
			print(info.Name(), " ", info.Size(), s, "\n")
		}
	default:
		println("Unknown command: ", cmdFirst)
	}

}
