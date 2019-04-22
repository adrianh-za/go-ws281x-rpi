package main

import (
	ws2811 "github.com/adrianh-za/ws281x-rpi"
)

const (
	ledCount int = 32
	ledBrightness int = 96
	ledChannel int = 0
)

func main() {
	opt := ws2811.DefaultOptions
	opt.Channels[ledChannel].Brightness = ledBrightness
	opt.Channels[ledChannel].LedCount = ledCount
	
	var device *ws2811.WS2811
	device, err := ws2811.MakeWS2811(&opt)
	device.Init()
	defer device.Fini()
	println(err)
	device.Clear(ledChannel, ledCount)
}