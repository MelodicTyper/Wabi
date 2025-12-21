package syscalls

// Developement syscalls, not meant to be anything close to the final version

import (
	"machine"
	"os"
	"io"
	//"tinygo.org/x/drivers/flash"
	//"tinygo.org/x/tinyfs"
	"tinygo.org/x/tinyfs/littlefs"
	"github.com/soypat/cyw43439"
)


var ledDevice *cyw43439.Device

func InitLED () {
	dev := cyw43439.NewPicoWDevice()
	cfg := cyw43439.DefaultWifiConfig()
	
	err := dev.Init(cfg)

	if err != nil {
		panic(err)
	}
	ledDevice = dev
}

func TurnLEDOn () {
	ledDevice.GPIOSet(0, true)
}

func TurnLEDOff () {
	ledDevice.GPIOSet(0, false)
}



var (
	blockdev = machine.Flash
	fs = littlefs.New(blockdev)
)


func InitFS () {
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
	println(fs.Size())
	
}

func WriteTestFile () {
	f, err := fs.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC);
	defer f.Close()
	
	buf := make([]byte, 64)
	content := "Hi!"
	copy(buf, content)
	f.Write(buf)
	if err != nil {
		println("error opening %s: %s\r\n", "test.txt", err.Error())
		return
	}
	
}

func ReadTestFile () {
	f, err := fs.Open("test.txt")
	if err != nil {
		println("Could not open: " + err.Error())
		return
	}
	defer f.Close();
	buf := make([]byte, 128)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			println("Error Reading File")
		}
		println(string(buf[:n]))
	}
}