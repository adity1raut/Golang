

package main 



/// This is a variadic function that takes a variable number of integer arguments
// and returns their product. It uses the `...` syntax to accept a slice of integers.
// Variadic functions are useful when you want to pass a variable number of arguments
// to a function without needing to define a fixed number of parameters.
// Variadic functions can be called with zero or more arguments, and they are often
// used for operations like summation, multiplication, or concatenation of strings.
// The function iterates over the provided integers, multiplying them together,
// and returns the final product. If no arguments are provided, it returns 1,
// which is the multiplicative identity (since multiplying by 1 does not change the value).
// Variadic functions are defined by placing `...` before the type in the function signature.
// This allows the function to accept any number of arguments of that type.
// The function can be called with any number of integer arguments, including none,
// and it will handle them appropriately.
// Variadic functions are a powerful feature in Go, allowing for flexible function calls
// and making it easier to work with collections of data without needing to define
// a specific number of parameters.
// This function multiplies a variable number of integers and returns the product.
// It demonstrates the use of variadic functions in Go, which allow for flexible
// function signatures that can accept a variable number of arguments.
// Variadic functions are defined by placing `...` before the type in the function signature.
// This allows the function to accept any number of arguments of that type.

func Multyply (nums ...int) int {
	result := 1 
	for _ , num := range nums {
		result *= num
	}
	return result
}

func main () {

	r := Multyply (2 ,3 , 5 ,7 , 9)
	println(r) // Output: 1890


	r = Multyply (2 ,3 , 5)
	println(r) // Output: 30
}