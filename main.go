package main

import (
	"time"
	"github.com/MelodicTyper/wabi/sysfuncs"
	"github.com/MelodicTyper/wabi/cmd"
	"github.com/MelodicTyper/wabi/boot"
)

func main () {
	
	
	fs := boot.InitFS()
	
	
	
	time.Sleep(2 * time.Second)
	println("System init.")
	print("\r\n> ") // Print prompt
	stdin := make([]byte, 0, 64)
	sysfuncs.InitLED()
	
	sysfuncs.TurnLEDOn()
	
	for {
		stdin = cmd.HandleCmdPrompt(stdin);
		time.Sleep(20 * time.Millisecond)
	}
	
	
}