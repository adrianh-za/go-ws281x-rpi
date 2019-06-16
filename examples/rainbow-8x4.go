package main

import (
	"time"
	ws2811 "github.com/adrianh-za/go-ws281x-rpi"
	"github.com/adrianh-za/go-utils/colorsys"
	"./utils"
)

const (
	ledBrightness int = 64
	ledChannel int = 0
	ledRows int = 4
	ledCols int = 8
)

func main() {
	//Setup LED options
	opt := ws2811.DefaultOptions
	opt.Channels[ledChannel].Brightness = ledBrightness
	opt.Channels[ledChannel].LedCount = (ledRows * ledCols)
	
	//Setup LEDs
	var device *ws2811.WS2811
	device, err := ws2811.MakeWS2811(&opt)
	if (err != nil) {
		println(err)
		return
	}
	device.Init()
	device.SetupExit(ledChannel, true)

	var hue = int64(0)
	var hueStep = int64(360 / 8)

	for {
		//Loop only for the LED cols count.  The loop will set LED accross all rows
		for pixelCount := 0; pixelCount < ledCols; pixelCount++ {
			var pixelHue = (hue + (hueStep * int64(pixelCount))) % 360
			var r, g, b = colorsys.Hsv2Rgb(float64(pixelHue), 1.0, 1.0)
			var hex = colorsys.RGBToHex(r, g, b)
			utils.VerbosePrintln("Hue: ", hue, " Hex: ", hex)
			
			//Set the LED colour in each row
			for rowCount := 0; rowCount < ledRows; rowCount++ {
				device.Leds(ledChannel)[pixelCount + (ledCols * rowCount)] = hex
			}
		}
		
		//Paint the frame
		device.WaitRender()

		//Delay and increase the hue
		time.Sleep(30 * time.Millisecond)
		hue = (hue + 5) % 360     //Step the HUE over one step
	}
}