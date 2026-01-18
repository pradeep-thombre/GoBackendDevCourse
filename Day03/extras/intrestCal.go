package extras

import "fmt"

// calculate simple intrest
func CalculateSimpleIntrest(principal int, rateOfIntrest float64, time int) {
	var simpleIntrest = float64(principal) * rateOfIntrest * float64(time) / 100
	// var simpleIntrest = principal * int(rateOfIntrest) * time / 100
	fmt.Println("Simple Intrest is", simpleIntrest)
}

// rate of Intrest = 7.5 --> 7
// principal = 2000000.00
// time =5.00
