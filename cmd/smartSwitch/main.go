package main

import (
	"MagicMirrorToolSet/pin"
	"fmt"
	"time"

	"gobot.io/x/gobot"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func main() {
	rpi := raspi.NewAdaptor()

	// 56:12
	// sensor := gpio.NewPIRMotionDriver(rpi, pin.GPIO20, time.Second*30)
	sensor := gpio.NewPIRMotionDriver(rpi, pin.GPIO20)
	relay := gpio.NewGroveRelayDriver(rpi, pin.GPIO19)

	var relayBehavior = func() {
		fmt.Println("-- relay behavior --")
		relay.On()
		time.Sleep(time.Millisecond * 100)
		relay.Off()
	}

	work := func() {

		// turn off monit at start up
		go relayBehavior()

		// HC-SR501 sensor behavior
		sensor.On(gpio.MotionDetected, func(data interface{}) {
			fmt.Printf("-- %s - motion detedted: %v --\n", now(), data)
			go relayBehavior()
		})
		sensor.On(gpio.MotionStopped, func(data interface{}) {
			fmt.Printf("-- %s - motion stopped: %v --\n", now(), data)
			go relayBehavior()
		})
		sensor.On(gpio.Error, func(err interface{}) {
			fmt.Println(err)
		})
	}

	robot := gobot.NewRobot(
		"smartSwitch",
		[]gobot.Connection{rpi},
		[]gobot.Device{sensor, relay},
		work,
	)
	robot.Start()
}
