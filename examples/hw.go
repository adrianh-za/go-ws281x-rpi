package main

import (
    "fmt"
    ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func main() {
    fmt.Println("*****************************")
    fmt.Println("* rpi_ws281x Hardware Check *")
    fmt.Println("*****************************")
    hw := ws2811.HwDetect()
    fmt.Printf("Hardware Type    : %d\n", hw.Type)
    fmt.Printf("Hardware Version : 0x%08X\n", hw.Version)
    fmt.Printf("Periph base      : 0x%08X\n", hw.PeriphBase)
    fmt.Printf("Video core base  : 0x%08X\n", hw.VideocoreBase)
    fmt.Printf("Description      : %v\n", hw.Desc)
}