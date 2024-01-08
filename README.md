#### About GO

- Go is a ==<u>open source</u>== programming language developed by Google from 2007
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

1. `package main` lets the Go compiler know that we want this code to compile and ==run as a standalone program, as opposed to being a library that's imported by other programs.== 
2. Go, C, and Rust are all languages where the code is first converted to machine code by the compiler before it's executed. - thats why Go is a compiler language
3. `import fmt` imports the `fmt` (formatting) package. The formatting package exists in Go's standard library and lets us do things like print text to the console.
4. `func main()` defines the `main` function. `main` is the name of the function that acts as the entry point for a Go program.

#### Variables in Go:

a value for eg: 120 or "hello" or 12.34 is literal in GO

+, -, *, / for arithmetic operation


```
fmt.Println(20 * 3) // Prints: 60  
fmt.Println(55.21 / 5) // Prints: 11.042  
fmt.Println(9 % 2) // Prints: 1
```

==const declaration== - const once declared cant be changed and has to be assigned value while initializing. Constants can't use the `:=` short declaration syntax. Constants can be character, string, boolean, or numeric values. They _can not_ be more complex types like slices, maps and structs

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

Some types can be converted the following way:

```
temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat) 

// Converting a float to integer will truncate the floating point portion
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

Return values may be given names, and if they are, then they are treated the same as if they were new variables defined at the top of the function.

Named return values are best thought of as a way to document the purpose of the returned values.

A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions. They can harm readability in longer functions.

```go
func getCoords() (x, y int){
  // x and y are initialized with zero values

  return // automatically returns x and y
}
```

is the same as:

```go
func getCoords() (int, int){
  var x int
  var y int
  return x, y
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

If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error that crashes the program.

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

An array is a collection of data elements of the same type, where we can access each **element** by an **index**. An array is a fixed size ordered list of elements with the same data type. Arrays are useful for collecting and accessing multiple related values.

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

All of this happens automatically using slices, while this is not possible with arrays

Go provides us with a function, `append` that handles adding to and resizing a slice

```
books := []string{"Tom", "Men"}  
books = append(books, "Frankenstein")  
books = append(books, "Dracula")  
fmt.Println(books) // [Tom Men Frankenstein Dracula]
```

To pass an array parameter into a function, we provide a local name, square brackets, and the data type. The difference between slice and array parameters is whether the number of elements is stated:

```
func printFirstLastArray(array [4]int) {    
	fmt.Println("First", array[0])    
	fmt.Println("Last", array[3])
}

func printFirstLastSlice(slice []int) {   
	length := len(slice)    
	if (length > 0) {        
		fmt.Println("First", slice[0])        
		fmt.Println("Last", slice[length-1])    
	}
}
```

Due to Go being a pass by value language, modifying a normal array parameter won’t create permanent change. Sometimes this can be useful in performing local calculations

```
// Changes to the array will only be local to the function  
func changeFirst(array [4]int, value int) {  
    array[0] = value  
}
```

In order to retain changes, a slice can be used

```
// Changes to the slice parameter will be permanent  
func changeFirst(slice []int, value int) {  
    if (len(slice) > 0) {  
        slice[0] = value  
    }  
}
```

#### Maps

A map is an unordered collection of **keys** and **values**.  Maps can be initialized with or without data.

Here is an example of connecting a key of type `string` with a value of type `int`:

|Key|Value|
|---|---|
|Joe|2126778723|
|Angela|4089978763|
|Shawn|3143776876|
|Terell|5026754531|
Unlike array, values in a map are not accessed by indices. Maps allow for very fast lookups by organizing the values for retrieval.

In Go, there are two ways to create a map.

==Creating a map with make==

We can use the `make` function to create an empty map. The format is:

```
variableName := make(map[keyType]valueType)
prices := make(map[string]float32)
```

==Creating a map with values==
If we know some map values, we can specify them as follows:

```
variableName := map[keyType]valueType{    
  name1: value1,    
  name2: value2,    
  name3: value3,
}

contacts := map[string]int{  
    "Joe":    2126778723,  
    "Angela": 4089978763,  
    "Shawn":  3143776876,  
    "Terell": 5026754531,  
}
```

to access a value in map 

	variable := yourMap[keyValue]

If a key is not in the map, a default value for value type is returned. We can also get a second return value to determine if the key is in the map. We can look up values with a key. We can also get a `status` value to determine if the key was set in the map.

```
customer,status := customers["billy"]  
  
if status {  
  fmt.Println("we found the customer")  
} else {  
  fmt.Println("no such customer!")  
}
```

Maps are also easy to add key-value pairs or to change the value of an existing pair.

	yourMap[newKey] = newValue

So to add a new customer balance, we could do:

```
customers["Samantha"] =  1.25
```

to change an existing value 

```
customers["Samantha"] =  2.75
```

Go allows us to remove elements using the **`delete`** function:

```
delete(yourMap, keyValueToDelete)

delete(contacts, "Gary") 

// If we call the `delete` function with a key that is not in the map nothing bad happens.
```

#### Structs

a way to group several variables into one custom data type. These types make the code cleaner, more intuitive, and less error-prone. grouping together related variables is done using `struct`

A group of related variables can be defined as a struct. Each variable within a struct is known as a field.

Define structs

A struct must be defined before it can be used in a program. The definition of a struct includes its name and its fields. A **field** is one of the internal variables  inside a struct. We use the following template:

```
// Struct names begin with a capital letter in Go
type NameOfStruct struct {  
	// Struct fields go here
}
```

if we want to define a 2D point with an x and y coordinate. We could define two variables `x` and `y` and use them throughout our program. instead we can create a struct called `Point` which contains both coordinates. Defining `Point` in this way logically groups together the relevant data types. We would define the struct for `Point` like so:

```
type Point struct { 
	x int  
	y int
}
```

Creating an Instance of a Struct

To use a struct we defined, we have to create an instance of it.

	p1 := Point{x: 10, y: 12} 
	or 
	var p1 = Point{x: 10, y: 12}

Go allows us to rely on default values as well. We can omit fields

	p1 := Point{x: 10}  
	// y will be set to 0

we can omit all fields to rely only on default values:

```
p1 := Point{} // x and y will be set to 0
```

The order of our struct definition allows us to avoid labeling our fields. The values are assigned from left to right according to how the fields are defined in the struct from top to bottom.

	p1 := Point{10, 12}  // Same as var p1 = Point{10, 12}
	When not using labels, we must provide values for every field;

Access and modify a struct’s variables

access individual fields within struct using the name of the variable, `.`, and the name of the field

```
john := Student{"John", "Smith", 14, 9}
fmt.Println(john.firstName)
```

modify the value of a field with an assignment statement:

```
john.age = 15
```

Functions that Access a Struct

```
type Rectangle struct {  
  length float32  
  width  float32  
}

func (rectangle Rectangle) area() float32 {  
  return rectangle.length * rectangle.width  
}

// functions associated with a struct are written outside of the struct!
```

Defining a function in this way will only pass in a copy of the rectangle: that is, we will not be able to use the function to alter the value of a field.

to call the `area()` function like so:

```
rectangle.area()
```


Pointers to a Struct

Without pointers, when a variable is passed into a function, only a copy of it is used inside the function. We can use pointers to modify values in our structs within a function.

```
type Employee struct {  
  firstName string  
  lastName string  
  age int  
  title string  
}
```
We must first create an instance of `Employee` and then we create a pointer that will point to this instance.

```
func main() {  
  steve := Employee{“Steve”, “Stevens”, 34, “Junior Manager”}  
  pointerToSteve := &steve  
}
```

We can now use this pointer to change the values of the fields for `steve`. There are two ways to do this in Go:

```
(*pointerToSteve).firstName or  pointerToSteve.firstName
```

We can use these pointers to modify structs in our functions. Consider the following example:

```
func (rectangle *Rectangle) modify(newLength float32){  
	rectangle.length = newLength
}
// notice `rectangle` is also a pointer
```

 It is dereferenced without the use of the dereferencing operator just like `pointerToSteve`

Arrays of Structs

to deal with many structs of the same type, We can use them in an array together. Let’s say we want to create an array of the following points: {1, 1} {7, 27} {12, 7} {9, 25}

We create an array of `Point`s like so:

```
points := []Point{{1, 1}, {7, 27}, {12, 7}, {9, 25}}
```

If the points have names, we can also create the array like this:

```
a = {1, 1}
b = {7, 27}
c = {12, 7}
d = {9, 25}
points := []Point{a, b, c, d}
```

access the contents of this array same as array. We can also access and modify the fields of each of the array elements.

```
points := []Point{{1, 1}, {7, 27}, {12, 7}, {9, 25}}
fmt.Println(points[0]) // Output will be {1, 1}

points := []Point{{1, 1}, {7, 27}, {12, 7}, {9, 25}}
points[1].x = 8
points[1].y = 16
fmt.Println(points[1]) // Output will be {8, 16}
```

Nested Structs

When we have complex groups of fields in our structs they can be combined into their own struct.

```
type Name struct{  
  firstName string  
  lastName string  
}  
  
type Employee struct{  
  name Name  
  age int  
  title string  
}

the `Employee` struct, has two separate fields for the first and last name of the employee. We can combine those two into their own struct called `Name`.
```

We create an instance of `Employee` like 

```
carl := Employee{Name{“Carl”, “Carlson”}, 32, “Engineer”}
```

To access the fields of the nested struct (`Name` in this case), we chain together the field accesses like so:

```
fmt.Println(carl.name.lastName) // Output will be “Carlson”
```

define the employee struct with the `Name` struct anonymously like so:

```
type Employee struct{  Name  age int  title string}

// notice the `Name` file has no associated variable name with it
```

Composing a struct in this way allows us to access the `firstName` and `lastName` fields directly from the `Employee` struct.

```
carl := Employee{Name{“Carl”, “Carlson”}, 32, “Engineer”}
fmt.Println(carl.firstName) // Output will be “Carl”
fmt.Println(carl.lastName) // Output will be “Carlson”
```

We cannot have two anonymous fields of the same type (i.e., two `Name` fields) as that would make it impossible to tell which field is being accessed (which `firstName` if two anonymous `Name` fields). An anonymous field is used to field access easier and leads to cleaner code.

#### Errors in Go

Go programs express errors with `error` values. An Error is any type that implements the simple built-in [error interface](https://blog.golang.org/error-handling-and-go):

```go
type error interface {
    Error() string
}
```

When something can go wrong in a function, that function should return an `error` as its last return value. Any code that calls a function that can return an `error` should handle errors by testing whether the error is `nil`.

```go
// Atoi converts a stringified number to an interger
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
    // because "42b" isn't a valid integer, we print:
    // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
    // Note: 'parsing "42b": invalid syntax' is returned by  .Error() method
    return
}
// if we get here, then i was converted successfully
```

A `nil` error denotes success. A non-nil error denotes failure.

### The Error Interface

Because errors are just interfaces, you can build your own custom types that implement the `error` interface. Here's an example of a `userError` struct that implements the `error` interface:

```go
type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}
```

It can then be used as an error:

```go
func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}
```

Go programs express errors with `error` values. Error-values are any type that implements the simple built-in [error interface](https://blog.golang.org/error-handling-and-go).

In Go, an `error` is just another value that we handle like any other value – no special keywords 

The Go standard library provides an "errors" package that makes it easy to deal with errors.

```go
var err error = errors.New("something went wrong")
```

