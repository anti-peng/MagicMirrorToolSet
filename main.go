//usr/bin/env go run $0 "$@"; exit
package main

import (
	"fmt"
	"os"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	pin := rpio.Pin(GPIO21)
	pin.Output()

	for i := 0; i < 20; i++ {
		pin.Toggle()
		time.Sleep(time.Second / 2)
	}
}
