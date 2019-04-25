package main

import (
	ws2811 "github.com/adrianh-za/ws281x-rpi"
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
	device.Init()
	println(err)

	//Clear the LED hat
	device.ClearAll(ledChannel, opt.Channels[ledChannel].LedCount)
	device.Wait()
}