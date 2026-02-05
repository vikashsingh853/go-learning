package main
import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var name string = "Go Developer"
	var age int = 5
	fmt.Println(name)
	fmt.Println(age)

	// short variable declaration
	country :="India"
	fmt.Println(country)

	// // input 
	// var nameInput string
	// var ageInput int
	// fmt.Print("Enter your name:")
	// fmt.Scanln(&nameInput)
	// fmt.Print("Enter your age:")
	// fmt.Scanln(&ageInput)
	// fmt.Println("Name:", nameInput)
	// fmt.Println("Age:", ageInput)

	userInput()
}

func userInput(){
	var name string
	var country string
	var language string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your country: ")
	fmt.Scanln(&country)

	fmt.Print("Enter your favourite programming language: ")
	fmt.Scanln(&language)

	fmt.Println("My name is ", name)
	fmt.Println("I am from ", country)
	fmt.Println("I Love ", language)
}