# Hello World!

> [!IMPORTANT]
> - The extension of the programs must be **(file_name.go)**
> - Go is strictly typed, which means once you choose a type for a variable you can't change it later.

> [!NOTE]
> Go is a compiled language. Let's remember that compilation is the process of translating the source code that you write into a lower-level language.

## C-like Syntax
Go has C-like Syntax which means that it shares syntactical similarities with the C programming language. Go doesn't strictly adhere to the traditional C-like syntax, but it does share some similarities. Go has its unique features and design principles, it was developed with the aim of being simple, efficient, and easily readable, but some aspects that may be considered C-like are:
- Go uses braces ({}) to define blocks of code. However, it does not require semicolons at the end of statements, making the code cleaner and less prone to errors.
- Go follows a C-like approach to variable declarations, where the type comes after the variable name. `var name type`
- Go supports if-else statements, loops (for and a variation of while), and switch statements for flow control, which aligns with the C-like paradigm.
- Go incorporates common operators seen in C, including arithmetic, logical, and bitwise operators.

## Running Go Code
Let's start by creating a simple program and learning how to compile and execute it. 
> [!NOTE]
> Usually, the entry point file is called `main.go`.
Copy the following code:
```  GO
package main

import "fmt"

func main() {

    // Prints Hello world in Go
    fmt.Println("Hello World GOLANG :)")
}
```

> [!NOTE]
> A line that starts with // is known as comment. The Go compiler ignores comments and helps us to understand how the code works. Go supports 2 different styles of comments: //is a single-line comment and /**/is a multi-line comment.

Open a shell/command prompt and run the program by entering:

``` Shell
go run main.go
```
There you go, you successfully compiled and executed your first Go program! :D

> [!NOTE]
> Newlines, spaces and tabs are known as whitespace (because you can't see them). Go mostly doesn't care about whitespace, we use it to make programs easier to read. (You could remove this lines and the program would behave in exactly the same way).
>
> Despite this, go has a built-in tool called `gofmt` that uses a predefined set of rules to format your code in a standard way. This is useful when working in a team, as it means that everyone's code looks the same. You can manually run `gofmt` on a file by running `gofmt -w file_name.go` in your terminal or you can set up your editor to automatically run `gofmt` every time you save a file (The official extension for VSCode does this by default).

In Go, the entry point to a program has to be a function called main within a package main. The `main` package in the Go language contains a main function, which shows that the file is executable. Go has a number of built­in functions, such as `println`, which can be used without reference but we can’t get very far without making use of Go’s standard library. The `import` keyword is used to declare the packages that are used by the code in the file.

> [!IMPORTANT]
> Go is strict about importing packages and using variables. It will not compile if you import a package but don’t use it nor if you declare a variable but don’t use it.

## A Quick Look Over Go's Packages
Go programs are read top to bottom, left to right. (like a book) The first line says this:
```Go
package main
```
This is known as a “package declaration”. Every Go program must start with a package declaration. Packages are Go's way of organizing and reusing code. There are two types of Go programs: executables and libraries. Executable applications are the kinds of programs that we can run directly from the terminal. Libraries are collections of code that we package together so that we can use them in other programs. We will explore libraries in more detail later, for now just make sure to include this line in any program you write.

## "fmt" Package
`fmt` is pronounced “fumpt” and is one of Go’s core packages. It's mainly used for printing information to the terminal. The `fmt` package has a broader purpose like helping us format data, for this reason, it's sometimes referred to as the format package. The package has three sets of functions based on their usage

### Functions used to format and print data
- `Print()` : When the arguments are strings, it concatenates them without any spacing and prints the result to the console. When none of the arguments is a string, the Print function adds spaces between them.
```Go
// String arguments
fmt.Print("How", "are", "you", "?") // Output: Howareyou?

// None string arguments
fmt.Print(10, 20) // Output: 10 20
```
- `Println()` : Always adds a space between its arguments and appends a new line or a line break at the end.
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
    - `%v` represents the named value in its default format. This is the more generic verb since it adapts to the type of the value and displays its string representation.
        - `%+v` adds field names when printing `struct`s.
    - `%s` expects the named value to be a string type.
    - `%d` expects the named value to be an integer type.
    - `%f` expects the named value to be a float type.
    - `%t` expects the named value to be a boolean type.
    - `%T` represents the type for the named value.

You can find a complete list of verbs [in the official documentation](https://pkg.go.dev/fmt).

 ```Go
package main

import "fmt"

// We haven't seen structs yet, but for now you can think of them as named collections of fields
type Person struct {
	Name string
	Age  int
}

func main() {

	var name string = "Lukman"
	fmt.Printf("My name is %s\n", name) // Output: My name is Lukman

	var age int = 23
	fmt.Printf("My age is %d\n", age) // Output: My age is 23

	person1 := Person {
		Name: name,
		Age:  age,
	}

	fmt.Printf("%v\n", person)  // Output: {Lukman 23}
	fmt.Printf("%+v\n", person) // Output: {Name:Lukman Age:23}
}

```
  
### Functions that formats the data but prints nothing
- `Sprint()` : Doesn’t print out anything. Rather, it returned a value. Afterward, we can use the returned value for later usage.
```Go
var user string = "Gopher"
var feedback string = "Nice book!"

// Please note that we needed to add the space between the arguments ourselves
var userFeedback string = fmt.Sprint(user, " feedback on your book is: ", feedback)

fmt.Println(userFeedback) // Output: Gopher feedback on your book is: Nice book!
```
- `Sprintln()` : Works like fmt.Sprint() but it automatically includes spaces between the arguments for us (just like fmt.Println() and fmt.Print())
```Go
var quote string = fmt.Sprintln("see here,", "no spaces!")
fmt.Print(quote) // Prints: see here, no spaces!
```
- `Sprintf()` : When we have to interpolate a string, without printing it, then we can use fmt.Sprintf(). Just like fmt.Printf(), fmt.Sprintf() can also use verbs.
```Go
var user string = "userA"
var winner string = fmt.Sprintf("The winner is… %s!", user)

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

<a href="./02-variables-and-data-types/README.md" >
	<strong>➡️ Next topic</strong>
</a>
</div>
