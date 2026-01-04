package main

import (
	"time"
	"github.com/MelodicTyper/wabi/sysfuncs"
	"github.com/MelodicTyper/wabi/cmd"
	"github.com/MelodicTyper/wabi/boot"
)

func main () {
	
	
	littlefs := boot.InitFS()
	fs := &sysfuncs.Filesystem{Internal: littlefs}
	
	//cyw43439 := boot.InitCYW43439()
	
	
	
	time.Sleep(2 * time.Second)
	println("System init.")
	print("\r\n> ") // Print prompt
	stdin := make([]byte, 0, 64)
	sysfuncs.InitLED()
	
	sysfuncs.TurnLEDOn()
	
	print(fs.Size())
	cmd.SetFS(fs)
	for {
		stdin = cmd.HandleCmdPrompt(stdin);
		time.Sleep(20 * time.Millisecond)
	}
	
	
}