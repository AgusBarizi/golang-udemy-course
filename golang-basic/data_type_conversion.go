package main

import "fmt"

func main() {
	var nilai32 int32 = 12500
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println("int32", nilai32)
	fmt.Println("konversi int32 ke int64", nilai64)
	fmt.Println("konversi int32 ke int8", nilai8) //nilai tidak valid karena tidak dpt menampung data

	name := "Agus"
	charIndexAsByte := name[0]
	charIndexAsString := string(charIndexAsByte)

	fmt.Println("first character on Agus Barizi is", charIndexAsString)

}
