package main

import (
	"fmt"
	"time"
	ws2811 "github.com/adrianh-za/ws281x-rpi"
	"github.com/adrianh-za/utils-golang/colorsys"
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
	device.SetupExit(ledChannel, ledCount)
	println(err)

	var hue = int64(0)
	var hueStep = int64(360 / ledCount)

	for {
		
		for pixelCount := 0; pixelCount < ledCount; pixelCount++ {
			var pixelHue = (hue + (hueStep * int64(pixelCount))) % 360
			var r, g, b = colorsys.Hsv2Rgb(float64(pixelHue), 1.0, 1.0)
			var hex = rgbToColor(r, g, b)
			fmt.Println("Hue: ", hue, " Hex: ", hex)
			device.Leds(ledChannel)[pixelCount] = hex
		}

		device.Wait()
		device.Render()
		time.Sleep(40 * time.Millisecond)
		hue = (hue + hueStep) % 360	//Step the HUE over one step
	}
}

func rgbToColor(r uint32, g uint32, b uint32) uint32 {
	//return ((r >> 8) & 0xff) << 16 + ((g >> 8) & 0xff) <<8 + ((b >> 8) & 0xff)
	//return ((r & 0xff) << 16) + ((g & 0xff) << 8) + (b & 0xff);
	return ((1 << 24) + (r << 16) + (g << 8) + b)
}

func rgbaToColor(r uint32, g uint32, b uint32, a uint32) uint32 {
	return ((r & 0xff) << 24) + ((g & 0xff) << 16) + ((b & 0xff) << 8) + (a & 0xff);
}
