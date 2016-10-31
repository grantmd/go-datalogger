package main

import (
	"golang.org/x/exp/io/i2c"
	"log"
)

var (
	powerD *i2c.Device
)

func init() {
	powerD, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-0"}, 0x34)
	if err != nil {
		log.Println("Error opening power I2C device:", err)
	}
	_ = powerD
}

func PowerStatus() (buffer []byte) {
	buffer = make([]byte, 2)
	err := powerD.ReadReg(0x00, buffer)
	if err != nil {
		log.Println("Error reading power I2C device:", err)
	}

	return buffer
}
