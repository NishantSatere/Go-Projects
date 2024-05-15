package main

import (
	// "crypto/rand"
	"fmt"
	// "math/rand"
	// "sync"
	// "time"
	"sync"
)

// func DoWork() int {
// 	time.Sleep(time.Second)
// 	return rand.Intn(100)
// }

//	func attak(target string, attacked chan bool) {
//		time.Sleep(time.Second)
//		fmt.Println("Attacking", target)
//		attacked <- true
//	}

//	func throw(channel chan string) {
//		n := 10
//		for i := 0; i < n; i++ {
//			score := rand.Intn(10)
//			channel <- fmt.Sprintf("The score is %d", score)
//		}
//		close(channel)
//	}

// **************** Select Statement ****************
// func rouglyFair() {
// 	ninja1 := make(chan string)
// 	close(ninja1)
// 	ninja2 := make(chan string)
// 	close(ninja2)

// 	ninja1cnt := 0
// 	ninja2cnt := 0

// 	for i := 0; i < 100; i++ {
// 		select {
// 		case msg1 := <-ninja1:
// 			fmt.Println(msg1)
// 			ninja1cnt++
// 		case msg2 := <-ninja2:
// 			fmt.Println(msg2)
// 			ninja2cnt++
// 		}
// 	}

// 	println("ninja1cnt", ninja1cnt)
// 	println("ninja2cnt", ninja2cnt)
// }

// func passmessage(channel chan string, message string) {
// 	channel <- message
// }

// Alwqays use a pointer to the WaitGroup

// ***************** WaitGroup *****************
// func attakVillan(villan string, beeper *sync.WaitGroup) {
// 	fmt.Println("Killed", villan)
// 	beeper.Done()
// }

func main() {

	// ***************** Mutex *****************

	// ************************ WaitGroup ************************
	// it is an alternative to synchedronize goroutines

	// createing a waitgroup
	// var beeper sync.WaitGroup

	// evilNinjas := []string{"Tommy", "Jerry", "Spike"}

	// villan := "Tommy"

	// adding a goroutine to the waitgroup
	// beeper.Add(len(evilNinjas))

	// for _, ninja := range evilNinjas {
	// 	// pass by reference
	// 	go attakVillan(ninja, &beeper)
	// }

	// go attakVillan(villan, &beeper)
	// wait for the goroutine to finish
	// beeper.Wait()
	// fmt.Println("Attack Finished")

	// ninja1 := make(chan string)
	// ninja2 := make(chan string)

	// go passmessage(ninja1, "ninja1")
	// go passmessage(ninja2, "ninja2")
	// rouglyFair()

	// channel := make(chan string)
	// go throw(channel)

	// ******************* Range Over Channel *******************
	// for {
	// 	mesg, open := <-channel

	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(mesg)
	// }

	// now := time.Now()

	// ***************** defer statement *****************
	// defer func() {
	// 	fmt.Println(time.Since(now))
	// }()

	// smokeSignal := make(chan bool)

	// evilNinja := "tommy"
	// go attak(evilNinja, smokeSignal)
	// fmt.Println(<-smokeSignal)
	// time.Sleep(time.Second * 2)

	// Create a channel
	// dataChannel := make(chan int)

	// **************************Create a goroutine to send data to the channel

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
