package main

import (
	"fmt"
)

func main() {

  menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

  fmt.Println("The menu:")
  
  for index, value := range menu{
    fmt.Println("Index: ", index, "Menu Item: ", value)
  }
  

  orders := map[string]string{
    "John": "Cheeseburgers",
    "Janet": "Hamburgers",
    "Jordan": "Fries",
  }
  
  orders["James"] = "Chicken Sandwich"
  
    fmt.Println("\nThe friend's orders:")
  
   for key, value := range orders{
    fmt.Println("Friend Name : ", key, "Order: ", value)
   }
}  