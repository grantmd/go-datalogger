package main

import (
	"fmt"
	"github.com/kylelemons/gousb/usb"
	"time"
)

func main() {
	// Only one context should be needed for an application.  It should always be closed.
	ctx := usb.NewContext()
	defer ctx.Close()

	c := time.Tick(1 * time.Second)
	for now := range c {
		//fmt.Printf("%v\tpower\t%s\n", now, PowerStatus())
	}
}
