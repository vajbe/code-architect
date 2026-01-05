package main

import "log"

func Gen() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, val := range []int{2, 3, 5} {
			ch <- val
		}
	}()
	return ch
}

func AdTen(in <-chan int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for val := range in {
			ch <- val + 10
		}
	}()
	return ch
}

func Sq(in <-chan int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for val := range in {
			ch <- val * val
		}
	}()
	return ch
}

func main() {

	for val := range Sq(AdTen(Sq(Gen()))) {
		log.Println(val)
	}

}
