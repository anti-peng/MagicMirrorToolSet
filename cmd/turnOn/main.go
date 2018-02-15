package main

// works with crontab
// sudo crontab -e
// - 0 0 * * * /home/pi/go/bin/main
// - 0 6 * * * /home/pi/go/bin/main
// sudo /etc/init.d/cron restart

import (
	"MagicMirrorToolSet/pin"
	"time"

	"gobot.io/x/gobot"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {

	quit := make(chan int)

	rpi := raspi.NewAdaptor()
	relay := gpio.NewGroveRelayDriver(rpi, pin.GPIO19)

	work := func() {
		relay.On()
		time.Sleep(time.Millisecond * 100)
		relay.Off()
		quit <- 0
	}

	robot := gobot.NewRobot(
		"switch",
		[]gobot.Connection{rpi},
		[]gobot.Device{relay},
		work,
	)

	robot.Start(false)
	<-quit
}
