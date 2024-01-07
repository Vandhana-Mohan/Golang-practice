# Golang-practice
#### About GO

- Go is a <u>open source</u> programming language developed by Google from 2007
- Go is associated with backend development due to its features that make it well-suited for building scalable and concurrent server-side applications.
- Developers use Go to create backend services, APIs, micro services, and other server-side components.
- Go and Golang both same - Go is official name while Golang is a more like nickname.
- Go is often used for cloud computing and web development, but it also has applications in DevOps, data science, and AI. 
#### Features GO

- open source
- performant
- multiple-cores
- concurrency
- compiled network
- clean syntax
- powerful standard library
- garbage collection
- portable (compiles on many OS's)

#### Why GO

1. Efficient Compilation
2. Efficient Execution
3. Ease of programming
#### Similarity between JavaScript and Go:

1. **C-Like Syntax:**
    - Both JavaScript and Go have C-like syntax, which means they share some common structural elements like loops, conditionals, and functions.
2. **Concurrent Programming:**
    - Go is known for its built-in support for concurrent programming. While JavaScript is single-threaded, Go has goroutines and channels for concurrent execution. Understanding the concepts of concurrency in Go may be a new and valuable experience for a JavaScript developer.
3. **Garbage Collection:**
    - Both JavaScript and Go have automatic garbage collection, meaning developers don't need to manage memory manually. If you're used to this feature in JavaScript, it will be a familiar concept in Go as well.
4. **Simple and Readable Code:**
    - Go, like JavaScript, values simplicity and readability. The language is designed to be straightforward and easy to understand, which can help with a smooth learning curve.
5. **Strong Typing:**
    - While both languages are statically typed, Go's type system is more explicit compared to JavaScript. If you are comfortable with static typing in JavaScript, adapting to Go's type system should not be too challenging.
#### Differences between JavaScript and Go:

1. **Backend vs. Frontend:**
    - While JavaScript is commonly used for both front-end (browser-based) and back-end (Node.js) development, Go is often associated with back-end development. Learning Go may expose you to different types of applications and use cases.
2. **Learning New Concepts:**
    - Go introduces some unique concepts such as goroutines, channels, and defer statements. These might be new to someone coming from JavaScript, but they can be powerful tools once understood.
3. **Ecosystem and Libraries:**
    - The ecosystems and libraries of JavaScript and Go are different. Familiarizing yourself with the Go ecosystem and its standard library will be part of the learning process.


Javascript Vs Go
1. function  || func
2. var x = 5 || var x int = 5 or x:= 5
3. cant declare with int and then change to string
4. no while in Go
5. only == for conditional check
6. comments :  // and /* ...  */

#### Start Go:

create mod: this is how we initialize our project module
```go mod init example```

run mod file:
``` go run main.go ``` it will find main and execute code

To check the role of a function,  go doc followed by a package or the function of a package. For example:

```
$ go doc fmt
```

To find more information about a package’s function:

```
$ go doc fmt.println
```

#### Variables in Go:

a value for eg: 120 or "hello" or 12.34 is literal in GO

+, -, *, / for arithmetic operation


```
fmt.Println(20 * 3) // Prints: 60  
fmt.Println(55.21 / 5) // Prints: 11.042  
fmt.Println(9 % 2) // Prints: 1
```

==const declaration== - const once declared cant be changed and has to be assigned value while initializing

```
const funFact = "Hummingbirds"  
fmt.Println(funFact)
```

for numeric, there is int, float and complex

- int is whole counting numbers
- float is fractional
- complex is with imaginary unit i

There are 11 different integer types, 2 different floating-point types, and 2 different complex number types.

Signed integers is negative, but unsigned integers is only positive.

for Bool - true / false

==var declaration==

```
var lengthOfSong uint16  // u - unsigned
var isMusicOver bool  
var songRating float32
var hello int8
```

to assign values to variables, we can use the assignment operator (`=`) followed by a value.

```
var kilometersToMars int32
kilometersToMars = 62100000

or another way is 

var kilometersToMars int32 = 62100000

```

eg: 

```
func main() {

var numOfFlavors int
numOfFlavors = 57

fmt.Println(numOfFlavors)

var flavorScale float32 = 5.8

fmt.Println(flavorScale)

}
```

To declare string variables - remember string text always in double quotes only

```
var nameOfSong string  
var nameOfArtist string
nameOfSong = "Stop Stop"  
nameOfArtist = "The Julie Ruin"
var description string  
description = nameOfSong + " is by: " + nameOfArtist + "."  
fmt.Println(description)  
// Prints "Stop Stop is by: The Julie Ruin.

// use the `+` operator to join them known as string concatenation
```

Default Values

```
var classTime uint32  
var averageGrade float32  
var teacherName string  
var isPassFail bool  
  
fmt.Println(classTime) // Prints 0  
fmt.Println(averageGrade) // Prints 0  
fmt.Println(teacherName) // Doesn't print anything  
fmt.Println(isPassFail) // Prints false
```


#### Inferring Variables

declare a variable without  stating its type using the short declaration `:=` operator

```
nuclearMeltdownOccurring := true  
radiumInGroundWater := 4.521  
daysSinceLastWorkplaceCatastrophe := 0  
externalMessage := "Everything is normal. Keep calm and carry on."

```

we can also write above as 

```
var nuclearMeltdownOccurring = true  
var radiumInGroundWater = 4.521  
var daysSinceLastWorkplaceCatastrophe = 0  
var externalMessage = "Everything is normal. Keep calm and carry on."

notice - var is used in this case and := is replaced with ==
```

===Recommended===  use int instead of int8 or int16


+=, -=, *=, /=  are  ===shorthand operators===

===Multiple Variable declaration===

```
var part1, part2 string  
part1 = "To be..."  
part2 = "Not to be..."

for this syntax, both variables must be the same type.
```

```
If we already know what values we want to assign our variables we can use `:=`

quote, fact := "Bears, Beets, Battlestar Galactica", true  
fmt.Println(quote) // Prints: Bears, Beets, Battlestar Galactica  
fmt.Println(fact) // Prints: true
```

eg :-

```
func main () {

var magicNum, powerLevel int32

magicNum = 2048

powerLevel = 9001

fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)

amount, unit := 10, "doll hairs"

fmt.Println(amount, unit, ", that's expensive...")
}
```

eg: 

```
package main

import "fmt"

func main(){

publisher, writer, artist, title := "DizzyBooks Publishing Inc.", "Tracey Hatchet", "Jewel Tampson", "Mr. GoToSleep"

year, pageNumber := 1997, 14

grade := 6.5

fmt.Println(title, "written by", writer, "drawn by", artist, "and published by", publisher, "released in the year", year, "and the page", pageNumber, "has a grade of", grade)

title, writer, artist, year, pageNumber, grade = "Epic Vol. 1", "Ryan, N. Shawn", "Phoebe Paperclips", 2013, 160, 9.0

fmt.Println(title, "written by", writer, "drawn by", artist, "and published by", publisher, "released in the year", year, "and the page", pageNumber, "has a grade of", grade)

}
```

#### The fmt Package

`fmt` is one of Go’s core packages.  fmt helps us format data hence known as format package.
In addition to println, we have print, printf, sprint, sprintln, sprintf, scan and many more.

Println - prints to terminal and has default styling built in. it prints every argument with space and adds "\\n" at the end.
for eg :

```
fmt.Println("Println", "formats", "really well")
fmt.Println("Right?")
```

Which prints:

```
Println formats really well
Right?
```

print does not have the default formatting and styling.

for eg:

```
fmt.Print("The answer is", ": ")
fmt.Print("12")
```

The above code snippet would print:

```
The answer is: 12 (notice no default spacing and no line break at the end.)
```

Println and Print has the ability to concatenate strings i.e. combine different strings into a single string.

With `fmt.Printf()`, we can _interpolate_ strings, or leave placeholders in a string and use values to fill in the placeholders. 

```
guess := "C"
fmt.Printf("Is %v your final answer?", guess)
// Prints: Is C your final answer?

selection1 := "soup"  
selection2 := "salad"  
fmt.Printf("Do I want %v or %v?", selection1, selection2)  
// Prints: Do I want soup or salad?

Notice that the placement of the arguments matters! 

remember : printf has no line break at the end.

```

Go has a variety of useful verbs (check [their documentation](https://golang.org/pkg/fmt/#hdr-Printing) for a comprehensive list

%T - specifies the type of variable

%d - displays the int,  %f - displays the float and %.2f - displays only 2 decimals of float

for eg:

```
specialNum := 42  
fmt.Printf("This value's type is %T.", specialNum)  
// Prints: This value's type is int.  
  
quote := "To do or not to do"  
fmt.Printf("This value's type is %T.", quote)  
// Prints: This value's type is string.

votingAge := 18  
fmt.Printf("You must be %d years old to vote.", votingAge)  
// Prints: You must be 18 years old to vote.

gpa := 3.8  
fmt.Printf("You're averaging: %f.", gpa)  
// Prints: You're averaging 3.800000.

gpa := 3.8  
fmt.Printf("You're averaging: %.2f.", gpa)  
// Prints: You're averaging 3.80.
```

`fmt.Sprint` and `fmt.Sprintln` functions are part of the `fmt` package and are used for formatting and concatenating strings. 

```
grade := "100"  
compliment := "Great job!"  
teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)  
  
fmt.Print(teacherSays)  
// Prints: You scored a 100 on the test! Great job!

fmt.Sprintln does the same but adds "\n" at the end of line.
```

If we need to interpolate a string, without printing it, then we can use `fmt.Sprintf()`.

Just like `fmt.Printf()`, `fmt.Sprintf()` can also use verbs:

```
correctAns := "A"
answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)
fmt.Print(answer) 
// Prints: And the correct answer is… A!

`fmt.Sprintf()` works very similarly to `fmt.Printf()`, the major difference is that `fmt.Sprintf()` returns its value instead of printing it out!
```

Scan - allows us to get user input

```
fmt.Println("How are you doing?")  
  
var response string  
fmt.Scan(&response)  
  
fmt.Printf("I'm %v.", response)

How are you doing?  
good  
I'm good.

How are you doing?  
not bad  
I'm not. (notice Only the `not` part is saved, since it was separated from `bad` by a space.)
```

#### Conditionals

If statement

similar concept to javascript
```alarmRinging := true  
if alarmRinging {  
  fmt.Println("Turn off the alarm!!")  
}

if (alarmRinging) {  
  fmt.Println("Turn off the alarm!!")  
}
The parentheses are optional but typically, we’ll see the `if` statement written without parentheses.
```

if... else statement

```
isHungry := false  
if isHungry {  
  fmt.Println("Eat the cookie")  
} else {  
  fmt.Println("Step away from the cookie...")  
} Notice, the `else` statement does _not_ accept a condition.
``` 

if... else if... else statement

```position := 2  
  
if position == 1 {  
  fmt.Println("You won the gold!")  
} else if position == 2 {  
  fmt.Println("You got the silver medal.")  
} else if position == 3 {  
  fmt.Println("Great job on bronze.")  
} else {  
  fmt.Println("Sorry, better luck next time?")  
}
```

switch statement

```
clothingChoice := "sweater"  
  
switch clothingChoice {  
case "shirt":  
  fmt.Println("We have shirts in S and M only.")  
case "polos":  
  fmt.Println("We have polos in M, L, and XL.")  
case "sweater":  
  fmt.Println("We have sweaters in S, M, L, and XL.")  
case "jackets":  
  fmt.Println("We have jackets in all sizes.")  
default:  
  fmt.Println("Sorry, we don't carry that")  
}  
// Prints: We have sweaters in S, M, L, and XL.
```

Comparison and Logical operators

| operator | Meaning |
| ---- | ---- |
| `==` | Is equal to |
| `!=` | Is NOT equal to |
| `<` | Less than |
| `>` | Greater than |
| `<=` | Less than or equal to |
| `>=` | Greater than or equal to |
| `&&` | And |
| \|\| | Or |
| `!` | Not |
#### Scoped Short Declaration Statement

```x := 8  
y := 9  
if product := x * y; product > 60 {  
  fmt.Println(product, "  is greater than 60")  
}

Notice that `product` is separated from the condition using a semi-colon `;`

switch season := "summer" ; season {  
case "summer":  
  fmt.Println("Go out and enjoy the sun!")  
}

and the variable product and season is local to that block.
```

Randomizing

Go has a `math/rand` library that helps us generate a random integer.
```
import (  "math/rand"  "fmt")
func main() {  
	fmt.Println(rand.Intn(100))
}

In our `main` function, we’re printing a random number using `rand` and the `Intn()` method. With the argument of `100`, the maximum value the method will return is 99. Looking at the line `fmt.Println(rand.Intn(100))`, it should print a random integer from `0` to `99`.
```

But we need to use Go seed to actually make numbers random. so we use time for seed
- Go uses a seed value to as a starting point for generating random numbers.
- Unique seed values can be created using time, namely `rand.seed(time.Now().UnixNano())`

```package main  
  
import (  
  "fmt"  
  "math/rand"  
  "time"  
)  
  
func main() {  
  rand.Seed(time.Now().UnixNano())  
  fmt.Println(rand.Intn(100))  
}
```

#### Functions

```
func doubleNum(num int) int {  
  return num * 2  
}

fmt.Println(doubleNum(x)) // Prints: 10
```
Each function has its own specific scope, take a look at the code:

```
package main
import "fmt"
func performAddition() { 
	x := 5  
	y := 7  
	fmt.Println("The sum of", x, "and", y, "is", x + y)
}
func main() {  
	performAddition()  
	fmt.Println("What if", x, "was different?")
}
```

There are three different scopes present in this example:

- The global scope, which contains the function definitions for `main()` and `performAddition()`.
- `performAddition()` has a local scope, which defines `x` and `y`.
- `main()` has a local scope also. It can access `performAddition()` because that’s defined on the same scope level as `main()` but can’t access the internals of `performAddition`‘s scope (i.e., `x` or `y`).

A function can be given a return _type_, the type of a value that will be returned by the function. At the call site, the return value can be stored within a variable of the same type as the function’s return.

```
func getLengthOfCentralPark() int32 {  
	var lengthInBlocks int32  
	lengthInBlocks = 51  
	return lengthInBlocks
}

func main() {  
  var centralParkLength int32  
  centralParkLength = getLengthOfCentralPark()  
  fmt.Println(centralParkLength) // Prints: 51  
}
```

we can also provide functions with information using _parameters_. Function parameters are variables that are used within the function to use in some sort of computation or calculation. When calling a function, we give _arguments_, which are the values that we want those parameter variables to take. We give our function parameters types when defining the function:

```
func multiplier(x int32, y int32) int32 {  return x * y}
```

Since both parameters have the same type, we could write it as:

```
func multiplier(x, y int32) int32 {  
   return x * y
}
```

call our function with literal values as arguments:

```
func main() {  
   var product int32  
   product = multiplier(25, 4)  
   fmt.Println(product) // Prints: 100
}
```

We can also call our function with variables as arguments:

```
func main() {  
  var mainX, mainY, newProduct int32  
  mainX = 6  
  mainY = 7  
  newProduct = multiplier(mainX, mainY)  
  fmt.Println(newProduct) // Prints: 42
}
```

Functions also have the ability to return multiple values. 

```
func GPA(midtermGrade float32, finalGrade float32) (string, float32){  
averageGrade := (midtermGrade + finalGrade) / 2  
var gradeLetter string  
	if averageGrade > 90 {    
		gradeLetter = "A"  
	} else if averageGrade > 80 {    
		gradeLetter = "B"  
	} else if averageGrade > 70 {    
		gradeLetter = "C"  
	} else if averageGrade > 60 {   
		gradeLetter = "D"  
	} else {    
		gradeLetter = "F"  
	}  
	return gradeLetter, averageGrade 
}
```

### Deferring Resolution

We can delay a function call to the end of the current scope by using the `defer` keyword. `defer` tells Go to run a function, but at the end of the current function. 

#### Addresses

The space that the computer allocates is called an _address_. Each address is marked as a unique numerical value.

To find a variable’s address we use the `&` operator followed by the variable name, like so:

```
x := "My very first address"
fmt.Println(&x) // Prints 0x414020
```

#### Pointers

Pointers are variables that specifically store addresses.

We even set the data type of the addresses’ value for the pointer. For instance:

```
var pointerForInt *int  
  
minutes := 525600  
  
pointerForInt = &minutes  
  
fmt.Println(pointerForInt) // Prints 0xc000018038

`pointerForInt` will store the address of a variable that has an `int` data type. `*` operator signifies that this variable will store an address and the `int` portion means that the address contains an integer value.

or in short

minutes := 55  
  
pointerForInt := &minutes
```

 address -  values are stored
 pointers - keep track of address

Dereferencing / indirecting : use pointer to access the address and change its value

```lyrics := "Moments so dear"  
pointerForStr := &lyrics  
  
*pointerForStr = "Journeys to plan"  
  
fmt.Println(lyrics) // Prints: Journeys to plan
```

Go is a pass-by-value language while Js which is pass by reference language

==pass-by-value== In Go, when you pass a variable to a function, you're passing a copy of the value, not the original variable itself. Any modifications made to the parameter inside the function do not affect the original variable outside the function.

```func addHundred(num int) {  
  num += 100  
}  
  
func main() {  
  x := 1  
  addHundred(x)  
  fmt.Println(x) // Prints 1  
}
```

but if we do using returned value

```func addHundred(num int) int {  
  num += 100  
  return num
}  
  
func main() {  
  x := 1  
  x = addHundred(x)  
  fmt.Println(x) // Prints 101
}
```

or using pointers

```
func addHundred (numPtr *int) {  
  *numPtr += 100  
}  

// function has a parameter of a pointer for an integer. By passing the value of // a pointer (which is an address) to `addHundred()`, we  dereference 
// the address and add `100` to its value. 
  
func main() {  
  x := 1  
  addHundred(&x)  // `addHundred()` expects a pointer for an argument
  fmt.Println(x) // Prints 101  
}
```

The primary difference in above 2 ways lies in how the modification of the original value is done: one uses a returned value, and the other uses a pointer. Using a return value may be more straightforward in many cases. Using pointers can be more efficient when working with large data structures, but it also adds complexity. 

- The `*` operator can be used to assign a pointer the type of the value its address holds.
- The `*` operator can also be used to dereference a pointer and assign a new value.

#### Looping

Classic for loop (definite loop)

```
for number := 0; number < 5; number++ {  
  fmt.Print(number, " ")  // Output: 0 1 2 3 4
}  

no () around for loop, and := to initialize the variable within for loop stmt
```

For as a While Loop  (Indefinite loop)

```
number := 0 // Initialize a variable to be used inside the loop  
for number < 5 {  
  fmt.Println(number)  
  number++ // Update the variable being used  
}
```

A `break` statement changes when a loop will end. While a `continue` statement changes what will happen in each loop.

#### Looping and Arrays

Each map and array has a set amount of items that they contain. In Go, the `range` keyword can be used to work through these items one at a time within a loop. For example:

```
Array:

letters := []string{"A", "B", "C", "D"}  
for index, value := range letters {  
  fmt.Println("Index:", index, "Value:", value)  
}

Output:
Index: 0 Value: A  
Index: 1 Value: B  
Index: 2 Value: C  
Index: 3 Value: D
```

```
map:

addressBook := map[string]string{  
  "John": "12 Main St",  
  "Janet": "56 Pleasant St",  
  "Jordan": "88 Liberty Ln",  
}  
for key, value := range addressBook {  
  fmt.Println("Name:", key, "Address:", value)  
}

output:
Name: John Address: 12 Main St  
Name: Janet Address: 56 Pleasant St  
Name: Jordan Address: 88 Liberty Ln
```

#### Arrays

An array is a collection of data elements of the same type, where we can access each **element** by an **index**.

==create array==
```
var playerScores [4]int  // provide the number of elements
fmt.Println(playerScores)

this example creates an empty array of integer values with space for 4 elements.
```

```
triangleSides := [3]int{15, 26, 30} // array with size and the element

triangleAngles := [...]int{30, 60, 90} 

// the compiler determine the length automatically using `...` ellipsis syntax.

```

==Access Array Values with Indices==

```
students = [3]string{"Jill", "Fred", "Sasha"}  
 
fmt.Println(students[0]) // Go uses `0` as the first index of the array
```

==Modify array values==

syntax : array[index] = value

```
myArray := [4]int{10, 24, 5, 47}
myArray[2] = 33 // output : 10, 24, 33, 47
```

#### Slices

**Slices** are a data collection type similar to arrays, but they have the ability to change their size while arrays has a fixed size.

==Create slice ==

We can create a slice from an array, or by itself.

```
// creating slice by itself
  
var numberSlice []int  // creating empty slice 
stringSlice := []string{}  // creating empty slice 
  
names := []string{"Kathryn", "Martin", "Sasha", "Steven"} // create slice with elements 

```

```
// creating slice from array

array := [5]int{2, 5, 7, 1, 3}  
 
sliceVersion := array[:] // This is a slice of the whole array 
fmt.Println(sliceVersion)  // output : [2 5 7 1 3]
partialSlice := array[2:5] // This slice contains the elements at index 2, 3, 4 
fmt.Println(partialSlice)  
// [7 1 3] Modifying the slice will still update the original array.
```

==len==
`len` is a function which returns the length of an array or slice passed into it.
```
favoriteThings := [2]string{"Raindrops on Roses", "Whiskers on Kittens"}  
fmt.Println(len(favoriteThings))  // 2  

```

Arrays only have a length, but slices have length and capacity.

A slice is resizeable, so there is a difference between:

- Its length, the current number of elements it holds
- Its **capacity**, the number of elements it can hold before needing to resize itself.

A slice’s capacity can be accessed through the `cap` function

```
slice := []string{"Fido", "Fifi", "FruFru"}  
  
fmt.Println(slice, len(slice), cap(slice)) // [Fido Fifi FruFru] 3 3  
slice = append(slice, "FroFro")  
// After appending an element when the slice is at capacity  
// The slice will double in capacity, but increase its length by 1  
fmt.Println(slice, len(slice), cap(slice)) // [Fido Fifi FruFru FroFro] 4 6
```

when we added an element to a slice which was at full capacity the following occured:

- The new element was still able to be added
- The length increased to fit the new element
- The capacity doubled in size.

All of this happens automatically using slices, while this is not possible with arrays!

