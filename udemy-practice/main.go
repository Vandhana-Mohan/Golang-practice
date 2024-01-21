package main

import (
	"fmt"
	"math/rand"
)

/*
create a random int between 0 and 250
*/
func main() {

	x := rand.Intn(250)

	fmt.Println(x)
}
