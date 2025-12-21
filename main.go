package main

import (
	"time"
	"github.com/MelodicTyper/wabi/syscalls"
	"github.com/MelodicTyper/wabi/cmd"
)

func main () {
	time.Sleep(2 * time.Second)
	println("System init.")
	print("\r\n> ") // Print prompt
	stdin := make([]byte, 0, 128)
	syscalls.InitLED()
	
	syscalls.TurnLEDOn()
	
	for {
		stdin = cmd.HandleCmdPrompt(stdin);
		time.Sleep(20 * time.Millisecond)
	}
}