# For loop

**For** loop It is a control structure used to execute code repeatedly until a condition is met.

> [!NOTE]
> In Go there is only one type of loop, which is the for.

## Syntax

1. Initialization
2. Condition
3. Update

> [!IMPORTANT]
> Parentheses are not required around the condition.

```go
for [initialization]; [condition]; [update] {
    // code
}
```

**Example:**

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

### For While

```go
i := 1

for i <= 5 {
    fmt.Println(i)
    i++
}
```

### Do-While

```go
i := 1

for {
    fmt.Println(i)
    i++
    if i > 5 {
        break
    }
}
```

### For Range

The `range` keyword is used to iterate over collections of data such as array, slice, string, map.

> [!NOTE]
> It uses two variables, the first one representing the index and the second one the values.
**Syntax:**

```go
for i, value := range [array | slice | string | map] {
    //code
}
```

**Example:**

```go
// Array
var arr = [5]int{1,2,3,4,5}

for i, value := range arr {
    fmt.Printf("Indice: %d, Valor: %d\n", i, value)
}
```

**Example2:**

> [!IMPORTANT]
> We can use the blank operator to ignore either of the two variables.

```go
var arr = [3]string{"Daniela", "Pedro", "Andres"}
// Operator blank _
for _, value := range arr {
    fmt.Println(value)
}
```

**Example3:**

> [!IMPORTANT]
> Since version 1.22 you can use range with an integer as argument

```go

for i := range 5 {
    fmt.Println(i)
}

// Output:
// 0
// 1
// 2
// 3
// 4
```
