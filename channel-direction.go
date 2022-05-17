package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

/*
	Level akses channel bisa ditentukan, apakah hanya sebagai penerima, pengirim, atau penerima
	sekaligus pengirim. Konsep ini disebut dengan channel direction.

	channel chan string => Dapat mengirim dan menerima data
	channel chan<-  string => Hanya dapat mengirim data
	channel <-chan string => Hanya dapat menerima data
*/

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)

	go sendData(messages)
	receiveData(messages)
}

func sendData(channel chan<- int) { // channel hanya mengirim data
	for i := 0; true; i++ {
		channel <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Millisecond)
	}
}

// Channel Timeout
func receiveData(channel <-chan int) { // channel hanya menerima data
loop:
	for {
		select {
		case data := <-channel:
			fmt.Print("Receive data : ", data, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}
