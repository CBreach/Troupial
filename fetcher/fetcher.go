package fetcher

import (
	"github.com/gocolly/colly/v2"
)

/*
	defines the basic struct of a fetcher
*/

type Fetcher struct {
	rawUrl    string
	Host      string
	Collector *colly.Collector
}
