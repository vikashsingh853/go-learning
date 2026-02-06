package main

import (
	"fmt"
	"myapp/calculator"
	"myapp/mathutils"
)

func main() {
	fmt.Println("Hello, World!")
	add := calculator.Add(10, 20)
	subtract := calculator.Subtract(20, 10)

	fmt.Println("Addition:", add)
	fmt.Println("Subtraction:", subtract)

	multiply := mathutils.Multiply(10, 20)
	divide := mathutils.Divide(20, 10)

	fmt.Println("Multiplication:", multiply)
	fmt.Println("Division:", divide)
}
