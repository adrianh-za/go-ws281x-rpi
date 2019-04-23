package main

import (
	"fmt"
	"time"
	ws2811 "github.com/adrianh-za/ws281x-rpi"
	"github.com/adrianh-za/utils-golang/colorsys"
)

const (
	ledBrightness int = 48
	ledChannel int = 0
	ledRows int = 4
	ledCols int = 8
)

func main() {
	opt := ws2811.DefaultOptions
	opt.Channels[ledChannel].Brightness = ledBrightness
	opt.Channels[ledChannel].LedCount = ledRows * ledCols
	
	var device *ws2811.WS2811
	device, err := ws2811.MakeWS2811(&opt)
	device.Init()
	println(err)

	var hue = int64(0)

	for {
		var r, g, b = colorsys.Hsv2Rgb(float64(hue), 1.0, 1.0)
		var hex = colorsys.RGBToHex(r, g, b)
		fmt.Println("RGB: ", r, g, b,  "Hue: ", hue, " Hex: ", hex)
		device.SetAll(ledChannel, opt.Channels[ledChannel].LedCount, hex)
		device.Render()
		
		time.Sleep(10 * time.Millisecond)
		hue = (hue + 1) % 360
	}
}