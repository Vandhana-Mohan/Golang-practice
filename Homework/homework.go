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

Create a function that calculates the area of a rectangle.
If the input is not a valid pair of numbers (length and width),
output the message "Please enter valid dimensions." Return two arguments:
The first argument is a string "The area of the rectangle is calculated."
The second argument is the calculated area.
Example output for valid dimensions: "The area of the rectangle is calculated", 30.

write a program that uses zero value, short declaration operator, multiple initializations,
var when specifity is required, and blank identifier

*/

package main

import "fmt"

var hello = "hi"

const number = 42.42

func calculateAreaRectangle(len, width int) (string, int) {
	area := len * width
	return "The area of the rectangle is calculated : ", area
}

func checkNumberEvenOrOdd(num int) string {
	if num%2 == 0 {
		return fmt.Sprintf("%d is an even number", num)
	} else {
		return fmt.Sprintf("%d is an odd number", num)
	}
}
func main() {

	var a int                          // zero value
	b := 32                            // short declaration operator
	c, d := 34, "hello"                // multiple initialization
	af, bf, cf, _, ef := 1, 2, 3, 4, 5 // blank identifier
	var ff float32 = 32.34             // var when specificity is required

	fmt.Printf("this %v variable is declared outside main and is type %T\n", hello, hello)
	fmt.Printf("this %v variable is declared outside main and is type %T\n", number, number)

	var num int
	fmt.Print("Enter a number : ")
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Please Enter a valid number.")
	} else {
		result := checkNumberEvenOrOdd(num)
		fmt.Println(result)
	}

	var length, width int

	fmt.Print("Enter the length of the rectangle : ")
	_, err1 := fmt.Scan(&length)

	fmt.Print("Enter the width of the rectangle : ")
	_, err2 := fmt.Scan(&width)

	if err1 != nil || err2 != nil {
		fmt.Println("Please enter valid number.")
	} else {
		result, value := calculateAreaRectangle(length, width)
		fmt.Println(result, value)
	}

	fmt.Println(a, b, c, d, af, bf, cf, ef, ff)

}
