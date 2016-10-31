package main

import (
	"github.com/kylelemons/gousb/usb"
	"github.com/kylelemons/gousb/usbid"
	"log"
)

var (
	ctx *usb.Context
)

func init() {
	// Only one context should be needed for an application.  It should always be closed.
	ctx = usb.NewContext()
	defer ctx.Close()

	log.Println("Scanning for USB video devices")

	devs, err := ctx.ListDevices(func(desc *usb.Descriptor) bool {
		log.Printf("  Protocol: %s\n", usbid.Classify(desc))
		return false
	})

	// All Devices returned from ListDevices must be closed.
	defer func() {
		for _, d := range devs {
			d.Close()
		}
	}()

	// ListDevices can occaionally fail, so be sure to check its return value.
	if err != nil {
		log.Fatalf("list: %s", err)
	}
}
