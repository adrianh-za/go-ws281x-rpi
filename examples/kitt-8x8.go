package main

import (
	"fmt"
	"time"

	"github.com/adrianh-za/utils-golang/colorsys"
	ws2811 "github.com/adrianh-za/ws281x-rpi"
)

const (
	ledBrightness int = 128
	ledChannel    int = 0
	ledRows       int = 4
	ledCols       int = 8
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

	//Do some LED processing from here onwards
	var ledPositions = GetLedOuterBoundry(ledRows, ledCols, false)
	var ledPositionsLength = len(ledPositions) //Just for performance, store the var
	fmt.Println(ledPositions)

	for {
		for count := 0; count < ledPositionsLength; count++ {
			var led1 = ledPositions[(ledPositionsLength+count)%ledPositionsLength]
			var led2 = ledPositions[((ledPositionsLength+count)-1)%ledPositionsLength]
			var led3 = ledPositions[((ledPositionsLength+count)-2)%ledPositionsLength]
			var led4 = ledPositions[((ledPositionsLength+count)-3)%ledPositionsLength]
			fmt.Println("Count: ", count, " LEDs: ", led1, led2, led3, led4)

			//Clear all LEDs
			device.ClearAll(ledChannel, opt.Channels[ledChannel].LedCount)

			//Set the LED colors
			device.Leds(ledChannel)[led1] = colorsys.RGBToHex(255, 0, 0)
			device.Leds(ledChannel)[led2] = colorsys.RGBToHex(96, 0, 0)
			device.Leds(ledChannel)[led3] = colorsys.RGBToHex(32, 0, 0)
			device.Leds(ledChannel)[led4] = colorsys.RGBToHex(8, 0, 0)

			//Paint the LEDs
			device.Render()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func GetLedOuterBoundry(rows int, columns int, reverse bool) []int {
	var totalLEDs = rows * columns
	var ledIndexes = make([]int, 0)

	//First row
	for count := 0; count < columns; count++ {
		ledIndexes = append(ledIndexes, count)
	}
	//First column
	for count := 2; count < rows; count++ {
		ledIndexes = append(ledIndexes, (count*columns)-1)
	}
	//Second row (opposite direction)
	for count := totalLEDs; count > (totalLEDs - columns); count-- {
		ledIndexes = append(ledIndexes, count-1)
	}
	//Second column (opposite direction)
	for count := totalLEDs - (columns * 2); count > 0; count = count - columns {
		ledIndexes = append(ledIndexes, count)
	}

	//Return the array of INT, reverse is specified
	if !reverse {
		return ledIndexes
	} else {
		for i, j := 0, len(ledIndexes)-1; i < j; i, j = i+1, j-1 {
			ledIndexes[i], ledIndexes[j] = ledIndexes[j], ledIndexes[i]
		}
		return ledIndexes
	}
}