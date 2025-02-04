package main

import (
	"fmt"
	// "go/version"
)
var name = "lucky"

type Player struct{
	name string
	age int
	attackPower int

}
func getName(player Player) string{ // this could also be achieve using function receive
	return player.name
}

// function receive

func (player Player)getAge() int {
	return player.age
}
func main(){
	fmt.Println(name)
	version:=2
	version2:=5
	fmt.Println(version)
	fmt.Println(version2)
	player := Player{
		name: "luckify",
		age: 3,
		attackPower: 100,

	}
	getn := getName(player)
	fmt.Printf("this is player %v\n", player)
	fmt.Printf(" the player name is %+v\n", getn)

	fmt.Printf("this is the player's age %+v\n", player.getAge())
}