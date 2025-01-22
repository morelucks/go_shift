package main

import (
	"fmt"
	// "os/user"
	// "os/user"
)
var z int 


func main() {
	x := 3;
	y := "new code"
	z=6;


  fmt.Println("Hello World!", x , y)
  fmt.Println(z)
  fmt.Println("the sum of x and y is : ", add(4,5))
	uname:="main user"
	updateName(uname)
  fmt.Println(uname)

  fmt.Println("the memory address of uname is : ", &uname)
//   fmt.Println("the memory address of uname is : ", &user)


mp:=&uname
fmt.Println("pointer mp", mp)
// deferencing a pointer
fmt.Println("dereference mp: ", *mp)
}


func add(x int, y int) int{
	return x+y
}
func updateName(user string){
	user="user update"
	fmt.Println("the memory address of X is : ", &user)

}
