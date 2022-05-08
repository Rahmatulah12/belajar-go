package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		Channel digunakan untuk menghubungkan gorutine satu dengan goroutine lainnya.
		Mekanismenya adalah dengan cara serah-terima data lewat channel tersebut. Goroutine
		pengirim dan penerima harus berada pada channel yang berbeda (konsep ini disebut
		buffered channel). Pengiriman dan penerimaan data pada channel bersifat blocking atau
		synchronous.
	*/
	runtime.GOMAXPROCS(2)

	// Buat channel
	var messages = make(chan string)

	var sayHello = func(name string) {
		var data = fmt.Sprintf("Hello %s", name)
		messages <- data // kirim data ke channel messages
	}

	go sayHello("John")
	go sayHello("Jane")
	go sayHello("Jack")

	var message1 = <-messages // menerima data dari channel messages
	fmt.Println(message1)

	var message2 = <-messages // menerima data dari channel messages
	fmt.Println(message2)

	var message3 = <-messages // menerima data dari channel messages
	fmt.Println(message3)

}
