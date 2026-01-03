package main

import (
	"fmt"
	"log"
)

func PanicExample() (result int) {

	defer func() {
		fmt.Println("You're returning ", result)
		r := recover()
		if r != nil {
			log.Println("Catching panic-2 error; Err:", r)
		}
		result += 10
	}()

	defer func() {
		r := recover()
		if r != nil {
			log.Println("Catching panic-1 error; Err:", r)
		}
	}()

	num := 10

	if num == 10 {
		panic("Jo")
	}

	return 10
}
