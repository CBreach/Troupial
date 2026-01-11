package fetcher

import (
	"log"
	"testing"
)

var tests = []struct {
	rawUrl   string
	cleanUrl string
	wantErr  bool
}{
	{
		"https://www.linkedin.com/jobs/collections/recommended/?currentJobId=3864453531",
		"www.linkedin.com",
		false,
	},
	{
		"https://www.indeed.com/jobs?q=&l=San+Diego%2C+CA&from=searchOnHP&vjk=e8f3a5ed795808b8",
		"www.indeed.com",
		false,
	},
	{
		"tmks/www.gatoVonis....com",
		"invalid URL",
		true,
	},
}

// tests behavior of the urlDetector unexported function
// takes raw URL and attempts to parse its host, if fails at doing so then it returns an error
func TestUrlDetector(t *testing.T) {
	for _, currTest := range tests {
		cleanUrl, err := urlDetector(currTest.rawUrl)
		if err == nil && currTest.wantErr {
			log.Printf("unmatching err outputs on invalid URL, expecting %t and got %t instead", currTest.wantErr, err == nil)
			t.Fail()

		} else if err != nil && !currTest.wantErr {
			log.Printf("unmatching err outputs on valid URL, expecting %t and got %t instead", currTest.wantErr, err == nil)
			t.Fail()
		}

		if cleanUrl != currTest.cleanUrl {
			log.Printf("Parsing error, expected %s and got %s instead", currTest.cleanUrl, cleanUrl)
			t.Fail()
		}

	}
}

func TestNewFetcher(t *testing.T) {
	for _, currTest := range tests {
		f, err := NewFetcher(currTest.rawUrl)
		if err != nil && !currTest.wantErr {
			log.Printf("error: %v", err)
			t.Fail()
		}
		if f == nil && !currTest.wantErr {
			log.Printf("error creating a fetcher, URL: %s likely Invalid", currTest.rawUrl)
		}
	}
}
