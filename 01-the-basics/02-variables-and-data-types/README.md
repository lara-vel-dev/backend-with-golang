# Variables and data types 
> [!NOTE]
> Don't forget to add the main function before adding the code below

## Data types in Go
Data types are a set of related values that describe the operations that can be done on them and define the way they are stored. Go comes with several built-in data types.

### Basic types

- `int` is used to declare integers.
- `uint` is used to declare unsigned (positive or zero) integers.
- `float32` is used to declare 32-bit floating numbers.
- `float64` is used to declare 64-bit floating numbers.
- `string` is used to declare strings.
- `bool` is used to declare boolean values.

> [!NOTE]
> As said in the official [documentation](https://go.dev/ref/spec#Numeric_types), the size of `int` and `uint` depends on the architecture of the machine on which your program is running.

On a 32-bit machine, an `int` and `uint` will be 32 bits wide, while on a 64-bit machine, they will be 64 bits wide. Even though the above is not a problem in most cases, you can specify the size of `int`s and `uint`s by using the following types:

- `int8` and `uint8` : 8-bit signed and unsigned integers (from -128 to 127 and from 0 to 255 respectively). 
- `int16` and `uint16` : 16-bit signed and unsigned integers (from -32768 to 32767 and from 0 to 65535 respectively).
- `int32` and `uint32` : 32-bit signed and unsigned integers (from -2147483648 to 2147483647 and from 0 to 4294967295 respectively).
- `int64` and `uint64` : 64-bit signed and unsigned integers (from -9223372036854775808 to 9223372036854775807 and from 0 to 18446744073709551615 respectively).

#### Example
``` GO
var a int
var b uint
var c float64
var d float32
var e string
var f bool
```

### Advanced types

- `byte` is an alias for `uint8`. It is commonly used when working with files or streams of data.
- `rune` is an alias for `int32`. It is commonly used when working with `chars` and Unicode code points.
- `complex64` is used to declare complex numbers with float32 real and imaginary parts
- `complex128` is used to declare complex numbers with float64 real and imaginary parts. 

#### Example
``` GO
var a byte

var b rune
b = 'Ś'

// Here, we print the Unicode code point of the character
fmt.Printf("%U\n", a) // Output: U+015A

var c complex64
c = 2.5 + 3.5i

var d complex128
```

### Strings
A string is a sequence of characters with a definite length used to represent text. Go strings are made of individual bytes, usually one for each character. Strings can be created using double quotes or backticks. The difference between these is that double-quoted strings can't contain newlines and allow special escape sequences. Some common operations with strings are:
```Go
// Finds length of a string
len("Hello")

// Access a particular character in the string
"Hello"[3]

// Concatenates 2 strings together
"Hello " + "world"
```

Additionally, Go offers a built-in package called `strings` that provides many useful functions. You can read more about it [here](https://pkg.go.dev/strings). Some of the most common functions are:

```Go
// Returns true if the string contains the substring
strings.Contains("Hello world", "world") // true

// Replace a substring with another
strings.Replace("Hello world world", "world", "replaced", 1) // Hello replaced world
strings.Replace("Hello world world", "world", "replaced", 2) // Hello replaced replaced

// Replace all occurrences of a substring
strings.ReplaceAll("Hello world world", "world", "replaced") // Hello replaced replaced
// Equivalent to: strings.Replace("Hello world world", "world", "replaced", -1) // Hello replaced replaced

// Count the number of occurrences of a substring
strings.Count("Hello world world", "world") // 2

// Trim spaces from the beginning and end of the string
strings.TrimSpace(" Hello world ") // Hello world
// Equivalent to: strings.Trim(" Hello world ", " ") // Hello world

// Trim a substring 
strings.Trim("Hello world", "Hd") // ello worl
strings.TrimLeft("Hello world", "Hd") // ello world
strings.TrimRight("Hello world", "Hd") // Hello worl

// Change the case of the string
strings.ToUpper("Hello world") // HELLO WORLD
strings.ToLower("Hello world") // hello world
strings.ToTitle("Hello world") // HELLO WORLD
```

> [!NOTE]
> In the last example, even though the `ToUpper` and `ToTitle` functions return the same result, they are not equivalent in some cases. For example, although the character `ǆ` (Unicode code point `U+01C6`) seems to be an string of two characters (`d` and `ž`), it is actually a single character (technically called a `digraph`). The `ToTitle` function will return `ǅ` (Unicode code point `U+01C5`), while the `ToUpper` function will return `Ǆ` (Unicode code point `U+01C4`).

## Things you should know before declaring variables in Go
- In Go, variables are explicitly declared and used by the compiler.
- A variable's name must start with a letter and may contain letters, numbers or underscore symbol.
- We use cammelCase convention.
>[!NOTE]
> Cammel case is a style for writing compound words in which the first letter of each new word or phrase is capitalized, for example `var lastName string` .
- `var` or `const` declares 1 or more variables.
### Example
  ``` GO
  // With const
  const pi, gravity float64 = 3.14, 9.81
  ftm.Println(pi,gravity)

 // With var
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
	triangleArea := base * height / 2
  ftm.println("The area is:", triangleArea)
  ```

## Scope
In Go, scope refers to the part of the program where a variable or identifier is accessible. It determines where you can use that variable and modify its value. There are 3 types of scope
- **Local**: Are declared within functions, blocks of code enclosed in curly braces `{}` and are only accesible within the code block where they are declared.
  ```go
  package main
  
  import "fmt"

  func addNumbers() {

    // Local variable
    var sum int
    sum = 5 + 9
  }


  func main() {

    addNumbers()

    // Can't access sum out of its local scope
    fmt.Println("Sum is", sum)
  }
  ```
- **Global**: Are declared outside of any function but not within a package block. These variables are accessible from any part of the program.
  ```go
  package main
  
  import "fmt"

  // Declare global variable before main function
  var sum int

  func addNumbers () {

    // Local variable
    sum = 9 + 5
  }


  func main() {

    addNumbers()

    // Can access sum
    fmt.Println("Sum is", sum)
  }
  ```
- **Package**: Are declared outside of any function but within a package block. These variables are accessible throughout the entire package, which encompasses all the Go source code files within the same directory.
  ```go
  package greetings // Package name

  var message string = "Hello, from the package!"

  func printMessage() {
    fmt.Println(message)
  }

  func main() {
    fmt.Println(message) // Accessing the package-level variable from main
  }
  ```

<div>
<a href="./01-the-basics/01-hello-world/README.md" >
	<strong>⬅️ Previous topic</strong>
</a>
&emsp;
<a href="./01-the-basics/03-operators/README.md" >
	<strong>➡️ Next topic</strong>
</a>
</div>

