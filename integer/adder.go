package integer

import "fmt"

/*
Add generates the sum of two numbers

Parameters:
  - `x` : the first number
  - `y` : the second number

Returns:
  - The sum of the two numbers
*/
func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
