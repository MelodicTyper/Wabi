package syscalls

import (
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