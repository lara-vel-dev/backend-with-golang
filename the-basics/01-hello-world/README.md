# Hello World!

> [!IMPORTANT]
> - The extension of the programs must be **(file_name.go)**
> - Go is strictly typed, which means once you choose a type for a variable you can't change it later.

> [!NOTE]
> Go is a compiled language. Let's remember that compilation is the process of translating the source code that you write into a lower-level language.

## C-like Syntax
Go has C-like Syntax which means that it shares syntactical similarities with the C programming language. Go doesn't strictly adhere to the traditional C-like syntax, but it does share some similarities. Go has its unique features and design principles, it was developed with the aim of being simple, efficient, and easily readable, but some aspects that may be considered C-like are:
- Go uses braces ({}) to define blocks of code. However, it does not require semicolons at the end of statements, making the code cleaner and less prone to errors.
- Go follows a C-like approach to variable declarations, where the type comes after the variable name. For example: `var x int`
- Go supports if-else statements, loops (for and a variation of while), and switch statements for flow control, which aligns with the C-like paradigm.
- Go incorporates common operators seen in C, including arithmetic, logical, and bitwise operators.
 
```  GO
package main

import "fmt"

func main() {
	fmt.Println("Hello World GOLANG :)")
}
```

If we want to run our program in GO, we have to use 

``` Shell
go run name_of_program.go
```
