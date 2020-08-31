package main

import "fmt"

func arrays() {
	fmt.Println("---% array %---")
	grades := [3]int{2, 3, 1}

	fmt.Printf("%v\n", grades)

	var students [3]string

	students[0] = "John"
	students[1] = "Ali"

	fmt.Printf("%v, length:%v\n", students[0], len(students))

	slice()

}

func slice() {

	fmt.Println("---% slice %---")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := a[:] //all elements

	c := a[3:] //from 4th elemts

	d := a[3:6]

	e := a[:5]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	arr := make([]int, 2, 100)

	fmt.Println(arr)

	arr = append(a, 1, 2, 3)
}
