package main

import (
	"fmt"
	// "go/format"
)

func main()  {
	dataChan:=make(chan int)
	go func ()  {
		// background thread
		for i:=0; i<100; i ++ {
			dataChan <-i
		}
	}()
		//  main thread
	for n:= range dataChan{
	
		fmt.Printf("n =%d\n", n)

	}

	// dataChan <- 78 
	// // dataChan <- 723


	// n:= <- dataChan

	// fmt.Printf("n =%d\n", n)
}