package main

import "fmt"


func main(){
	fmt.Println("♥️") // print a emoji
	fmt.Println(`hi im practicing udemy exercise 1 with emoji and string template literals`) // print raw string literal

	var a int
	var b float64
	var c complex64
	var d string 
	var e rune 
	var f byte 
	var g bool

	add, bdd:= 23, "hi"
	fmt.Println("a", a, "b", b, "c", c, "d", d, "e", e, "f", f, "g", g, add, bdd)

	for i:= 1; i <= 10; i++{
		fmt.Println(i)
	}
}