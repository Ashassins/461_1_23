package testfiles

import (
	"testing"

	"github.com/19chonm/461_1_23/fileio"
	"github.com/19chonm/461_1_23/worker"
)

var url_ch_size = 100
var rating_ch_size = 5

// will check for activity on rating channel being filled up
// func TestManagerStartsWorkers(t *testing.T) {

// 	url_ch := make(chan string, url_ch_size)

// }

// testing for rating channel and workers
func TestManagerGoodInput(t *testing.T) {
	url_ch := make(chan string, url_ch_size)
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	url_ch <- "https://github.com/expressjs/express"
	url_ch <- "https://github.com/facebook/react"
	url_ch <- "https://github.com/request/request"

	worker.StartWorkers(url_ch, rating_ch)

	//testing logic
}

func TestManagerBadInput(t *testing.T) {
	url_ch := make(chan string, url_ch_size)
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	url_ch <- "https://github.com/facebook/react"
	url_ch <- "https://badurl.com/blabla/test"
	url_ch <- "https://www.npmjs.com/package/react"

	worker.StartWorkers(url_ch, rating_ch)

	//testing logic
}
