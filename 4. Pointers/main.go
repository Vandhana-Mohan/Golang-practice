package main

import "fmt"

func main(){
	var pointerForInt *int  
	
	minutes := 525600  
	
	pointerForInt = &minutes  
	
	fmt.Println(pointerForInt)

	lyrics := "Moments so dear"  

	pointerForStr := &lyrics  
	
	*pointerForStr = "Journeys to plan"  
	
	fmt.Println(lyrics) 

	greeting := "Hello there!"
	
	// Call your brainwash() below:
	
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)
}

// explaining pass by value

func brainwash(saying *string) {
	*saying = "Beep Boop"
}
