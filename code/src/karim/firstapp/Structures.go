package main

import (
	"fmt"
)

// Doctor Sample Stuct
type Doctor struct {
	id        int
	name      string
	companion []string
}

func structExample() {

	fmt.Println("---% struct %---")
	aDoctor := Doctor{
		id:   1,
		name: "Jhon",
		companion: []string{
			"Jo", "Ram", "Jane",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.name)

	embadedStruct()
}

// Animal struct with Tag
type Animal struct {
	name   string
	origin string
}

//Bird Embaded struct
type Bird struct {
	Animal
	speed  int
	canFly bool
}

func embadedStruct() {
	b := Bird{
		Animal: Animal{name: "Emu", origin: "Australia"},
		speed:  48,
		canFly: false,
	}

	fmt.Println(b)

}
