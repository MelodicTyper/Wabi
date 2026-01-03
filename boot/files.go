package boot

import (
	"machine"
	"tinygo.org/x/tinyfs/littlefs"
)

var (
	blockdev = machine.Flash
	fs = littlefs.New(blockdev)
)

func InitFS () littlefs.LFS {
	fs.Configure(&littlefs.Config{
		CacheSize:     512,
		LookaheadSize: 512,
		BlockCycles:   100,
	})
	println(machine.FlashDataStart())
	println(machine.FlashDataEnd())
	
	if err := fs.Mount(); err != nil {
		println("Could not mount LittleFS filesystem: " + err.Error() + "\r\n")
		if err := fs.Format(); err != nil {
			println("Could not format LittleFS filesystem: " + err.Error() + "\r\n")
		} else {
			println("Successfully formatted LittleFS filesystem.\r\n")
		}
		if err := fs.Mount(); err != nil {
			panic("Mount failed after format")
		}
	} else {
		println("Successfully mounted LittleFS filesystem.\r\n")
	}
	println("FS Initialization compelte")
	//println(fs.Size())
	
}