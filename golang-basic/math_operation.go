package main

import "fmt"

func main() {
	var x = 10
	var y = 3
	var z = x + y
	fmt.Println("z = x + y", z)
	z += 10
	fmt.Println("z + 10", z)
	var a float32
	a = float32(x / y)
	fmt.Println("a = x / y", a)
}
