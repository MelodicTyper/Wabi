package cmd

import (
	"runtime"
	"github.com/MelodicTyper/wabi/sysfuncs"
	"bytes"
)

var m runtime.MemStats

var fs *sysfuncs.Filesystem

func SetFS(f *sysfuncs.Filesystem) {
	fs = f // TODO rework this to be cleaner
}

func createArgs (argsArray [][]byte, args []byte) int { // modify argsArray, takes in args
	argCount := 0

	if args != nil || !bytes.Equal(args, []byte("")) {
		for argCount < len(argsArray) {
			arg, rest, found := bytes.Cut(args, []byte(" "))
			if len(arg) > 0 {
				argsArray[argCount] = arg
				argCount++
			}
			
			if !found {
				break
			}
			
		
			
			args = rest
		}
	}
	return argCount
}
