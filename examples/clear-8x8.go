package main

import (
    utils "github.com/adrianh-za/go-ws281x-rpi/examples/utils"
    ws2811ext "github.com/adrianh-za/go-ws281x-rpi"
    ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
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
    utils.VerbosePrintln("LEDs initialized")

    //Clear the LED hat
    ws2811ext.ClearAll(device, ledChannel)
    ws2811ext.WaitRender(device)
    utils.VerbosePrintln("LEDs cleared")
}