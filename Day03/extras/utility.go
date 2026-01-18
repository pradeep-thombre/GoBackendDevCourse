package extras

import "fmt"

var randomVariable = 42
var Marks = "25.0"

func Addition(num1 int, num2 int) {
	fmt.Println("Inside Addition Function")
	var addition = num1 + num2
	fmt.Println("Addition is", addition)
}

func substraction() {}

func Multiply(num1 float64, num2 int) { // 11.2 * 3 -- 33

	fmt.Println("Inside Multiply Function")
	var multiplication = num1 * float64(num2)

	fmt.Println("multiplication is", multiplication)

}

func Division(num1 float32, num2 float32) {
	fmt.Println("Inside Division Function")
	var division = num1 / num2
	fmt.Println("Division is", division)
}

// Simple intrest = (P * R * T) / 100
