package main

import (
	"fmt"
)

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel is nil, must be initialized before use")
		a = make(chan int)
		fmt.Printf("Channel type is %T\n", a)
	}

	b := make(chan bool)
	go hello(b)
	<-b // Serve ad indicare che voglio leggere qualcosa dal canale b. E'una richiesta bloccante. Senza di questa, la goroutine hello non viene eseguita
	fmt.Println("This is the main function")

	n := 255
	c := make(chan int)
	go SumSquareDigit(n, c)
	sum := <-c

	fmt.Println("The sum of the squares of the single digits is", sum)

	sendChan := make(chan<- int) // crea un canale in cui si puo' solo inviare (unidirezionale)
	go SendData(sendChan)
	//fmt.Println(<- sendChan) // non va bene, perche' stiamo ricevendo da un canale che e' definito come sender
	convertChan := make(chan int)
	go SendData(convertChan) // anche se convertChan e' bidirezionale, posso passargli l'unidirezionale. Il viceversa NON e' consentito!
	fmt.Println(<-convertChan)

	consumerChan := make(chan int)

	go producer(consumerChan)
	for {
		v, ok := <-consumerChan
		if !ok {
			break
		}
		fmt.Println("Received ", v, ok)

	}

	go producer(c)
	for e := range c {          // I can use a for loop to iterate over the elements in the channel
		fmt.Println("Received ", e)
	}

}

func hello(done chan bool) {
	fmt.Println("This is the hello function")
	done <- true // Scrivi true nel canale. E' fondamentale, altrimenti avremmo un deadlock
	fmt.Println("Channel written")
}

func SumSquareDigit(n int, sumsquare chan int) {
	sum := 0
	for n != 0 {
		digit := n % 10
		sum += digit * digit
		n = n / 10
	}
	sumsquare <- sum
}

func SendData(sendchan chan<- int) {
	sendchan <- 10
}

func producer(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c) // I can close the channel when I am done producing
}
