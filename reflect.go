package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 10.5
	reflectValueOfNumber := reflect.ValueOf(number)
	fmt.Println(reflect.TypeOf(number))          // mengembalikan tipe data float64 (object reflect.Type)
	fmt.Println(reflect.ValueOf(number).Float()) // mengembalikan nilai/value number (object reflect.Value)
	// Mengakses nilai dalam bentuk interface
	fmt.Println(reflectValueOfNumber.Interface())           // mengembalikan interface{} (kosong)
	fmt.Println(reflectValueOfNumber.Interface().(float64)) // mengembalikan float64 (object float64)

	// pengecekan tipe data dengan method Kind()
	if reflectValueOfNumber.Kind() == reflect.Float64 {
		fmt.Println("number is float64")
	}

	// 	mengkases informasi property variable objek
	/*
		Reflect bisa digunakan untuk mengambil informasi semua property variabel objek cetakan
		struct, dengan catatan property-property tersebut bermodifier public.
	*/
	fmt.Println("")
	var people = &Peopele{
		Name: "John",
		Age:  30,
	}
	people.GetPropertyInfo()

	people.SetName("Jane")
	fmt.Println("")
	fmt.Println(people.Name)

	// mengakses informasi method variable objek
	var method = reflect.ValueOf(people).MethodByName("SetName")
	method.Call([]reflect.Value{reflect.ValueOf("John Doe")})
	fmt.Println(people.Name)
}

type Peopele struct {
	Name string
	Age  int
}

func (p *Peopele) GetPropertyInfo() {
	reflectValue := reflect.ValueOf(p)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem() // mengaganti nilai reflect value menjadi elemennya
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Nama :", reflectType.Field(i).Name)
		fmt.Println("Tipe Data :", reflectType.Field(i).Type)
		fmt.Println("Nilai :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (p *Peopele) SetName(name string) {
	p.Name = name
}
