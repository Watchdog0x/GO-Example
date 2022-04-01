package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// if condition_that_evaluates_to_boolean{
	//      perform action1
	// }else if condition_that_evaluates_to_boolean{
	//      perform action2
	// }else{
	//      perform action3
	// }

	price, inStock := 100, true

	if price >= 80 { // parenthesis are no required to enclose the testing condition
		fmt.Println("Too Expensive")
	}

	if price <= 100 && inStock { //the same with: if price <= 100 && inStock { }
		fmt.Println("Buy it!")
	}

	// In Go there is not such a thing like the Truthiness of a variable.
	// Error:
	// if price {
	//  fmt.Println("We have price!")
	// }

	// only one if branch will be executed
	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else { //executed only once if all the if branches are false (it's optional)
		fmt.Println("It's Expensive!")
	}

	// converting string to int:
	i, err := strconv.Atoi("45")

	// error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)

	}

	// simple (short) statement ->  the same effect as the above code
	// i and err are variables scoped to the if statement only
	if i, err := strconv.Atoi("34"); err == nil {
		fmt.Println("No error. i is ", i)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer! Error:", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*0.621)
	}

}
