package main

import "fmt"

func mapExample() {

	populations := make(map[string]int)

	populations = map[string]int{
		"Texus":      792303,
		"Florida":    692303,
		"New York":   18739430,
		"California": 233093,
	}

	fmt.Println(populations)

	populations["Alabama"] = 987979

	fmt.Println(populations)

	val, ok := populations["New York"]

	fmt.Printf("%v , %v \n", val, ok)
	fmt.Println(len(populations))
}
