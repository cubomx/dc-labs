package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)
var (
	quantityA int
	quantityB int
	file_ os.File
)

func in(in <-chan int, out chan<- int) {
	go count(time.Now())
	for {
		<-in
		out <- 1
		quantityA++
	}
}

func writeTo(line string){
	_, err := file_.WriteString(line)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

}

func count (start time.Time) {
	for {
		sum := quantityA+quantityB
		seconds := time.Since(start)
		if seconds >= time.Second{

			writeTo("Log report of communications between Goroutines\n")
			writeTo("--------------------------------------------------\n")
			writeTo("Time elapsed             : " + seconds.String() + "\n")
			writeTo("Amount of communications : " + strconv.Itoa(sum))
			os.Exit(0)
		}
	}
}

func out(in chan<- int, out <-chan int) {
	for {
		<-out
		in <- 1
		quantityB++
	}
}

func main() {
	file, err := os.Create("report.txt")
	file_ = *file
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	ping := make(chan int)
	pong := make(chan int)
	go in(ping, pong)
	go out(ping, pong)

	ping <- 1

	for {
		continue
	}
}