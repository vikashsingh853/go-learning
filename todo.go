package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func saveStask(tasks []Task){
  file,err:= os.Create("tasks.txt")

  if err!=nil{
	fmt.Println("Error creating file:", err)
	return
  }

  defer file.Close()

  for _,task:=range tasks{
	line:=fmt.Sprintf("%d,%s,%s,%t\n", task.id, task.Title, task.Description, task.Completed)
	_,err:=file.WriteString(line)

	if err!=nil{
		fmt.Println("Error writing to file:", err)
		return
	}
  }

}

func loadTasks() []Task{
	data,err:=os.ReadFile("tasks.txt")
	
	if err!=nil{
		fmt.Println("Error reading file:", err)
		return []Task{}
	}

	lines:=strings.Split(string(data), "\n")
	tasks:=[]Task{}

	for _,line:=range lines{
		if line==""{
			continue
		}

		parts:=strings.Split(line, ",")

		if len(parts)!=4{
			fmt.Println("Invalid task format:", line)
			continue
		}

		id,err:=strconv.Atoi(parts[0])
		if err!=nil{
			fmt.Println("Invalid task ID:", parts[0])
			continue
		}

		title:=parts[1]
		description:=parts[2]
		completed,err:=strconv.ParseBool(parts[3])
		if err!=nil{
			fmt.Println("Invalid completed status:", parts[3])
			continue
		}

		task:=Task{Title:title, Description:description, Completed:completed, id:id}
		tasks=append(tasks,task)
	}

	return tasks
}


func main() {
	
	tasks:= loadTasks()
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
			saveStask(tasks)
			fmt.Println("Task added successfully.")
		case 2:
			var id int
			var found bool
			fmt.Print("Enter task ID to mark as completed: ")
			fmt.Scanln(&id)

			for i :=range tasks{
				if tasks[i].id==id{
					found=true
					tasks[i].markCompleted()
					fmt.Println("Task marked as completed.")
					break
				}
			}
			if !found{
				fmt.Println("Task not found with ID:", id)
			}else{
				saveStask(tasks)
			}

		case 3: 
		 for _, task := range tasks{
			task.display() 
		}

		case 4:
			var id int
			var found bool
			fmt.Print("Enter task ID to delete s: ")
			fmt.Scanln(&id)

			for i:=range tasks{
				if tasks[i].id==id{
					found=true
					tasks=append(tasks[:i],tasks[i+1:]...)
					fmt.Println("Task deleted successfully.")
					break
				}
			}
			if !found{
				fmt.Println("Task not found with ID:", id)
			}else{
				fmt.Println("Task deleted successfully.")
				saveStask(tasks)
			}
			case 5:
			fmt.Println("Exiting...")
			saveStask(tasks)
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}