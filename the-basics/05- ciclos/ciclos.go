package main

import "fmt"

func main() {

	// for normalito
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for while
	counter := 0
	for counter <= 10 {
		fmt.Println("ESTO ES UN WHILE, PERO A LA VEZ UN FOR EN GO :)")
		counter++
	}

	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever > 12 {
			break
		}
	}
}
