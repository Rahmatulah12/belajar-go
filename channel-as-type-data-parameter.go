package main

import (
	"fmt"
	"runtime"
)

/*
	Variabel channel bisa di-passing ke fungsi lain sebagai parameter. Caranya dengan
	menambahkan keyword chan ketika deklarasinya.
*/

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan string)

	for _, each := range []string{"John Wick", "John Due", "Michael"} {
		go func(who string) { // eksekusi goroutine pada fungsi IIFE
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

// fungsi untuk menerima go routine lewat channel
func printMessage(what chan string) {
	fmt.Println(<-what)
}
