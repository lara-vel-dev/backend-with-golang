# Hello World!

> [!IMPORTANT]
> - The extension of the programs must be **(file_name.go)**
> - Go is strictly typed, which means once you choose a type for a variable you can't change it later.

> [!NOTE]
> Go is a compiled language. Let's remember that compilation is the process of translating the source code that you write into a lower-level language.

## C-like Syntax
Go has C-like Syntax which means that it shares syntactical similarities with the C programming language. Go doesn't strictly adhere to the traditional C-like syntax, but it does share some similarities. Go has its unique features and design principles, it was developed with the aim of being simple, efficient, and easily readable, but some aspects that may be considered C-like are:
- Go uses braces ({}) to define blocks of code. However, it does not require semicolons at the end of statements, making the code cleaner and less prone to errors.
- Go follows a C-like approach to variable declarations, where the type comes after the variable name. `var x int`
- Go supports if-else statements, loops (for and a variation of while), and switch statements for flow control, which aligns with the C-like paradigm.
- Go incorporates common operators seen in C, including arithmetic, logical, and bitwise operators.

## Running Go Code
Let's start by creating a simple program and learning how to compile and execute it. 
> [!NOTE]
> Usually, the entry point file is called `main.go`.

```  GO
package main

import "fmt"

func main() {

    fmt.Println("Hello World GOLANG :)")
}
```

Open a shell/command prompt and run the program by entering:

``` Shell
go run main.go
```

In Go, the entry point to a program has to be a function called main within a package main. The `main` package in the Go language contains a main function, which shows that the file is executable. Go has a number of built­in functions, such as println, which can be used without reference but we can’t get very far without making use of Go’s standard library. The `import` keyword is used to declare the packages that are used by the code in the file.

> [!IMPORTANT]
> Go is strict about importing packages. It will not compile if you import a package but don’t use it.

## "fmt" Package
`fmt` is pronounced “fumpt” and is one of Go’s core packages. It's mainly used for printing information to the terminal. The `fmt` package has a broader purpose like helping us format data, for this reason, it's sometimes referred to as the format package. The package has three sets of functions based on their usage

### Functions used to format and print data
- `Print()` : When the arguments are strings, it concatenates them without any spacing and prints the result to the console. When none of the arguements is a string, the Print function adds spaces between them.
```Go
// String arguments
fmt.Print("How", "are", "you", "?") // Output: Howareyou?

// None string arguments
fmt.Print(10, 20) // Output: 10 20
```
- `Println()` : Always adds a space between its arguements and appends a new line or a line break at the end.
```Go
/*
    Output:
    My name is Gopher
    New line
*/

fmt.Println("My", "name", "is", "Gopher")
fmt.Println("New", "line")
```
- `Printf()` : Provides custom formatting of a string using one or more verbs. A verb is a placeholder for a named value (constant or variable) to be formatted according to these conventions:
    - `%v` represents the named value in its default format.
    - `%d` expects the named value to be an integer type.
    - `%f` expects the named value to be a float type.
    - `%T` represents the type for the named value.
 ```Go
var name string ="Lukman"
fmt.Printf("My name is %v", name) // Output: My name is Lukman
```
  
### Functions that formats the data but prints nothing
- `Sprint()` : Doesn’t print out anything. Rather, it returned a value. Afterward, we can use the returned value for later usage.
```Go
var user string = "Gopher"
var Feedback string = "Nice book!"
var userFeedback string = fmt.Sprint(user, "feedback on your book is", feedback)

fmt.Print(userFeedback) // Output: Gopher feedback on your book is Nice book!
```
- `Sprintln()` : Works like fmt.Sprint() but it automatically includes spaces between the arguments for us (just like fmt.Println() and fmt.Print())
```Go
var quote string = fmt.Sprintln("see here,", "no spaces!")
fmt.Print(quote) // Prints see here, no spaces!
```
- `Sprintf()` : When we have to interpolate a string, without printing it, then we can use fmt.Sprintf(). Just like fmt.Printf(), fmt.Sprintf() can also use verbs.
```Go
var user string = "userA"
var winner string = fmt.Sprintf("The winner is… %v!", user)

fmt.Print(winner) // Prints: The winner is is… userA!
```

### Functions to read user input from the terminal
- `Scan()` : Scans user input from the terminal and extracts text delimited by spaces into successive arguments. A newline is considered a space. This function expects an address of each argument to be passed.
```Go
fmt.Println("What's your name?")
fmt.Scan(&name) // Input: John

fmt.Println("What's your age?") 
fmt.Scan(&age) // Input: 25

fmt.Printf("%v is %d years old!", name, age) // Output: John is 25 years old!
```

> [!NOTE]
> If you’re ever stuck without internet access, you can get the documentation running locally via:
> `godoc ­http=:6060` and pointing your browser to http://localhost:6060

**[➡️ Next topic](https://github.com/lara-vel-dev/backend-with-golang/blob/main/the-basics/02-variables-and-data-types/README.md)**
