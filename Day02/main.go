package main

import (
	"fmt"
	"strconv"
)

// randomVariable := 42 // invalid outside function

var marks float64 // takes default value 0

func main() {
	// fmt.Println("Hello, World!")
	var name string = "Abhay" // declaration of string with initialization
	fmt.Println("My name is", name)

	// var surname string // declaration of string without initialization

	var location = "Pune" // declaration of string with initialization
	fmt.Println("I live in", location)

	game := "Football" // short declaration of string with initialization
	fmt.Println("My favorite game is", game)

	fmt.Println("my marks is", marks)

	// var num uint = 100 // unsigned int it contains only positive values 0,123...
	var num int = 100 // signed int it contains both negative and positive values -123,0,123...
	fmt.Println("num is", num)

	var isHappy bool = false
	fmt.Println("Am I happy?", isHappy)

	var temperature float64 = 36.666666666666666666666666666666 // float64 is default type for floating point numbers
	fmt.Println("Temperature is", temperature)                  // precision up to 15-16 decimal places

	var tempInFarenhite float32 = 98.6666666666666666 // float32 precision up to 6-7 decimal places
	fmt.Println("Temperature in Farenhite is", tempInFarenhite)

	var num1 int = 100
	var num2 int = 3

	fmt.Println(num1 + num2) // addition
	fmt.Println(num1 - num2) // subtraction
	fmt.Println(num1 * num2) // multiplication
	fmt.Println(num1 / num2) // division
	fmt.Println(num1 % num2) // modulus operator gives remainder

	var numFloat1 float64 = 100.1
	var numFloat2 float64 = 3.5

	fmt.Println(numFloat1 + numFloat2) // addition
	fmt.Println(numFloat1 - numFloat2) // subtraction
	fmt.Println(numFloat1 * numFloat2) // multiplication
	fmt.Println(numFloat1 / numFloat2) // division
	// fmt.Println(numFloat % numFloat) // modulus operator gives remainder (not valid for float)

	var num3 int = 10
	fmt.Println(float64(num3) * 3.51222) // conversion of int to float64

	var num4 float64 = 10.99
	fmt.Println(int(num4) * 3) // conversion of float64 to int (decimal part will be truncated)

	fmt.Println("int--> string", strconv.Itoa(num3)) // conversion of int to string

	var strNum string = "1234"
	var strInt, _ = strconv.Atoi(strNum) // conversion of string to int
	fmt.Println("string--> int", strInt)
}

// zero values
// int     -> 0
// float   -> 0.0
// string  -> ""
// bool    -> false
// pointers -> nil
