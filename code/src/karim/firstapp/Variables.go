package main

import "fmt"

/*
There are 3 variables scopes
local
package global
Global available to all packages var I string = "global"
*/

var glb int = 90

// I gloabal visibility
var I string = "global"

var (
	firstName string = "John"
	lastName  string = "Dev"
	age       int    = 30
)

func variables() {

	fmt.Println("----%Variables%----")
	var i int
	i = 42
	fmt.Println(i)

	var j int = 27
	fmt.Println(j)

	k := 30
	fmt.Println(k)
	fmt.Printf("%v, %T", k, k)
	fmt.Println(glb)
	var glb int = 100
	fmt.Println(glb)

	conversion()
}

func conversion() {
	fmt.Println("----%Conversion%----")
	var i int = 40
	j := float32(i)

	fmt.Println(j)
}
