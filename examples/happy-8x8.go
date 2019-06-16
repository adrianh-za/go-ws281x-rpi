package main

import (
	"time"

	"github.com/adrianh-za/go-utils/colorsys"
	ws2811 "github.com/adrianh-za/go-ws281x"
)

const (
	ledBrightness int = 48
	ledChannel    int = 0
	ledRows       int = 8
	ledCols       int = 8
)

func main() {
	//Setup LED options
	opt := ws2811.DefaultOptions
	opt.Channels[ledChannel].Brightness = ledBrightness
	opt.Channels[ledChannel].LedCount = (ledRows * ledCols)

	//Setup LEDs
	var device *ws2811.WS2811
	device, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		println(err)
		return
	}
	device.Init()
	device.SetupExit(ledChannel, true)

	//Set the different color pixels
	var yellowHex = colorsys.RGBToHex(255, 216, 0)
	var redHex = colorsys.RGBToHex(255, 0, 0)
	var blueHex = colorsys.RGBToHex(0, 0, 255)
	var clearHex = colorsys.RGBToHex(0, 0, 0)

	// Face matrix
	faceMatrix := [ledRows * ledCols]int{
		0, 0, 1, 1, 1, 1, 0, 0,
		0, 1, 1, 1, 1, 1, 1, 0,
		1, 1, 1, 2, 2, 1, 1, 1,
		1, 1, 2, 1, 1, 2, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 1, 3, 1, 1,
		0, 1, 1, 1, 1, 1, 1, 0,
		0, 0, 1, 1, 1, 1, 0, 0}

	//Do some painting
	for {
		for pixelCount := 0; pixelCount < len(faceMatrix); pixelCount++ {
			if faceMatrix[pixelCount] == 1 {
				device.Leds(ledChannel)[pixelCount] = yellowHex
			} else if faceMatrix[pixelCount] == 2 {
				device.Leds(ledChannel)[pixelCount] = redHex
			} else if faceMatrix[pixelCount] == 3 {
				device.Leds(ledChannel)[pixelCount] = blueHex
			} else {
				device.Leds(ledChannel)[pixelCount] = clearHex
			}
		}

		//Paint the frame and delay
		device.WaitRender()
		time.Sleep(5000 * time.Second)
	}
}
