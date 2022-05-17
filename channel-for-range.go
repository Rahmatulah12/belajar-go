package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	go SendMessage(messages)
	printMessage(messages)
}

func SendMessage(channel chan string) {
	for i := 0; i < 100; i++ {
		channel <- fmt.Sprintf("Mengirim data yang ke - %d", i)
	}

	close(channel) // menutup channel setelah 100 data dikirim
}

func printMessage(channel chan string) {
	for message := range channel {
		fmt.Println(message)
	}
}
