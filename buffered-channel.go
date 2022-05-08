package main

import (
	"fmt"
	"runtime"
)

/*
	Channel secara default adalah un-buffered, tidak di-buffer di memori. Ketika ada goroutine
	yang mengirimkan data lewat channel, harus ada goroutine lain yang bertugas menerima
	data dari channel yang sama, dengan proses serah-terima yang bersifat blocking.
	Maksudnya, baris kode di bagian pengiriman dan penerimaan data, tidak akan akan
	diproses sebelum proses serah-terima-nya selesai.
	Buffered channel sedikit berbeda. Pada channel jenis ini, ditentukan jumlah buffer-nya.
	Angka tersebut akan menjadi penentu kapan kita bisa mengirimkan data. Selama jumlah
	data yang dikirim tidak melebihi jumlah buffer, maka pengiriman akan berjalan
	asynchronous (tidak blocking).
	Ketika jumlah data yang dikirim sudah melewati batas buffer, maka pengiriman data hanya
	bisa dilakukan ketika salah satu data sudah diambil dari channel, sehingga ada slot channel
	yang kosong. Dengan proses penerimaan-nya sendiri bersifat blocking.
*/

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2) // buffered channel ada diparam 2, dimulai dari 0, berarti kapasitas 3

	// membuat goroutine
	go func() {
		for {
			i := <-messages
			fmt.Println("Receive data:", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Send data:", i)
		messages <- i
	}
}
