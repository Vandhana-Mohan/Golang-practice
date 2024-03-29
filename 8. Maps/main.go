package main

import "fmt"

func main() {
 orders := make(map[string]float32)
 fmt.Println(orders)

 donuts := map[string]int{
  "frosted": 10,
  "chocolate": 15,
  "jelly": 8,
 }
 fmt.Println(donuts)
 firstChoice := donuts["frosted"]
   fmt.Println(firstChoice) 

   secondChoice, status := donuts["bavarian cream"]
   fmt.Println(secondChoice) 
   fmt.Println(status) 

   if status {
      fmt.Println(secondChoice) 
   } else {
    fmt.Println("We do not sell that donut!") 
   }

   donuts["glazed"] = 12

  fmt.Println(donuts["glazed"])

   donuts["jelly"] = 3
   fmt.Println(donuts)

   delete(donuts, "chocolate")
  	fmt.Println(donuts)
}