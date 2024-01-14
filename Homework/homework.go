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

func main(){
	var num int
	fmt.Scan("Enter a number : ", &num)

	fmt.println(num)
}