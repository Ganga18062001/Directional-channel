package main

import (
	"fmt"
	"sync"
)

//Using buffer channel data is send or receive
// buffer full : sender block
//buffer empty : receiver block

var ch chan int = make(chan int,5)

func Sender(gm *sync.WaitGroup){
	defer gm.Done()
	ch <- 10
	ch <- 11
	ch <- 12
	ch <- 13
	ch <- 14
	ch <- 15
	ch <- 20
	fmt.Println("Data send successfully")

}
func Receiver(gm *sync.WaitGroup){
	defer gm.Done()
	val := <- ch
	fmt.Println(val)

	// val = <- ch
	// fmt.Println(val)

	// val = <- ch
	// fmt.Println(val)

	// val = <- ch
	// fmt.Println(val)

	// val = <- ch
	// fmt.Println(val)

	// val = <- ch
	// fmt.Println(val)

	// val = <- ch
	// fmt.Println(val)
}

func main(){
	var gm sync.WaitGroup
	gm.Add(2)
	go Sender(&gm)
	go Receiver(&gm)
	gm.Wait()

}