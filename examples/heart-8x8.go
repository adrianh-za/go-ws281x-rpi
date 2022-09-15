package main

import (
    "time"
    "github.com/adrianh-za/go-utils/colorsys"
    ws2811ext "github.com/adrianh-za/go-ws281x-rpi"
    ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
    ledBrightness int = 64
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
    ws2811ext.SetupExit(device, ledChannel, true)

    var redHex = colorsys.RGBToHex(255, 0, 0)
    var clearHex = colorsys.RGBToHex(0, 0, 0)

    heartMatrix := [ledRows * ledCols]int{
        0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 1, 1, 0, 0, 0,
        0, 0, 1, 1, 1, 1, 0, 0,
        0, 1, 1, 1, 1, 1, 1, 0,
        1, 1, 1, 1, 1, 1, 1, 1,
        1, 1, 1, 1, 1, 1, 1, 1,
        0, 1, 1, 0, 0, 1, 1, 0,
        0, 0, 0, 0, 0, 0, 0, 0}

    for {
        for pixelCount := 0; pixelCount < len(heartMatrix); pixelCount++ {
            if heartMatrix[pixelCount] == 1 {
                device.Leds(ledChannel)[pixelCount] = redHex
            } else {
                device.Leds(ledChannel)[pixelCount] = clearHex
            }
        }

        //Paint the frame and delay
        ws2811ext.WaitRender(device)
        time.Sleep(5000 * time.Second)
    }
}
