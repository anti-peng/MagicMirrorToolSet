package main

import (
	"MagicMirrorToolSet/pin"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, pin.GPIO26)

	worker := func() {
		gobot.Every(time.Second*1, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot(
		"blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		worker,
	)
	robot.Start()
}
