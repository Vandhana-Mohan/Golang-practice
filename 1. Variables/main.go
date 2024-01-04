package main

import (
	"fmt"
	"time"
)

// import "math"

func main(){
	print ("hello")
	fmt.Println("Hello")
	fmt.Println("new")
	fmt.Printf("jkhkj dsssd sdssd")

	currentTime := time.Now()
	fmt.Println(currentTime)

	specificTime := time.Date(2022, time.January, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)
}
