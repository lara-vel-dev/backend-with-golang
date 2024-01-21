# Variables and data types 
> [!NOTE]
> Don't forget to add the main function before adding the code below

## Data types in Go

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

<div>
<a href="https://github.com/lara-vel-dev/backend-with-golang/blob/main/the-basics/01-hello-world" >
	<strong>⬅️ Previous topic</strong>
</a>
<a href="https://github.com/lara-vel-dev/backend-with-golang/blob/main/the-basics/03-operators" >
	<strong>➡️ Next topic</strong>
</a>
</div>

