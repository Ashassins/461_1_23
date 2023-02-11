package testfiles

import (
	"testing"

	"github.com/19chonm/461_1_23/api"
)

// Since numbers would vary frequently, checking is done so that a valid response
// is received. Not for correctness
func TestGoodGithubNames(t *testing.T) {
	goodOwnerName := "facebook"
	goodRepoName := "react"
	m1, m2, m3, err := api.GetCorrectnessFactors(goodOwnerName, goodRepoName)
	if m1 == 0 && m2 == 0 && m3 == 0 && err != nil {
		t.Errorf("user got error %s", err)
	}

}

func TestBadGithubNames(t *testing.T) {
	badOwnerName := "somethingthatsurelydoesntexist"
	badRepoName := "xjaopsfuuoianlrjfkasdaliru"
	m1, m2, m3, err := api.GetCorrectnessFactors(badOwnerName, badRepoName)
	if err == nil || m1 != 0 || m2 != 0 || m3 != 0 {
		t.Errorf("expected error, got %d, %d, %d", m1, m2, m3)
	}

}
