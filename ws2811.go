    
// Copyright 2018 Jacques Supcik / HEIA-FR
// Copyright 2019 Adrian Houghton
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Interface to ws2811 chip (neopixel driver). Make sure that you have
// ws2811.h and pwm.h in a GCC include path (e.g. /usr/local/include) and
// libws2811.a in a GCC library path (e.g. /usr/local/lib).
// See https://github.com/jgarff/rpi_ws281x for instructions

package ws2811

import "github.com/pkg/errors"

const (
	// DefaultDmaNum is the default DMA number. Usually, this is 5 ob the Raspberry Pi (NOW CHANGE TO 10)
	DefaultDmaNum = 10
	// RpiPwmChannels is the number of PWM leds in the Raspberry Pi
	RpiPwmChannels = 2
	// TargetFreq is the target frequency. It is usually 800kHz (800000), and an go as low as 400000
	TargetFreq = 800000
	// DefaultGpioPin is the default pin on the Raspberry Pi where the signal will be available. Note
	// that it is the BCM (Broadcom Pin Number) and the "Pin" 18 is actually the physical pin 12 of the
	// Raspberry Pi.
	DefaultGpioPin = 18
	// DefaultLedCount is the default number of LEDs on the stripe.
	DefaultLedCount = 32
	// DefaultBrightness is the default maximum brightness of the LEDs. The brightness value can be between 0 and 255.
	// If the brightness is too low, the LEDs remain dark. If the brightness is too high, the system needs too much
	// current.
	DefaultBrightness = 96 // Safe value between 0 and 255.
)

const (
	// HwVerTypeUnknown represents unknown hardware
	HwVerTypeUnknown = 0
	// HwVerTypePi1 represents the Raspberry Pi 1
	HwVerTypePi1 = 1
	// HwVerTypePi2 represents the Raspberry Pi 2
	HwVerTypePi2 = 2
)

// StateDesc is a map from a return state to its string description.
var StateDesc = map[int]string{
	0:   "Success",
	-1:  "Generic failure",
	-2:  "Out of memory",
	-3:  "Hardware revision is not supported",
	-4:  "Memory lock failed",
	-5:  "mmap() failed",
	-6:  "Unable to map registers into userspace",
	-7:  "Unable to initialize GPIO",
	-8:  "Unable to initialize PWM",
	-9:  "Failed to create mailbox device",
	-10: "DMA error",
	-11: "Selected GPIO not possible",
	-12: "Unable to initialize PCM",
	-13: "Unable to initialize SPI",
	-14: "SPI transfer error",
}

// HwDesc is the Hardware Description
type HwDesc struct {
	Type          uint32
	Version       uint32
	PeriphBase    uint32
	VideocoreBase uint32
	Desc          string
}

// ChannelOption is the list of channel options
type ChannelOption struct {
	// GpioPin is the GPIO Pin with PWM alternate function, 0 if unused
	GpioPin int
	// Invert inverts output signal
	Invert bool
	// LedCount is the number of LEDs, 0 if channel is unused
	LedCount int
	// StripeType is the strip color layout -- one of WS2811StripXXX constants
	StripeType int
	// Brightness is the maximum brightness of the LEDs. Value between 0 and 255
	Brightness int
	// WShift is the white shift value
	WShift int
	// RShift is the red shift value
	RShift int
	// GShift is the green shift value
	GShift int
	// BShift is blue shift value
	BShift int
	// Gamma is the gamma correction table
	Gamma []byte
	// Must capture Interrupt and SIGTERM signals to handle program exit
	CaptureExit bool
	// Must the LEDs be cleared on exit
	ClearOnExit bool
}

// Option is the list of device options
type Option struct {
	// RenderWaitTime is the time in µs before the next render can run
	RenderWaitTime int
	// Frequency is the required output frequency
	Frequency int
	// DmaNum is the number of a DMA _not_ already in use
	DmaNum int
	// Channels are channel options
	Channels []ChannelOption
}

// DefaultOptions defines sensible default options for MakeWS2811
var DefaultOptions = Option{
	Frequency: TargetFreq,
	DmaNum:    DefaultDmaNum,
	Channels: []ChannelOption{
		{
			GpioPin:     DefaultGpioPin,
			LedCount:    DefaultLedCount,
			Brightness:  DefaultBrightness,
			StripeType:  WS2812Strip,
			Invert:      false,
			Gamma:       nil,
			CaptureExit: true,
			ClearOnExit: true,
		},
	},
}

// Leds returns the LEDs array of a given channel
func (ws2811 *WS2811) Leds(channel int) []uint32 {
	return ws2811.leds[channel]
}

// SetLedsSync wait for the frame to finish and replace all the LEDs
func (ws2811 *WS2811) SetLedsSync(channel int, leds []uint32) error {
	if err := ws2811.Wait(); err != nil {
		return errors.WithMessage(err, "Error setting LEDs")
	}
	l := len(leds)
	if l > len(ws2811.leds[channel]) {
		return errors.New("Error: Too many LEDs")
	}
	for i := 0; i < l; i++ {
		ws2811.leds[channel][i] = leds[i]
	}
	return nil
}

// StatusDesc returns the description of a status code
func StatusDesc(code int) string {
	desc, ok := StateDesc[code]
	if ok {
		return desc
	}
	return "Unknown"
}

var gamma8 = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5,
	5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10,
	10, 10, 11, 11, 11, 12, 12, 13, 13, 13, 14, 14, 15, 15, 16, 16,
	17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 24, 24, 25,
	25, 26, 27, 27, 28, 29, 29, 30, 31, 32, 32, 33, 34, 35, 35, 36,
	37, 38, 39, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 50,
	51, 52, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 66, 67, 68,
	69, 70, 72, 73, 74, 75, 77, 78, 79, 81, 82, 83, 85, 86, 87, 89,
	90, 92, 93, 95, 96, 98, 99, 101, 102, 104, 105, 107, 109, 110, 112, 114,
	115, 117, 119, 120, 122, 124, 126, 127, 129, 131, 133, 135, 137, 138, 140, 142,
	144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 167, 169, 171, 173, 175,
	177, 180, 182, 184, 186, 189, 191, 193, 196, 198, 200, 203, 205, 208, 210, 213,
	215, 218, 220, 223, 225, 228, 231, 233, 236, 239, 241, 244, 247, 249, 252, 255,
}

// 4 color R, G, B and W ordering
const (
	// SK6812StripRGBW is the RGBW Mode
	SK6812StripRGBW = 0x18100800
	// SK6812StripRBGW is the StripRBGW Mode
	SK6812StripRBGW = 0x18100008
	// SK6812StripGRBW is the StripGRBW Mode
	SK6812StripGRBW = 0x18081000
	// SK6812StrioGBRW is the StrioGBRW Mode
	SK6812StrioGBRW = 0x18080010
	// SK6812StrioBRGW is the StrioBRGW Mode
	SK6812StrioBRGW = 0x18001008
	// SK6812StripBGRW is the StripBGRW Mode
	SK6812StripBGRW = 0x18000810
	// SK6812ShiftWMask is the Shift White Mask
	SK6812ShiftWMask = 0xf0000000
)

// 3 color R, G and B ordering
const (
	// WS2811StripRGB is the RGB Mode
	WS2811StripRGB = 0x100800
	// WS2811StripRBG is the RBG Mode
	WS2811StripRBG = 0x100008
	// WS2811StripGRB is the GRB Mode
	WS2811StripGRB = 0x081000
	// WS2811StripGBR is the GBR Mode
	WS2811StripGBR = 0x080010
	// WS2811StripBRG is the BRG Mode
	WS2811StripBRG = 0x001008
	// WS2811StripBGR is the BGR Mode
	WS2811StripBGR = 0x000810
)

// Predefined fixed LED types
const (
	// WS2812Strip is the WS2812 Mode
	WS2812Strip = WS2811StripGRB
	// SK6812Strip is the SK6812 Mode
	SK6812Strip = WS2811StripGRB
	// SK6812WStrip is the SK6812W Mode
	SK6812WStrip = SK6812StripGRBW
)
