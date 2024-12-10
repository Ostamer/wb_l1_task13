package main

import "fmt"

func main() {
	// Два числа, которые мы хотим обменять
	a := 4
	b := 100

	fmt.Println("До обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// Обмен значениями при помощи XOR
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("После обмена:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
