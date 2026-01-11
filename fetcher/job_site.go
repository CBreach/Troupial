package fetcher

import (
	//"github.com/gocolly/colly/v2"
	"fmt"
	"net/url"

	"github.com/gocolly/colly/v2"
)

/*
.- helper function to parse and detect the website in question, to later decide what parser to use
.- it returns an error if the url passed is not valid
*/
func urlDetector(rawUrl string) (string, error) {
	// this will be a very simplified version, but it'll do for now
	// a domain must begin with the (http or https) protocol and must end with a top level domain (.com)
	//TODO: add logic for urls containing a country domain such as "goodle.com.uk"
	u, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return "invalid URL", err
	}
	fmt.Println(u.Host)
	return u.Host, nil
}

// Initializes a new fetcher object that will be used to make the web requests through colly
func NewFetcher(rawUrl string) (*Fetcher, error) {
	host, err := urlDetector(rawUrl)
	if err != nil {
		return nil, err
	}

	// this is a bit of a weird syntax but it basically initializes a new fetcher and returns its address along with a null error
	return &Fetcher{
		rawUrl:    rawUrl,
		Host:      host,
		Collector: colly.NewCollector(),
	}, nil

}
