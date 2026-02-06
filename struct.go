package main

import "fmt"

type User struct{
	name string
	age int
}

func (u User) display(){
	fmt.Printf("Name: %s, Age: %d\n", u.name, u.age)
}

func main(){
	users := []User{}

	fmt.Println("Enter 1 to add user, 2 to display users, 3 to exit:")

	for{
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice{
		case 1:
			var name string
			var age int
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter age: ")
			fmt.Scanln(&age)
			user :=User{name:name, age:age}
			users=append(users,user)
			fmt.Println("User added successfully.")
		case 2:
			for _, user := range users{
				user.display()
			}
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice! Please try again.")
			
		}
	}

}