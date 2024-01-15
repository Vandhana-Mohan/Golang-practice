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

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// explaining pass by value

func brainwash(saying *string) {
	*saying = "Beep Boop"
}
