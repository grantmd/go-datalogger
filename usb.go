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
		if desc.Class != 0xef {
			return false
		}

		log.Printf("  Protocol: %s\n", usbid.Classify(desc))
		log.Printf("  Class: %#x, SubClass: %#x, Protocol: %#x\n", desc.Class, desc.SubClass, desc.Protocol)

		for _, cfg := range desc.Configs {
			log.Printf("  %s:\n", cfg)
			for _, alt := range cfg.Interfaces {
				log.Printf("    --------------\n")
				for _, iface := range alt.Setups {
					log.Printf("    %s\n", iface)
					log.Printf("      %s\n", usbid.Classify(iface))
					for _, end := range iface.Endpoints {
						log.Printf("      %s\n", end)
					}
				}
			}
			log.Printf("    --------------\n")
		}

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
