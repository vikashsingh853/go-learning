package main

import "fmt"

func libraryFunction(){
	fmt.Println("This is a library function.")
	students :=make(map[string]int)

	for{
		var choice int
		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice{
		case 1:
			var name string
			var age int
			fmt.Print("Enter name:")
			fmt.Scanln(&name)
			fmt.Print("Enter age:")
			fmt.Scanln(&age)
			students[name]=age
			fmt.Println("Student added successfully.")
		case 2:
			for name,age:=range students{
				fmt.Printf("Name: %s, Age: %d\n", name, age)
			}

		case 3:
			fmt.Println("Exiting...")
			return

			default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}

}

func main(){
	var nums [5]int

	nums[0]=10
	nums[1]=20
	nums[2]=30
	nums[3]=40
	nums[4]=50

	// fmt.Println("Array:", nums)
	// fmt.Println("First element:", nums[0])
	// fmt.Println("Length of array:", len(nums))

	// fmt.Println("Iterating over array:")

	// // for i:=0;i<len(nums);i++{
	// // 	fmt.Println(nums[i])
	// // }

	// names :=[]string{"John", "Alice", "Bob"}
	// fmt.Println("Names Array:", names)

	// for i,name:=range names{
	// fmt.Printf("Index: %d, Name: %s\n", i, name)
	// }

	// names = append(names, "Eve")
	// fmt.Println("Updated Names Array:", names)

	// // map
	// ages :=make(map[string]int)
	// ages["John"]=30
	// ages["Alice"]=25
	// ages["Bob"]=35

	// fmt.Println("Ages Map:", ages)
	// fmt.Println("John's age:", ages["John"])

	// delete(ages, "Alice")
	// fmt.Println("Updated Ages Map:", ages)

	// for name,age:=range ages{
	// 	fmt.Printf("%s is %d years old.\n", name, age)
	// }

	libraryFunction()
}