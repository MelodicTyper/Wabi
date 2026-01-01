package cmd

import (
	//"bytes"
	//"fmt"
	"machine"
	"runtime"
	//"bytes"
	"github.com/MelodicTyper/wabi/syscalls"
)

func HandleCmdPrompt (inputBuffer []byte) []byte {
	data, err := machine.Serial.ReadByte();
	
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
	return inputBuffer;
}

var m runtime.MemStats

func ProcessCmd (cmd []byte) {
	

	
	switch string(cmd) {
	case "hi":
		println("Hello!")
		
		
    	runtime.ReadMemStats(&m)
	    // HeapSys = Bytes reserved by the runtime for the heap (Total Heap Size)
	    // HeapAlloc = Bytes of allocated heap objects (Used)
	    // Free = Total Heap Size - Used
	    free := m.HeapSys - m.HeapAlloc
		percent := (m.HeapAlloc *100 )/ m.HeapSys
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
		syscalls.TurnLEDOn()
	case "off":
		syscalls.TurnLEDOff()
	case "fs-init":
		syscalls.InitFS()
	case "fs-test-file-create":
		syscalls.WriteTestFile()
	case "fs-test-file-read":	
		syscalls.ReadTestFile()
	default:
		println("Unknown command: ")
	}
	
}