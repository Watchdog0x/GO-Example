package main

import (
	"fmt"
	"strconv"
)

func main() {

	//** NUMERIC TYPES **//

	// uint8      the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// uint     either 32 or 64 bits
	// int      same size as uint

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers
	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32

	//int type
	var i1 int8 = -128     //min value
	fmt.Printf("%T\n", i1) // => int8

	var i2 uint16 = 65535  //max value
	fmt.Printf("%T\n", i2) // => int16

	var i3 int64 = -324_567_345  // underscores are used to write large numbers for a better readability
	fmt.Printf("%T\n", i3)       // => int64
	fmt.Printf("i3 is %d\n", i3) // => i3 is -324567345 (underscores are ignored)

	//float64 type
	var f1, f2, f3 float64 = 1.1, -.2, 5. // trailing and leading zeros can be ignored
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	//rune type
	var r rune = 'f'
	fmt.Printf("%T\n", r) // => int32 (rune is an alias to int32)
	fmt.Printf("%x\n", r) // => 66,  the hexadecimal ascii code for 'f'
	fmt.Printf("%c\n", r) // => f

	//bool type
	var b bool = true
	fmt.Printf("%T\n", b) // => bool

	//string type
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s) // => string

	//array type
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers) // =>  [4]int

	//slice type
	var cities = []string{"London", "Bucharest", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities) // => []string

	//map type
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
	}
	fmt.Printf("%T\n", balances) // => map[string]float64

	//struct type
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you) // => main.Person

	//pointer type
	var x int = 2
	ptr := &x                                                 // pointer to int
	fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr) // => ptr is of type *int with value 0xc000016168

	//function type
	fmt.Printf("%T\n", f) // => func()

	var x = 3   //int type
	var y = 3.2 //float type

	// x = x * y //compile error ->  mismatched types

	x = x * int(y) // converting float64 to int
	fmt.Println(x) // => 9

	y = float64(x) * y //converting int to float64
	fmt.Println(y)     // => 28.8

	x = int(float64(x) * y)
	fmt.Println(x) // => 259

	//In Go types with different names are different types.
	var a int = 5   // same size as int64 or int32 (platform specific)
	var b int64 = 2 // int and int64 are not the same type

	// a = b // error: cannot use b (type int64) as type int in assignment
	a = int(b) // converting int64 to int (explicit conversion required)

	// preventing unused variable error
	_ = a

	//** CONVERTING NUMBERS TO STRINGS AND STRINGS TO NUMBERS **//

	s := string(99)            // int to rune (Unicode code point)
	fmt.Println(s)             // => 99, the ascii code for symbol c
	fmt.Println(string(34234)) // => 34234 is the unicode code point for 薺

	// we cannot convert a float to a string similar to an int to a string
	// s1 := string(65.1) // error

	// converting float to string
	var myStr = fmt.Sprintf("%f", 5.12)
	fmt.Println(myStr) // => 5.120000

	// converting int to string
	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1) // => 34234

	// converting string to float
	var result, err = strconv.ParseFloat("3.142", 64)
	if err == nil {
		fmt.Printf("Type: %T, Value: %v\n", result, result) // => Type: float64, Value: 3.142
	} else {
		fmt.Println("Cannot convert to float64!")
	}

	// Atoi(string to int) and Itoa(int to string).
	i, err := strconv.Atoi("-50")
	s = strconv.Itoa(20)
	fmt.Printf("i Type is %T, i value is %v\n", i, i) // => i Type is int, i value is -50
	fmt.Printf("s Type is %T, s value is %q\n", s, s) // => s Type is string, s value is "20"

	// 1. int to string
	s := strconv.Itoa(i)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// 2. string to int
	x, _ := strconv.Atoi(s2)
	fmt.Printf("f Type is %T, f value is %d\n", x, x)

	// 3. float64 to string
	y := strconv.FormatFloat(f, 'f', -1, 64)
	fmt.Printf("y's type: %T, y's value: %s\n", y, y)

	// 4. string to float64
	e, _ := strconv.ParseFloat(s1, 64)
	fmt.Printf("e's type: %T, e's value: %.2f\n", e, e)

	type age int        //new type, int is the underlying type
	type oldAge age     //new type, int (not age) is the underlying type
	type veryOldAge age //new type, int (not age) is the underlying type

	// new type speed (underlying type uint)
	type speed uint

	// s1, s2 of type speed
	var s1 speed = 10
	var s2 speed = 20

	// performing operations with the new types
	fmt.Println(s2 - s1) // -> 10

	// uint and speed are different types (they have different names)
	var x uint

	// x = s1  //error different types

	// correct
	x = uint(s1)
	_ = x

	// correct
	var s3 speed = speed(x)
	_ = s3

	// declaring a variable of type uint8
	var a uint8 = 10
	var b byte // byte is an alias to uit8

	// even though they have different names, byte and uit8 are the same type because they are aliases
	b = a // no error
	_ = b

	// declaring a new alias named second for uint
	// type alias_name = type_name
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint

	//no need to convert operations (same type)
	fmt.Printf("Minutes in an hour: %d\n", hour/60) // => Minutes in an hour: 60

}

func f() {
}
