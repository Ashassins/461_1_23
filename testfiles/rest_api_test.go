package testfiles

import (
	"testing"
	"os"
	api "github.com/19chonm/461_1_23/api"
)
// {"license":{"key":"mit","name":"MIT License","url":"https://api.github.com/licenses/mit"}}

// Input URL Tests
func TestGoodInput(t *testing.T) {
	var goodInputUrl string = "https://github.com/qiangxue/go-rest-api"
	var correctUser string = ""
	var correctRepo string = ""
	correctToken, correctOk := os.LookupEnv("GITHUB_TOKEN")

	user, repo, token, ok := api.ValidateInput(goodInputUrl)
	if user != "" {
		t.Errorf("user got: %s, want: %s", user, correctUser)
	} 
	if repo != "" {
		t.Errorf("repo got: %s, want: %s", user, correctRepo)
	}
	if token != "" {
		t.Errorf("token got: %s, want: %s", user, correctToken)
	}
	if ok != "" {
		t.Errorf("ok got: %b, want: %b", ok, correctOk)
	}
}

func TestBadInput(t *testing.T) {
	var badInputUrl string = ""
	var badUser string = "badUser"
	var badRepo string = "badRepo"
	var badToken string = "badToken"
	var badOk bool = false

	user, repo, token, ok := api.ValidateInput(badInputUrl)
	if user != "" {
		t.Errorf("user got: %s, want: %s", user, badUser)
	} 
	if repo != "" {
		t.Errorf("repo got: %s, want: %s", user, badRepo)
	}
	if token != "" {
		t.Errorf("token got: %s, want: %s", user, badToken)
	}
	if ok != "" {
		t.Errorf("ok got: %b, want: %b", ok, badOk)
	}
}

// Scoring Algorithm Tests
// Overall Scores
func TestGoodScoringAlgorithm(t *testing.T) {

}

func TestBadScoringAlgorithm(t *testing.T) {

}

// Rating Channel
func TestGoodRatingChannel(t *testing.T) {
	// var thing = false
	// if thing != false {
	// 	t.Errorf(" ")
	// }
}

func TestBadRatingChannel(t *testing.T) {
	// var thing = false
	// if thing != false {
	// 	t.Errorf(" ")
	// }
}

// ADD: Correctness Parameters, 
// Num Contributors, 
// License Compatibility, 
// Response Time, etc.

// // Input GitHub Token Tests
// func TestGoodToken(t *testing.T) {

// }

// func TestBadToken(t *testing.T) {

// }
