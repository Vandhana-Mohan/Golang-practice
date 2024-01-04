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