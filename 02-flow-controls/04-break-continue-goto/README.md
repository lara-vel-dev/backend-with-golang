# Break, Continue and Goto

When you're writing code, you often want to control how and when certain parts of your program run. Think of your program as a set of instructions where you might sometimes need to pause, skip steps or jump to a specific parts of the instruction, this is where `break`, `continue` and `goto` come in.

## Break
The `break` statement is used to end the execution of `for`, `switch`, or `select` statement in which it appears. Once `break` statement is executed, the program exits the loop or case block and continues with the code immediately following it.  

**Behavior:**
- Applicable only within `for`, `switch`, and `select` statements. 
- It does not evaluate further conditions or iterations once executed. 

### Example
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            fmt.Println("Exiting loop at i =", i)
            break // Exits the loop when i equals 5
        }
        fmt.Println("i =", i)
    }
    fmt.Println("Loop has ended.")

    /*
    Output: 
    i = 1
    i = 2
    i = 3
    i = 4
    Exiting loop at i = 5
    Loop has ended.
    */
}
```
- The `for` loop iterates over numbers from 1 to 10.
- When `i == 5`, the `break` statement is triggered, exiting the loop.
- The program skips the remaining  iteration (6 to 10) and proceeds after the loop.

## Continue
The `continue` statement skips the current iteration of the loop in which it appears and immediately jumps to the next iteration. Unlike `break`, it does not end the loop; instead, it bypasses any code following continue for that iteration.

**Behavior:**
- Only valid within `for` loops. 
- Commonly used to skip unwanted iterations based on a condition.

### Example
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue // Skip even numbers
        }
        fmt.Println("Odd number:", i)
    }

    /*
    Output: 
    Odd number: 1
    Odd number: 3
    Odd number: 5
    Odd number: 7
    Odd number: 9
    */
}
```
- The loop iterates over numbers 1 to 10.
- The condition `i%2 == 0` checks if i is even.
- For even numbers, the continue statement is executed, skipping the `fmt.Println` statement for that iteration.

## Goto
The `goto` statement provides unconditional jumps to a labeled statement within the same function. While it is a powerful tool, its use is discouraged except in specific scenarios like error handling or breaking out of deeply nested loops, as it can lead to less readable and maintainable code.

**Behavior**
- Requires a label to jump to, defined as an identifier followed by a colon (`LabelName`:).
- Control is transferred directly to the labeled statement, skipping any intermediate code.

### Example
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            goto End // Jump to the labeled statement
        }
        fmt.Println("i =", i)
    }
    return

End:
    fmt.Println("Exited loop via goto.")

    /*
    Output: 
    i = 1
    i = 2
    i = 3
    i = 4
    Exited loop via goto.
    */
}
```
- Labels must be defined within the same function as the `goto` statement.
- Labels cannot be nested.
- Overusing `goto` can create complex and hard-to-follow logic, often referred to as “spaghetti code.”

### Good practices for using goto
1. **Use for Error Handling:** Ideal for consolidating cleanup or exit logic in case of errors.
2. **Break Nested Loops:** Useful for cleanly exiting multiple layers of loops.
3. **Use Descriptive Labels:** Labels should clearly describe their purpose.
4. **Keep goto Local:** Minimize the distance between `goto` and its target label to avoid confusion.
5. **Document Its Use:** Always explain why `goto` was chosen to maintain code clarity.
6. **Avoid Overuse:** Only use `goto` when it genuinely improves readability or avoids complexity.

### When not to use
- To replace standard loops or control structures.
- For regular program flow or logic, as it can lead to unmaintainable code.
- When it reduces clarity—prioritize readability and structure.
