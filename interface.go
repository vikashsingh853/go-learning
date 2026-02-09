package main

import "fmt"

type Speaker interface {
	Speak() string
	Dance() string
}

type Dog struct {}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Dance() string {
	return "The dog is dancing!"
}

func makeSpeak(s Speaker) {
	fmt.Println(s.Dance())
	fmt.Println(s.Speak())
}

func main(){
	d:=Dog{}
	makeSpeak(d)
}