# Switch case

The switch statement in Go, like in many other languages, is a selection control mechanism and it allows a variable or expression's value to change the flow of a program's executions. It works like an if-else ladder, checking multiple conditions at once. This statement starts with the keyword `switch` followed by the expression and then a series of `cases`.

## Example
```Go
var mcDifficulty string = "peaceful"

switch mcDifficulty {
    case "peaceful": fmt.Println("Enemies don't spawn.") // Prints this statement
    case "easy": fmt.Println("Enemies spawn but they deal with low damage.")
    case "normal": fmt.Println("Enemies spawn and deal between 1 and 1.5 hearts of damage.")
    case "hard": fmt.Println("Enemies spawn and deal more damage than on normal.") 
    default: fmt.Println("The difficulty you chose doesn't match the existing ones.") 
}
```

The value of the expression (in this case `mcDifficulty`) is compared to the expression following each `case` keyword (`peaceful`, `easy`, `normal`, `hard`). If they are equal to the case expression, then the corresponding statement followed by the `:` will be executed.


> [!NOTE]
> Like an `if` statement, each case is checked top down and the first one to succeed is chosen. A switch also has a default case that will be executed in case none of the `cases` matches the value (similar to how the `else` works in the `if` statement).

You can also compare multiple values in the same `case` expression. To do so, you'll need to separate each value with a comma `","` such as the following example:

## Multiple case values example
```Go
package main 

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // package that converts strings into numbers and vice versa
)

func main() {
    var grade int
    scanner := bufio.NewScanner(os.Stdin)
    // Gets user's input
    scanner.Scan()

    // Stores user's input into grade
    grade = strconv.ParseInt(scanner.Text())

    switch grade {
        case 10: fmt.Println("Congrats! You got a perfect grade")
        case 9, 8: fmt.Println("You passed the subject!")
        case 7, 6: fmt.Println("You passed, but you can do better :)")
        default: fmt.Println("You failed the subject :(")
    }
}
```

