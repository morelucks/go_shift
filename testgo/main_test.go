package main

import (
	"testing"
	
)
func TestPlayer (t *testing.T){
	t.Log("this test is runing...")
	expected:=Player{
		name:"lucky",
		hp: 10,

	}
	have:=Player{
		name: "lucky",
		hp: 102,
	}

	// if have.hp != expected.hp && have.name!=expected.name{
	// 	t.Errorf("expect %v and got %v", expected, have)
	// }

}