package main
import "fmt"

func swap(a *int,b *int){
	*a,*b=*b,*a
}

func main(){
	x:=10
	y:=200

	fmt.Printf("Before swap: x=%d, y=%d\n", x, y)

	swap(&x, &y)

	fmt.Printf("After swap: x=%d, y=%d\n", x, y)


}