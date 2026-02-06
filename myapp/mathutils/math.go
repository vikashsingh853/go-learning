package mathutils

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	if b != 0 {
		return a / b
	}
	return 0 // Return 0 for division by zero
}