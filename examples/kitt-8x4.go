package main

import (
    "time"
    "github.com/adrianh-za/go-utils/colorsys"
    utils "github.com/adrianh-za/go-ws281x-rpi/examples/utils"
    ws2811ext "github.com/adrianh-za/go-ws281x-rpi"
    ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
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
    if (err != nil) {
        println(err)
        return
    }
    device.Init()
    ws2811ext.SetupExit(device, ledChannel, true)

    //Do some LED processing from here onwards
    var ledPositions = utils.GetLedOuterBoundry(ledRows, ledCols, false)
    var ledPositionsLength = len(ledPositions) //Just for performance, store the var
    utils.VerbosePrintln(ledPositions)

    for {
        for count := 0; count < ledPositionsLength; count++ {
            var led1 = ledPositions[(ledPositionsLength+count)%ledPositionsLength]
            var led2 = ledPositions[((ledPositionsLength+count)-1)%ledPositionsLength]
            var led3 = ledPositions[((ledPositionsLength+count)-2)%ledPositionsLength]
            var led4 = ledPositions[((ledPositionsLength+count)-3)%ledPositionsLength]
            utils.VerbosePrintln("Count: ", count, " LEDs: ", led1, led2, led3, led4)

            //Clear all LEDs
            ws2811ext.ClearAll(device, ledChannel)

            //Set the LED colors
            device.Leds(ledChannel)[led1] = colorsys.RGBToHex(255, 0, 0)
            device.Leds(ledChannel)[led2] = colorsys.RGBToHex(96, 0, 0)
            device.Leds(ledChannel)[led3] = colorsys.RGBToHex(32, 0, 0)
            device.Leds(ledChannel)[led4] = colorsys.RGBToHex(8, 0, 0)

            //Paint the LEDs
            ws2811ext.WaitRender(device)
            time.Sleep(100 * time.Millisecond)
        }
    }
}
