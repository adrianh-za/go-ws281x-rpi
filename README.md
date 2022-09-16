# Go-WS281x-RPi

Library written in Go to allow controlling of the WS281x LED strips on a Raspberry Pi.

## Usage ##

This needs quite a few steps, so try keep up!

First, grab the C library

1.  `sudo apt-get update`  (update RPi's software)

2.  `sudo apt-get install scons gcc git`  (get needed packages)

3.  We need to disable the audio, due to PCM

    `sudo nano /etc/modprobe.d/snd-blacklist.conf`

    and add this into file

    `blacklist snd_bcm2835`

4.  We need to edit configuration of RPi

    `sudo nano /boot/config.txt`

    and then add a `#` in front of

    `dtparam=audio=on`

    to comment it out.

    For headless, you may also need to add the following two line to config.txt

    `hdmi_force_hotplug=1`
    `hdmi_force_edid_audio=1`

5.  `sudo reboot`

6. `go get github.com/rpi-ws281x/rpi-ws281x-go` so that the `go/src/github.com` directory is created.

    => An error will likely occur if `https://github.com/jgarff/rpi_ws281x` is not installed.

7.  browse to `go/src/github.com` and then create the `jgarff` directory with  `mkdir jgarff` if it doesn't exist.  

8.  browse to `go/src/github.com/jgarff` and type  `git clone https://github.com/jgarff/rpi_ws281x`  (the C library that drives the LEDs)

9.  browse to  `go/src/github.com/jgarff/rpi_ws281x` github repo and type `scons` to compile library

10. copy *.h files to /usr/local/include

    `sudo cp *.h /usr/local/include`

11. copy *.a files to /usr/local/lib

    `sudo cp *.a /usr/local/lib`
    
12. `go get github.com/rpi-ws281x/rpi-ws281x-go`
    `go get github.com/adrianh-za/go-ws281x-rpi`
    `go get github.com/adrianh-za/go-utils/colorsys`

13. browse to $/go/src/github.com/adrianh-za/go-ws281x-rpi/examples

14. `sudo -E go run [filename].go`  (check filenames below)

15. ctrl-c to quit

Some extra help

<b><a href="https://tutorials-raspberrypi.com/connect-control-raspberry-pi-ws2812-rgb-led-strips/" target="_blank">Connect and Control WS2812 RGB LED Strips via Raspberry Pi</a></b>

Example filenames

** The filename contains the LED matrix size

* clear-8x4.go
* clear-8x8.go
* colours-8x4.go
* colours-8x8.go
* happy-8x8.go
* heart-8x8.go
* hw.go
* kitt-8x4.go
* kitt-8x8.go
* rainbow-8x4.go
* rainbow-8x8.go
* rainbow-snake-8x4.go
* rainbow-snake-8x8.go

## Acknowledgements ##

Utilises the WS281x library written by <b><a href="https://github.com/jgarff" target="_blank">Jeremy Garff</a></b>

Borrows heavily from the work done by <b><a href="https://github.com/supcik" target="_blank">Jacques Supcik</a></b>

## Gits ##

https://github.com/jgarff/rpi_ws281x

https://github.com/rpi-ws281x/rpi-ws281x-go
