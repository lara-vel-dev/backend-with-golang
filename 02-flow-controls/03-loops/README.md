# Loops

## For loop
The `for` statement allows us to repeat a list of statements multiple times. The basic `for` loop has 3 components that are separated by a semicolon:
- `initial statement`: this is executed before the first iteration.
- `condition expression`: this is evaluated before every iteration.
- `post statement`: this is executed at the end of every iteration.

The initial statement is often a short variable declaration and the variables declared there are visible only in the scope of the `for` statement. The loop will stop iterating once the condition is fulfilled.

### Example
``` Go
var sum int = 0
for i := 0; i < 10; i++ {
    sum += i
}

fmt.Println(sum) // 45
```

> [!NOTE]
> Unlike other languages, there are no parentheses surrounding the 3 components of the `for` statement and the braces are always required.

## While loop
The `while` loop allows us to execute a block of code until a certain condition is met. Unlike other programming languages, Golang doesn't has a keyword for a `while` loop but, we use the `for` loop to perform this functionality. This loop only has 1 component, the `condition`. The loop evaluates the `condition` and if its `true`, the statements inside the loop are executed and the `condition` is evaluated again. If the `condition` is `false`, the loop ends.

### Example 
```Go 
var number int = 1

for number <= 5 {
    fmt.Println(number) // 1 2 3 4 5
    number++
}
```

## Do-while loop
The `do-while` loop is a variant of the `while` loop. This execute the code block once, before checking the condition, then it will repeat the loop as long as the condition is `true`. We use the same `for` loop to provide the functionality.

### Example 
```Go 
var number int = 1

for {
    fmt.Println(number)
    number++

    if number > 5 {
        break
    }
}
```

## Range loop
The `range` loop is used with various data structures for iterating over an element. It is most commonly used in loops for iterating over elements of an array, map, slice and string. The syntax is as follows:

```Go 
for index, value := range datastructure {
    fmt.Println(value)
}
```
The `index` is the value that is been accesed, `value` is the actual value that we have in each iteration and `data structure` is used for holding the values that are accesed by the loop.

### For-range with arrays and slices
We can use the `for-range` loop to access the individual index and elements of an array.

```Go
// Creates an array
numArray := [3]int{10, 20, 30}

for idx, item := range numArray {
    fmt.Printf("index: %d, value: %d", idx, item) // index: 0, value: 10 index: 1, value: 20 index: 2, value: 30
}
```
### For-range with maps
We can use the `for-range` loop with map to access key-value pairs. 

```Go
// Creates a map
subjectGrades := map[string]float32{"Maths": 9.5, "English": 9.2, "History": 8.4}

for subject, grade := range subjectGrades {
	fmt.Println(subject, ":", grade) // Maths : 9.5 English : 9.2 History : 8.4
}
```

### For-range with strings
We can use the `for-range` with string to access individual characters of a string along with their respective index. 

``` Go
var str string = "Hello"

for i, item := range str {
    fmt.Printf("%d=%c\n", i, item) // 0=H 1=e 2=l 3=l 4=o 
}
```

## For-forever loop
When the condition is omitted, the loop runs forever (until you break out of it or return from the enclosing function), in Go that's what we call a `for-forever` loop. This loop is commonly used in scenarios where the program should run indefinitely, such as:
- Waiting for user input.
- Continuously listening for network connections.
- Running a server that should never stop unless explicitly terminated.

### Example
```Go
for {
    fmt.Println("This loop will run indefinitely")
}
```

<div>
<a href="https://github.com/lara-vel-dev/backend-with-golang/blob/main/02-flow-controls/02-switch/README.md" >
	<strong>⬅️ Previous topic</strong>
</a>
&emsp;
<a href="https://github.com/lara-vel-dev/backend-with-golang/blob/main/02-flow-controls/04-break-continue-goto/README.md" >
	<strong>➡️ Next topic</strong>
</a>
</div>
