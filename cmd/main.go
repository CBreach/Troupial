package main

import (
	"fmt"
	"log"

	"github.com/CBreach/Troupial/fetcher"
	//"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("PIpeline begins")
	f, err := fetcher.NewFetcher("https://www.linkedin.com/jobs/collections/recommended/?currentJobId=3864453531")
	if err != nil {
		log.Fatalf("Could not initialize fetcher, error: %v", err)
	}
	f.Fetch()
}
