package main

import "fmt"

func add(a int,b int)int{
	return a+b
}

func substract(a int,b int)int{
	return a-b
}

func multiply(a int,b int)int{
	return a*b
}

func divide(a int,b int)int{
	 if b!=0 {
		return a/b
	}

	return 0
}

func main(){
	var num1 int
	var num2 int

	var operator string
	
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	for operator!="exit"{
		if operator=="+"{
			fmt.Println("Result: ", add(num1,num2))
		}else if operator=="-"{
			fmt.Println("Result: ", substract(num1,num2))
		}else if operator=="*"{
			fmt.Println("Result: ", multiply(num1,num2))
		}else if operator=="/"{
			if num2!=0{
				fmt.Println("Result: ", divide(num1,num2))
			}else{
				fmt.Println("Error: Division by zero is not allowed.")
			}
		}else{
			fmt.Println("Invalid operator! Please enter one of +, -, *, /.")
			break
		}
		
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
		fmt.Print("Enter operator (+, -, *, /): ")
		fmt.Scanln(&operator)
	}
}