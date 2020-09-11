package main

import (
	"github.com/Electro3/setinterval"
	"log"
)

func main() {
	log.Println("linea 1")
	line := 0
	interval := setinterval.Start(func() {
		log.Println("linea 2")
		line++
	}, 1000, true)
	
	for {
		if line == 10 {
			interval <- true
			break
		}
	}
	log.Println("linea 3")
}
