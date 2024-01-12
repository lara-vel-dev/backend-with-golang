# Variables and data types 
NOTE: You have to have in mind the last code structure.

## Data types in Go
- `int` is used to declare integers.
- `float32` is used to declare 32-bit floating numbers.
- `float64` is used to declare 64-bit floating numbers.
- `string` is used to declare strings.
- `bool` is used to declare boolean values.
### Example
``` GO
var a int
var b float64
var c float32
var d string
var e bool
```

## Things you should know before declaring variables in Go
- In Go, variables are explicitly declared and used by the compiler.
- `var` or `const` declares 1 or more variables.
### Example
  ``` GO
  // with const
  const pi, gravity float64 = 3.14, 9.81
  ftm.Println(pi,gravity)

 // with var
  var age, year = 23, 2024
  ftm.Println(age, year)
  ```
- `var` can change its value overtime but it can't change its type.
- `const` can't change its value nor its type.
- Go will infer the type of initialized variables.
- Variables declared without a corresponding initialization are zero-valued.
- The := syntax is shorthand for declaring and initializing a variable. This syntax is only available inside functions.
### Example
  ``` GO

  // Calculating the area of a triangle
	const base float64 = 20
	const height float64 = 10
	area_triangle := base * height / 2
  ftm.println("The area is:", area_triangle)
  ```  
