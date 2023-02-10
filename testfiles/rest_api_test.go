package testfiles

import (
	"testing"
	"os"
	"github.com/19chonm/461_1_23/api"
)

var inputUrl string = "https://github.com/qiangxue/go-rest-api"
// {"license":{"key":"mit","name":"MIT License","url":"https://api.github.com/licenses/mit"}}

// Expected Values for Tests
var correctUser string = ""
var correctRepo string = ""
// var correctToken string = ""
// var correctOk bool = true
correctToken, correctOk := os.LookupEnv("GITHUB_TOKEN")

var badUser string = ""
var badRepo string = ""
var badToken string = ""
var badOk bool = false
// Input URL Tests
func TestGoodInput(t *testing.T) {
	user, repo, token, ok := validateInput(inputUrl)
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
	user, repo, token, ok :=  validateInput(inputUrl)
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

// ADD: Correctness Parameters, 
// Num Contributors, 
// License Compatibility, 
// Response Time, etc.


// // Rating Channel
// func TestGoodRatingChannel(t *testing.T) {

// }

// func TestBadRatingChannel(t *testing.T) {

// }
// // Input GitHub Token Tests
// func TestGoodToken(t *testing.T) {

// }

// func TestBadToken(t *testing.T) {

// }
