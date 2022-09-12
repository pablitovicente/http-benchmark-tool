package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type ReqStat struct {
	StatusCode int
	ExecTime time.Duration
}

func makeRequest (url string, result chan ReqStat) {
	start := time.Now()
	// execute request
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	// calculate how long it took
  elapsed :=  time.Since(start)
	// Send the results over channel
	result <- ReqStat{
		StatusCode: res.StatusCode,
		ExecTime: elapsed,
	}
}

func main(){
	requestNumber := 10000

	results := make(chan ReqStat)
	for i := 0; i < requestNumber; i ++ {
		go makeRequest("http://localhost:3000/hello", results)
	}

	allResults := make([]ReqStat, 0)

	for result := range results {
		fmt.Printf(".")
		allResults = append(allResults, result)
		if(len(allResults) == requestNumber) {
			log.Println("We've got all results closing channel")
			close(results)
		}
	}

	log.Printf("%+v\n", allResults)
}