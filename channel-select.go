package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Numbers:", numbers)

	var channel1 = make(chan float64)
	go getAvarage(numbers, channel1)

	var channel2 = make(chan float64)
	go getMax(numbers, channel2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-channel1: // menerima data
			fmt.Printf("Average \t: %.2f \n", avg)
		case max := <-channel2:
			fmt.Printf("Max \t: %.2f \n", max)
		}
	}
}

func getAvarage(numbers []float64, channel chan float64) {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	channel <- sum / float64(len(numbers))
}

func getMax(numbers []float64, channel chan float64) {
	var max = numbers[0]

	for _, value := range numbers {
		if max < value {
			max = value
		}
	}
	channel <- max
}
