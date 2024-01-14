/*
Create functions that check if a number is odd or even
If the input is not a number
output a message "Please enter a number"
return two arguments
The first argument is a string "This number is an even number."
The second argument is a number 20.
"This number is an even number", 20.
If the number is odd, it should output the following
The first argument is a string "This number is an odd number."
The second argument is a number 11.
"This number is an odd number", 11.

*/

package main

import "fmt"

func checkNumberEvenOrOdd(num int) (string){
	if(num % 2 == 0){
		return fmt.Sprintf("%d is an even number", num)
	} else {
		return fmt.Sprintf("%d is an odd number", num)
	}
}
func main(){
	var num int
	fmt.Print("Enter a number : ")
	_,err := fmt.Scan(&num)

	if(err != nil){
		fmt.Println("Please Enter a valid number.")
	} else {
		result := checkNumberEvenOrOdd(num)
		fmt.Println(result)
	}
}
