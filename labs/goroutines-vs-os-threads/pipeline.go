
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	mu     sync.Mutex
	fil os.File
	sinceStart time.Time
)

func writeTo(line int){
	mu.Lock()
	seconds := time.Since(sinceStart)
	_, err := fil.WriteString(strconv.Itoa(line) + "--" + seconds.String() +  "\n")
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	mu.Unlock()
}

func go1 (in <-chan int){
	x := <-in
	writeTo((x))
	out := make(chan int)


	go go1(out)
	out <- 1 + x
	for {
		time.Sleep(time.Millisecond* 10)
	}
}

func main () {
	runtime.GOMAXPROCS(1)
	fmt.Println("START")
	file_, err := os.Create("report.txt")
	if err != nil {

		fmt.Println(err)

		os.Exit(0)

	}
	fil = *file_
	ch1 := make(chan int)
	sinceStart = time.Now()
	go go1(ch1)
	ch1 <- 1
}
