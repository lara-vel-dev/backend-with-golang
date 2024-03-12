# Reading values

## Using the `fmt` package

Reading values from the standard input is a common task in programming. In Go, you can use the `Scan`, `Scanf` and `Scanln` functions from the `fmt` package to read values from the standard input.

The `Scan` function reads space-separated values from the standard input and stores them in the variables passed as arguments:

```go
package main

import "fmt"

func main() {
    // We can read a single value (until the first space)
    var name string
    fmt.Print("Enter your name: ") 
    fmt.Scan(&name) // Input: John Doe
    fmt.Println("Hello", name) // Output: Hello John

    // Or multiple values separated by spaces
    var name2 string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scan(&name2, &age) // Input: John 23
    fmt.Println("Hello", name2, "you are", age, "years old") // Output: Hello John you are 23 years old
}
```

> [!NOTE]
> Newlines are considered spaces by the `Scan` function, so, the input: 
> ```bash
> John
> 23
> ```
> will be read as `John` and `23`.

With the `Scanf` function, you can also read space-separated values from the standard input, but you can specify the format of the input using a format string (just like the examples of the `Printf` function in the [first section](https://github.com/lara-vel-dev/backend-with-golang/tree/main/the-basics/01-hello-world#functions-used-to-format-and-print-data)): 

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    // %s is a placeholder for a string and %d is a placeholder for an integer
    fmt.Scanf("%s %d", &name, &age)
    fmt.Println("Hello", name, "you are", age, "years old")

    // If you want to check for errors, you can use the returned value
    var name2 string
    var age2 int
    fmt.Print("Enter your name and age: ")
    readedLen, err := fmt.Scanf("%s %d", &name2, &age2)
    fmt.Println("Readed", readedLen, "values")

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Hello", name2, "you are", age2, "years old")
    } 
}
```

The `Scanln` function is similar to the `Scan` function, but **newlines are not considered spaces and it reads values until a newline is found**:

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scanln(&name, &age) // Input: John 23
    fmt.Println("Hello", name, "you are", age, "years old") // Output: Hello John you are 23 years old
}
```

> [!NOTE]
> As newlines **are not** considered spaces by the `Scanln` function and it reads values until a newline is found, the input:
> ```bash
> John
> 23
> ```
> will be read as `John`, and the `Scanln` function will return an error due to the unexpected newline.

## Using the `bufio` package

If you need to read multiple lines or a single line that may contain spaces, you can use the `bufio` package: 

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var currentFullName string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter your full name: ")

    // Read until a newline is found (Including the newline character)
		input, _ := reader.ReadString('\n')

    // Remove the newline character
		currentFullName = strings.TrimSpace(input)

    // Break the loop if the input is empty
		isEmpty := len(currentFullName) == 0
		if isEmpty {
			break
		}

		fmt.Println("Your full name is: " + currentFullName)
	}
}
```