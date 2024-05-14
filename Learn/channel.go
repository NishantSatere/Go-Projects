package main

import (
	// "crypto/rand"
	"fmt"
	"math/rand"
	// "sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func attak(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Attacking", target)
	attacked <- true
}
func main() {

	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)

	evilNinja := "tommy"
	go attak(evilNinja, smokeSignal)
	fmt.Println(<-smokeSignal)
	// time.Sleep(time.Second * 2)

	// Create a channel
	// dataChannel := make(chan int)

	// Create a goroutine to send data to the channel

	// go func() {
	// 	wg := sync.WaitGroup{}
	// 	for i := 0; i < 50; i++ {
	// 		wg.Add(1)
	// 		go func() {
	// 			defer wg.Done()
	// 			res := DoWork()
	// 			dataChannel <- res
	// 		}()
	// 	}
	// 	wg.Wait()
	// 	// if we don't close the channel, the for loop will keep waiting for more data and lead to deadlock
	// 	close(dataChannel)
	// }()

	// Read data from the channel
	// for n := range dataChannel {
	// 	fmt.Println(n)
	// }

	// We converted the channel to a buffered channel
	// dataChannel := make(chan int, 2)

	// Create a goroutine to send data to the channel
	// go func() {
	// dataChannel <- 42
	// dataChannel <- 43
	// // }()

	// Read data from the channel
	// n := <-dataChannel
	// fmt.Println(n)

	// n = <-dataChannel
	// fmt.Println(n)
}
