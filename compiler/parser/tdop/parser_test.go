package tdop

import (
	"testing"

	"github.com/PuerkitoBio/goblin/compiler/scanner"
)

var (
	cases = []struct {
		src []byte
	}{
		0: {
			src: []byte(`return 5`),
		},
		1: {
			src: []byte(`aB := 5
return aB`),
		},
	}

	isolateCase = 1
)

func TestParse(t *testing.T) {
	Scanner = new(scanner.Scanner)
	for i, c := range cases {
		if isolateCase >= 0 && i != isolateCase {
			continue
		}

		Parse("test", c.src)
	}
}