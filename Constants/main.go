package main

func main() {
	const days int = 7
	const n, m int = 4, 5
	const g, g1 = 5, 6

	const (
		min  = -500
		min1 = 500
		min2 = 402
	)

	// Constants rules
	//1.you can not change a constant
	const temp int = 100
	// temp = 50

	//2. You can not initiate a constant at runtime
	// Because const belong to compile time
	// const power = math.Pow(2, 3)

	//3. you can not use a variable to initialize a constant
	//t := 5
	//const tc = t

	//4. You can use len function in constant because is workign on complie time
	const le = len("Hello")

	// untypede constant
	const a float64 = 5.1 // this is a typeed constant

	const b = 6.7 // this is untyped constant

	// in intyped constant it will assigne the type by it self so you can do this

	const x = 5
	const y = 2.2 * x

	// this will give a error becuase we can use const that way
	const c int = 5
	const d float32 = 2.3

	// const f float32 = d * c
	// const f = c * d

	// but this will work becuase we did use untype for x and y

	const q = x * y         // this type will be float64
	const f float32 = x * y // this type will be float32
}
