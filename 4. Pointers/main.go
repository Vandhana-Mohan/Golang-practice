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
}