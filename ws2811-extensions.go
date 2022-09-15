// This file extends the WS2811 package of github.com/rpi-ws281x/rpi-ws281x-go with
// the below "static" methods.

package ws2811ext

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/pkg/errors"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)


func SetupExit(ws2811 *ws2811.WS2811, channel int, clear bool) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		for range signalChan {
			if (clear) {
				ClearAll(ws2811, channel)
				ws2811.Render()
			}
			ws2811.Fini()
			os.Exit(1)
		}
	}()
}

// SetAll sets all the leds for matrix to the specified color
func SetAll(ws2811 *ws2811.WS2811, channel int, color uint32) {
	for led := 0; led < len(ws2811.Leds(channel)); led++ {	
		ws2811.Leds(channel)[led] = color
	}
}

// ClearAll clears all the leds (sets to 0x0000000) for matrix
func ClearAll(ws2811 *ws2811.WS2811, channel int) {
	SetAll(ws2811, channel, 0)
}

// WaitRender first waits and then renders
func WaitRender(ws2811 *ws2811.WS2811) (error) {
	if err := ws2811.Wait(); err != nil {
		return errors.WithMessage(err, "Error waiting LEDs")
	}
	if err := ws2811.Render(); err != nil {
		return errors.WithMessage(err, "Error rendering LEDs")
	}
	return nil
}

/*
// SetupExit captures Interrupt and SIGTERM signals to handle program exit 
func (ws2811 *ws2811.WS2811) SetupExit(channel int, clear bool) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		for range signalChan {
			if (clear) {
				ws2811.ClearAll(channel)
				ws2811.Render()
			}
			ws2811.Fini()
			os.Exit(1)
		}
	}()
}

// SetAll sets all the leds for matrix to the specified color
func (ws2811 *ws2811.WS2811) SetAll(channel int, color uint32) {
	for led := 0; led < len(ws2811.Leds(channel)); led++ {	
		ws2811.Leds(channel)[led] = color
	}
}

// ClearAll clears all the leds (sets to 0x0000000) for matrix
func (ws2811 *ws2811.WS2811) ClearAll(channel int) {
	ws2811.SetAll(channel, 0)
}

// WaitRender first waits and then renders
func (ws2811 *ws2811.WS2811) WaitRender() (error) {
	if err := ws2811.Wait(); err != nil {
		return errors.WithMessage(err, "Error waiting LEDs")
	}
	if err := ws2811.Render(); err != nil {
		return errors.WithMessage(err, "Error rendering LEDs")
	}
	return nil
}*/

