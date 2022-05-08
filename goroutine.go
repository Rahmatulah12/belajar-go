package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Goroutine adalah sebuah mini thread
	// Goroutine berjalan asynchronous
	/*
		Untuk menerapkan goroutine, proses yang akan dieksekusi sebagai goroutine harus
		dibungkus kedalam fungsi. Pada saat pemanggilan fungsi tersebut, ditambahkan keyword
		go didepannya. Dengan demikian proses yang ada dalam fungsi tersebut dideteksi
		sebagai goroutine baru.
	*/
	runtime.GOMAXPROCS(2) // menentukan jumlah proccessor yang aktif
	go print(10, "Hello") // goroutine yang akan dieksekusi
	print(10, "Apa kabar")

	var input string
	/*
		Fungsi ini akan meng-capture semua karakter sebelum user menekan tombol enter, lalu
		menyimpannya pada variabel.
	*/
	fmt.Scanln(&input) // blocking, untuk menunggu goroutine ini selesai dieksekusi
}

func print(till uint8, message string) {
	for i := uint8(0); i < till; i++ {
		fmt.Println(message, "yang ke", (i + 1))
	}
}
