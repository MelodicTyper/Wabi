package boot

import "github.com/soypat/cyw43439"

func InitCYW43439() *cyw43439.Device {
	dev := cyw43439.NewPicoWDevice()
	cfg := cyw43439.DefaultWifiConfig()
	
	err := dev.Init(cfg)

	if err != nil {
		panic(err)
	}
	return dev
}