package main

import (
	"fmt"
	"time"
)

func main() {
	from := 0
	to := 50000000
	sum := 0
	start := time.Now()
	sum += getSummen(from, to)
	timeS := time.Since(start)

	c := make(chan int)
	start = time.Now()
	//Hvorfor ikke prøve med litt multithread?
	go getSum(from, to/2, c)
	go getSum(to/2, to, c)
	x, y := <-c, <-c
	timeP := time.Since(start)

	fmt.Println("Go Routines:")
	fmt.Println(timeP)
	fmt.Println(x + y)
	fmt.Println("Typisk:")
	fmt.Println(timeS)
	fmt.Println(sum)
}

//gogetSum channels
func getSum(from, to int, c chan int) {
	sum := 0
	for i := from; i != to; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	c <- sum
}

//Typisk måte
func getSummen(from, to int) (sum int) {
	sum = 0
	for i := from; i != to; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	return
}
