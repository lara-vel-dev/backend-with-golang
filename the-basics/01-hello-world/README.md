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

In Go, the entry point to a program has to be a function called main within a package main. Go has a number of built­in functions, such as println, which can be used without reference but we can’t get very far without making use of Go’s standard library. The `import` keyword is used to declare the packages that are used by the code in the file.

> [!IMPORTANT]
> Go is strict about importing packages. It will not compile if you import a package but don’t use it.
