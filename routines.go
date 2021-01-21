package main

import (
	"fmt"
	"time"
)

func main() {

	go hello()
	//time.Sleep(100 * time.Millisecond)
	go numbers()
	go alphabets()
	time.Sleep(3 * time.Second)
	fmt.Println("\nThis is the main function")
}

func hello() {
	fmt.Println("This is the hello function")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
