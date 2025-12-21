package cmd

import (
	"machine"

	"github.com/MelodicTyper/wabi/syscalls"
)

func HandleCmdPrompt (inputBuffer []byte) []byte {
	data, err := machine.Serial.ReadByte();
	
	if err == nil {
		if data == '\r' || data == '\n' {
			if len(inputBuffer) > 0 {
				print("\r\n")
				ProcessCmd(string(inputBuffer))
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

func ProcessCmd (cmd string) {
	switch cmd {
	case "hi":
		println("Hello!")
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
		println("Unknown command: ", cmd)
	}
}