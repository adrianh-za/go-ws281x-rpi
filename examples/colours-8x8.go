package main

import (
	"time"
	ws2811 "github.com/adrianh-za/go-ws281x"
	"github.com/adrianh-za/go-utils/colorsys"
	"./utils"
)

const (
	ledBrightness int = 64
	ledChannel int = 0
	ledRows int = 8
	ledCols int = 8
)

func main() {
	//Setup LED options
	opt := ws2811.DefaultOptions
	opt.Channels[ledChannel].Brightness = ledBrightness
	opt.Channels[ledChannel].LedCount = ledRows * ledCols
	
	//Setup LEDs
	var device *ws2811.WS2811
	device, err := ws2811.MakeWS2811(&opt)
	if (err != nil) {
		println(err)
		return
	}
	device.Init()
	device.SetupExit(ledChannel, true)
	
	//Do some LED processing from here onwards
	var hue = int64(0)
	for {
		//Calculate RGB and convert to HEX
		var r, g, b = colorsys.Hsv2Rgb(float64(hue), 1.0, 1.0)
		var hex = colorsys.RGBToHex(r, g, b)
		utils.VerbosePrintln("RGB: ", r, g, b,  "Hue: ", hue, " Hex: ", hex)
		
		//Set LEDs
		device.SetAll(ledChannel, hex)
		device.WaitRender()
		
		//Sleep and increment the HUE
		time.Sleep(20 * time.Millisecond)
		hue = (hue + 1) % 360
	}
}