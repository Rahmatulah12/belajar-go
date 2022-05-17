package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		defer digunakan untuk mengakhirkan eksekusi baris kode. Ketika eksekusi sudah
		sampai pada akhir blok fungsi, statement yang di defer baru akan dijalankan.
	*/
	defer fmt.Println("dijalankan terakhir")
	/*
		Fungsi os.Exit() berada dalam package os . Fungsi ini memiliki sebuah parameter
		bertipe numerik yang wajib diisi. Angka yang dimasukkan akan muncul sebagai exit status
		ketika program berhenti.
	*/
	os.Exit(1)
	fmt.Println("dijalankan awal")
}
