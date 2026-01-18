package main

import (
	"day3/extras"
	"fmt"
)

func main() {
	fmt.Println("Welcome")
	fmt.Println(extras.Marks)

	extras.Marks = "55"
	fmt.Println("Updated Marks:", extras.Marks)
	// fmt.Println(extras.randomVariable)

	fmt.Println("Updated Marks:", extras.Marks)

	var age uint = 20

	fmt.Println("Age is", age)

	// age = -21

	fmt.Println("Updated Age is", age)

	var location string = ""

	fmt.Println("Location is", location)

	location = "Pune"

	fmt.Println("Updated Location is", location)

	// declaration byte

	var char byte = 'A' // byte is an alias for uint8
	fmt.Println("Character is", char)
	fmt.Println("character is", string(char))

	// declaration rune
	var runeVar rune = 'अ' // rune is an alias for int32
	fmt.Println("Rune is", runeVar)
	fmt.Println("rune is", string(runeVar))

	runeVar = '😀'

	fmt.Println("Updated Rune is", runeVar)
	fmt.Println("Updated rune is", string(runeVar))

	extras.Addition(10, 20)

	fmt.Println("End of Main")

	extras.Multiply(12.2, 10)
	extras.Multiply(20.5, 100)
	extras.Division(100, 3)

	var principal int = 2000000
	var rate float64 = 7.5
	var time int = 5

	extras.CalculateSimpleIntrest(principal, rate, time)
}

// Ques.  calculate Simple Intrest on loan amount of 20,00,000 at rate of 7.5% for 5 years
// assume time and pricipal will always be whole numbers

//AJXJW4561J
