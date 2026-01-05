package main

import (
	"log"
)

func ChannelExample() {

	ch := make(chan bool)
	ch <- true
	log.Println("Exitting...")
}
