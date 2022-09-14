package play

import (
	"testing"
)

func TestConst(t *testing.T) {
	geneS := func() string {
		return "GENE"
	}
	geneI := func() int {
		return 0
	}
	// const err = errors.New("OK") // error: is not constant
	// const g = geneS()             // error: is not constant
	// const i = geneI()            // error: is not constant
	geneS()
	geneI()
}
