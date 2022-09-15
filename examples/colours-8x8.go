package main

import (
    "time"
    "github.com/adrianh-za/go-utils/colorsys"
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
    ws2811ext.SetupExit(device, ledChannel, true)

    //Do some LED processing from here onwards
    var hue = int64(0)
    for {
        //Calculate RGB and convert to HEX
        var r, g, b = colorsys.Hsv2Rgb(float64(hue), 1.0, 1.0)
        var hex = colorsys.RGBToHex(r, g, b)
        utils.VerbosePrintln("RGB: ", r, g, b,  "Hue: ", hue, " Hex: ", hex)

        //Set LEDs
        ws2811ext.SetAll(device, ledChannel, hex)
        ws2811ext.WaitRender(device)

        //Sleep and increment the HUE
        time.Sleep(20 * time.Millisecond)
        hue = (hue + 1) % 360
    }
}