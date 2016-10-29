package main

import (
	"fmt"
	"github.com/kylelemons/gousb/usb"
	"golang.org/x/exp/io/i2c"
	"time"
)

func main() {
	// Only one context should be needed for an application.  It should always be closed.
	ctx := usb.NewContext()
	defer ctx.Close()

	powerD, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-0"}, 0x34)
	if err != nil {
		panic(err)
	}

	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v\tpower\t%s\n", now, getPowerStatus(powerD))
	}
}

func getPowerStatus(powerD *i2c.Device) (buffer []byte) {
	buffer = make([]byte, 2)
	err := powerD.ReadReg(0x00, buffer)
	if err != nil {
		panic(err)
	}

	return buffer
}
