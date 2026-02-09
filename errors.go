package main

import (
	"errors"
	"fmt"
)

func divide(a,b int)(int,error){
 if b==0{
	return 0,errors.New("division by zero")
 }

 return a/b,nil
}

func main(){

	val,err :=divide(10,2)

	if err!=nil{
	fmt.Println("Error:",err)
	return
	}

	fmt.Println("Result:",val)

}