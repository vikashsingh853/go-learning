package main

import "fmt"

type Task struct{
	Title string
	Description string
	Completed bool
	id int
}

func (t *Task) markCompleted(){
	t.Completed=true
}

func (t *Task) display(){
	fmt.Printf("ID: %d, Title: %s, Description: %s, isCompleted: %t\n", t.id, t.Title, t.Description,t.Completed)
}


func main() {
	tasks:= []Task{}
	fmt.Println("TODO: Implement the main function")
	fmt.Println("Enter 1 to add task, 2 to mark task as completed, 3 to display tasks, 4 to delete task, 5 to exit:")

	for{
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice{
		case 1:
			var title string
			var description string
			fmt.Print("Enter task title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter task description: ")
			fmt.Scanln(&description)
			uniqueId:=0
			if len(tasks)==0{
				uniqueId=1
			}else{
				uniqueId=tasks[len(tasks)-1].id+1
			}
			task:=Task{Title:title,Description:description,Completed:false,id:uniqueId}

			tasks=append(tasks,task)
			fmt.Println("Task added successfully.")
		case 2:
			var id int
			fmt.Print("Enter task ID to mark as completed: ")
			fmt.Scanln(&id)

			for i :=range tasks{
				if tasks[i].id==id{
					tasks[i].markCompleted()
					fmt.Println("Task marked as completed.")
					break
				}
			}

		case 3: 
		 for _, task := range tasks{
			task.display() 
		}

		case 4:
			var id int
			fmt.Print("Enter task ID to delete: ")
			fmt.Scanln(&id)

			for i:=range tasks{
				if tasks[i].id==id{
					tasks=append(tasks[:i],tasks[i+1:]...)
					fmt.Println("Task deleted successfully.")
					break
				}
			}

			case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
			
		}
	}
}