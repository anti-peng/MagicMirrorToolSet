package main

import (
	"MagicMirrorToolSet/pin"
	"fmt"
	"time"

	"gobot.io/x/gobot"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	rpi := raspi.NewAdaptor()

	sensor := gpio.NewPIRMotionDriver(rpi, pin.GPIO25, time.Second*10)
	led := gpio.NewLedDriver(rpi, pin.GPIO26)
	relay := gpio.NewGroveRelayDriver(rpi, pin.GPIO24)

	var relayBehavior = func() {
		relay.On()
		time.Sleep(time.Second * 1)
		relay.Off()
	}

	work := func() {
		sensor.On(gpio.MotionDetected, func(data interface{}) {
			fmt.Printf("-- motion detedted: %v\n --", data)
			led.On()
		})

		sensor.On(gpio.MotionStopped, func(data interface{}) {
			fmt.Printf("-- motion stopped: %v\n --", data)
			led.Off()
			// turn off some driver after 'MotionStopped'
			go relayBehavior()
		})
	}

	robot := gobot.NewRobot(
		"smartSwitch",
		[]gobot.Connection{rpi},
		[]gobot.Device{sensor, led},
		work,
	)
	robot.Start()
}
