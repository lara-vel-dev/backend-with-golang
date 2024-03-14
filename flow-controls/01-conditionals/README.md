# Conditionals

In Go, the syntax for conditionals is similar to other programming languages. You can use the `if` statement to execute a block of code if a condition is true, `else if` to check multiple conditions, and `else` to execute a block of code if none of the conditions are met.

```go
package main

import (
	"fmt"
)

func main() {
	playerExperience := 50

	if playerExperience >= 50 {
		fmt.Println("You are at an advanced level")
	} else if playerExperience >= 25 {
		fmt.Println("You are at an intermediate level")
	} else {
		fmt.Println("You are at a beginner level")
	}
}
```

Something different in Go is that you can declare a variable inside the `if` statement and use it only within the scope of the `if` block (including the `else if` and `else` blocks). This is called an [If with a short statement](https://go.dev/tour/flowcontrol/6) and is useful, for instance, when you want to evaluate the returned value of a function and use it in the `if` block.

```go
package main

import (
	"fmt"
)

func getPlayerExperience() int {
	return 50
}

func main() {
	if playerExperience := getPlayerExperience(); playerExperience >= 50 {
		fmt.Printf("You are at an expert level with %d experience\n", playerExperience)
	} else if playerExperience >= 25 {
		fmt.Printf("You are at an intermediate level with %d experience\n", playerExperience)
	} else {
		fmt.Printf("You are at a beginner level with %d experience\n", playerExperience)
	}

	// This won't work because playerExperience is only available in the if statement
	// fmt.Printf("You have %d experience\n", playerExperience)
}
```

In the example above, the `playerExperience` variable is declared and initialized with the value returned by the `getPlayerExperience` function. This variable is only available within the scope of the `if` block.

Ifs with a short statement can also be used to check for errors when calling a function that returns two or more values, usually the result and an error. You can use the short statement to check if the error is not `nil` and handle it accordingly.

```go
package main

import (
	"fmt"
)

func getPlayerExperience() (int, error) {
	return 50, nil
}

func main() {
	if playerExperience, err := getPlayerExperience(); err != nil {
		fmt.Println("Error getting player experience")
	} else if playerExperience >= 50 {
		fmt.Printf("You are at an expert level with %d experience\n", playerExperience)
	} else if playerExperience >= 25 {
		fmt.Printf("You are at an intermediate level with %d experience\n", playerExperience)
	} else {
		fmt.Printf("You are at a beginner level with %d experience\n", playerExperience)
	}
}
```

In the example above, the `getPlayerExperience` function returns two values: the player experience and an error. The short statement is used to initialize the `playerExperience` and `err` variables with the returned values, then, the `err` variable is checked in the `if` block to handle any errors and the `playerExperience` variable is used in the `else if` and `else` blocks.

> [!NOTE]
> Probably, the error handling could be improved, but this is just an example of how to use the ifs with a short statement. 