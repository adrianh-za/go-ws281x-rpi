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

    //Setup the vars
    var hue = int64(0)
    var hueStep = int64(360 / opt.Channels[ledChannel].LedCount)

    for {
        //Set each pixel to correct colour
        for pixelCount := 0; pixelCount < opt.Channels[ledChannel].LedCount; pixelCount++ {
            var pixelHue = (hue + (hueStep * int64(pixelCount))) % 360
            var r, g, b = colorsys.Hsv2Rgb(float64(pixelHue), 1.0, 1.0)
            var hex = colorsys.RGBToHex(r, g, b)
            utils.VerbosePrintln("Hue: ", hue, " Hex: ", hex)
            device.Leds(ledChannel)[pixelCount] = hex
        }

        //Paint the LEDs
        ws2811ext.WaitRender(device)

        //Sleep a bit and increment hue
        time.Sleep(40 * time.Millisecond)
        hue = (hue + hueStep) % 360    //Step the HUE over one step
    }
}