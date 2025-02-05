package main

import (
	"fmt"

	// "golang.org/x/text/number"
)

type NStorer interface{
    GetAll() ([]int,  error)
    Put(int) error
}

type MongoDB struct{

}

//implement mongodb

func (m MongoDB) GetAll() ([]int, error){
    return []int {1,3, 4}, nil
}

func (m MongoDB) Put(n int) error{
    fmt.Println("inserting ..", n)
    return nil
}
func main(){
	s := MongoDB{}
    numb, _:= s.GetAll()
    fmt.Println("Numbers: ", numb)
}