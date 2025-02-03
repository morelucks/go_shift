package main

import (
	// "crypto/rand"
	"fmt"
	"math/rand"
	"sync"
	"time"
	// "go/format"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main()  {
	dataChan:=make(chan int)
	go func ()  {
		// background threada
		wg:= sync.WaitGroup{}
		for i:=0; i<=100; i ++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <-result
			}()
			
		}
		wg.Wait()
		close(dataChan) // i.e after sending data to the channel you have to close it, if not it will result to deadlock
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