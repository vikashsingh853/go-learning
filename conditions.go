package main

import "fmt"

func main(){
	var guessBumber int
	secret :=7

	fmt.Print("Guess the number between 1 and 10: ")
	fmt.Scanln(&guessBumber)

	// if guessBumber==secret{
	// 	fmt.Println("Congratulations! You guessed the number.")
	// }else if guessBumber<secret{
	// 	fmt.Println("Too low! Try again.")
	// }else{
	// 	fmt.Println("Too high! Try again.")
	// }

	// switch guessBumber{
	// 	case secret:
	// 		fmt.Println("Congratulations! You guessed the number.")
	// 	case 1,2,3,4,5,6:
	// 		fmt.Println("Too low! Try again.")
	// 	case 8,9,10:
	// 		fmt.Println("Too high! Try again.")
	// 	default:
	// 		fmt.Println("Out of range! Please guess a number between 1 and 10.")
	// }

	for {
		if guessBumber==secret{
			fmt.Println("Congratulations! You guessed the number.")
			break
		}
		fmt.Print("Guess the number between 1 and 10: ")
		fmt.Scanln(&guessBumber)
	}
}