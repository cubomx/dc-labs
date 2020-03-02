package main

import (
	"flag"
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
	"time"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

type Link struct {
	depth int
	used bool
}


func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

var data = make([] string, 10)

func work (link string, actualDepth int, seen map[string]*Link, worklist chan [] string) {
	go func(link string) {
		var data = crawl(link)
		for _, child := range data{
			seen[child] = &Link{actualDepth + 1, false}
		}
		if (actualDepth == 0){
			fmt.Println(len(data))
		}
		worklist <- data
	}(link)
}

func main() {

	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() {
		worklist <- os.Args[1:] }()
	depth := flag.Int("depth", 1, "limit of depth crawling")
	flag.Parse()
	// Crawl the web concurrently.
	seen := make(map[string]*Link)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if seen[link] == nil {

				info := &Link{1, true}
				seen[link] = info
				fmt.Println("Level:   ", seen[link].depth)
				n++
				work(link, info.depth, seen, worklist)

			} else if !seen[link].used{
				seen[link].used = true
				if seen[link].depth <= *depth {
					fmt.Println("Level:   ", seen[link].depth)
					n++
					work(link, seen[link].depth, seen, worklist)
				}
			} else {
				n--
			}
			time.Sleep(time.Millisecond * 100)
		}
	}
}
