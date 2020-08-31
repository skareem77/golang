package main

import "fmt"

func complexNum() {

	fmt.Println("----%Complex number%----")
	var i complex64 = 1 + 2i
	fmt.Println(i)

	fmt.Printf("%v, %T\n", real(i), real(i))
	fmt.Printf("%v, %T\n", imag(i), imag(i))

	var n complex128 = complex(5, 7)

	fmt.Printf("%v, %T\n", n, n)
}
