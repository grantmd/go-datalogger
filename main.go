package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(1 * time.Second)
	for now := range c {
		//fmt.Printf("%v\tpower\t%s\n", now, PowerStatus())
		fmt.Printf("%v\n", now)
	}
}
