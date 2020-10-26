package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func sum(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total // send the total value to our channel name 'c'
}

func sum1(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total // send the total value to our channel name 'c'
	close(c)
}

func main() {
	s := []int{2, 4, 6, 12, 1000}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive data from the sum function

	fmt.Println(x, y, x+y)

	c1 := make(chan int, 10) // Buffer channel (with capacity)
	go sum1(s, c1)
	for i := range c1 {
		fmt.Println(i)
	}

	// OS LEVEL SIGNAL CHANNEL RECEIVER
	var wait time.Duration
	// Make graceful shutdown almost immediately at 1 Millisecond due to user may encounter errors
	flag.DurationVar(&wait, "graceful-timeout", time.Millisecond*1, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	c2 := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c2, os.Interrupt)

	// Block until we receive our signal.
	<-c2

	// Create a deadline to wait for.
	_, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	fmt.Println("the user press CTRL + C to exit our Go's app")
	os.Exit(0)
}
