package testfiles

import (
	"testing"

	"github.com/19chonm/461_1_23/fileio"
	"github.com/19chonm/461_1_23/worker"
)

/*
Test to make sure rating channel is filled up
*/

// func TestWorkerRunsTask(t *testing.T) {

// 	//url_ch := make(chan string, rating_ch_size)

// }

func TestWorkerGoodInput(t *testing.T) {
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	worker.RunTask("https://github.com/expressjs/express", rating_ch)
	if len(rating_ch) != 1 {
		t.Errorf("rating channel was not updated")
	}

}

func TestWorkerBadInput(t *testing.T) {
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	worker.RunTask("https://badurl.com/blabla/test", rating_ch)
	if len(rating_ch) != 0 {
		t.Errorf("rating channel should not have been updated")
	}
}
