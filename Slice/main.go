package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// declaring a string slice, by default is initialized with nil or uninitialized

	var cities []string

	fmt.Println("cities is equal to nil: ", cities == nil) // -> cities is equal to nil:  true

	fmt.Printf("cities: %#v\n", cities) // -> cities: []string(nil)

	// we can not assign elements to nil slice:

	// cities[0] = "Paris" // -> runtime error

	// declaring a slice using a slice literal

	numbers := []int{2, 3, 4, 5} // on the right hand-side of the equal sign is a slice literal

	fmt.Println(numbers) // => [2 3 4 5]

	// creating a slice using the make() built-in function

	// creating a slice with 2 int elements initialized with zero.

	nums := make([]int, 2) // the same as []int{0, 0}.

	fmt.Println(nums) // => [0 0]

	// declaring a slice using a slice literal

	type names []string

	friends := names{"Dan", "Maria"}

	fmt.Println(friends)

	// getting an element from the slice

	x := numbers[0]

	fmt.Println("x is", x) // => x is 2

	// modifying an element of the slice

	numbers[1] = 200

	fmt.Printf("%#v\n", numbers) // => []int{2, 200, 4, 5}

	// iterating over a slice

	for index, value := range numbers {

		fmt.Printf("index: %v, value: %v\n", index, value)

	}

	//iterating over a slice (C/C++, Java Style)

	for i := 0; i < len(numbers); i++ {

		fmt.Printf("index: %v, value: %v\n", i, numbers[i])

	}

	// slices with the same element type can be assigned to each other

	var n []int

	n = numbers

	_ = n

	//** COMPARING SLICES **//

	// a Go slice can only be compared to nil

	// uninitialized slice, equal to nil

	var nn []int

	fmt.Println(nn == nil) // true

	// empty slice but initialized, not equal to nil

	mm := []int{}

	fmt.Println(mm == nil) //false

	// we can not compare slices using the equal (=) operator

	// fmt.Println(nn == mm) //error -> slice can only be compared to nil

	// to compare 2 slices use a for loop to iterate over the slices and compare element by element

	// it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	var eq bool = true

	if len(a) != len(b) {

		eq = false

	}

	for i, valueA := range a {

		if valueA != b[i] {

			eq = false // don't check further, break!

			break

		}

	}

	if eq {

		fmt.Println("a and b slices are equal")

	} else {

		fmt.Println("a and b slices are not equal")

	}

	// arrays, slices and strings are sliceable
	// slicing doesn't modify the array or the slice, it returns a new one

	// declaring an [5]int array
	a := [5]int{1, 2, 3, 4, 5}

	// a slice expression is formed by specifying a start or a low bound and a stop or high bound like  a[start:stop].
	// this selects a range of elements which includes the element at index start, but excludes the element at index stop.

	// slicing an array returns a slice, not an array
	b := a[1:3]                                 // 1 is called start (included), 3 is called stop (excluded)
	fmt.Printf("Type: %T , Value: %#v\n", b, b) // => Type: []int , Value: []int{2, 3}

	// declaring a slice
	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:3]   //start included, stop excluded
	fmt.Println(s2) //[2 3]

	//for convenience, any of the indexis may be omitted.
	// a missing low index defaults to zero; a missing high index defaults to the length of the sliced operand.
	fmt.Println(s1[2:])       // => [3 4 5 6], same as s1[2:len(s1)]
	fmt.Println(s1[:3])       // => [1 2 3], same as s1[0:3]
	fmt.Println(s1[:])        // => [1 2 3 4 5 6], same with s1[0:len(s1)]
	fmt.Println(s1[:len(s1)]) // => => [1 2 3 4 5 6], returns the entire slice
	// fmt.Println(s1[:45])   //panic: runtime error: slice bounds out of range

	s1 = append(s1[:4], 100) // adds 100 after index 4 (excluded)
	fmt.Println(s1)          // -> [1 2 3 4 100]

	s1 = append(s1[:4], 200) // overwrites the last element
	fmt.Println(s1)          // -> [1 2 3 4 200]

	// a slice expression doesn't create a new backing array. The original and the returned slice are connected!
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3] //s3, s4 share the same backing array with s1

	s3[1] = 600     // modifying the backing array so s1, s3 and s4 are in fact modified!!
	fmt.Println(s1) // -> [10 600 30 40 50]
	fmt.Println(s4) // -> [600 30]

	// when a slice is created by slicing an array, that array becomes the backing array of the new slice.
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2                       // modifying the array
	fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

	// append() function creates a complete new slice from an existing slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	// newCars doesn't share the same backing array with cars
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"                              // only cars is modified
	fmt.Println("cars:", cars, "newCars:", newCars) // => cars: [Nissan Honda Audi Range Rover] newCars: [Ford Honda

	numbers := []int{2, 3}

	// append() returns a new slice after appending a value to its end
	numbers = append(numbers, 10)
	fmt.Println(numbers) //-> [2 3 10]

	// appending more elements at once
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers) //-> [2 3 10 20 30 40]

	// appending all elements of a slice to another slice
	n := []int{100, 200, 300}
	numbers = append(numbers, n...) // ... is the ellipsis operator
	fmt.Println(numbers)            // -> [2 3 10 20 30 40 100 200 300]

	s1 := []int{10, 20, 30, 40, 50}
	newSlice := s1[0:3]

	/*
		But, keep in mind that slice uses array in the backend.
		Slice by itself doesn’t store any data. Think of slice like a reference to an array. All it does is to describe part of the underlying array
	*/

	s1 := []int{10, 20, 30, 40, 50}
	newSlice := s1[0:3]
	fmt.Println(len(newSlice), cap(newSlice)) // The len is 3 and the cap is 5

	fmt.Printf("%#v %#v\n", &s1[0], &newSlice[0]) // -> (*int)(0xc0000da060) (*int)(0xc0000da060)

	newSlice = s1[2:5]
	fmt.Println(len(newSlice), cap(newSlice)) // The len is 3 and the cap is 3 there are only 3 items in the backend array

	// is because the backend array is the first element of the slice in the backend array; that's why capacity is 3 not 5
	// SLICING CREATES A NEW SLICE HEADER BUT THE BACKEND ARRAY IS THE SAME

	// To show a example that the backend array is the same

	fmt.Printf("%#v %#v \n", &s1[0], &newSlice[0]) // -> (*int)(0xc0000da060) (*int)(0xc0000da070)
	newSlice[0] = 500
	fmt.Println(s1) // -> [10 20 500 40 50]

	fmt.Printf("%#v %#v \n", &s1[0], &newSlice[0])

	// Slicing operations are cheaper then array operations her are a example where i show how to check the memory size

	array := [6]int{1, 2, 3, 4, 5}
	slice := []string{"1"}

	fmt.Printf("%#v %d\n", &slice[0], cap(slice))

	fmt.Printf("array's size in bytes: %d \n", unsafe.Sizeof(array)) // -> array's size in bytes: 40 is because 1 element is 8 bits / 1 byte
	fmt.Printf("slice's size in bytes: %d \n", unsafe.Sizeof(slice)) // -> slice's size in bytes: 24

	// Slicing operations are cheaper then array operations her are a example where i show how to check the memory size

	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	fmt.Printf("array's size in bytes: %d \n", unsafe.Sizeof(array)) // -> array's size in bytes: 40
	fmt.Printf("slice's size in bytes: %d \n", unsafe.Sizeof(slice)) // -> slice's size in bytes: 24

	//** Slice's Length and Capacity **//

	// The append() function creates a new backing array with a larger capacity

	nums := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 1, Capacity: 1

	nums = append(nums, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums = append(nums, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 5, Capacity: 8

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
	// if the slices don't have the same no of elements, it copies the minimum of length of the slices
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn) // => [10 20 30] [10 20 30] 3

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(years[:3], years[len(years)-3:]...) // it take the first and the last

	fmt.Println(newYears) // => [2000 2001 2002 2008 2009 2010]

}
