package ymltomd

import (
	"fmt"

	"github.com/atsushinee/go-markdown-generator/doc"
)

func (y *Ytm) GenMd() string {
	d := doc.NewMarkDown()
	d.WriteTitle("Test title", 1)
	d.WriteLines(3)
	d.WriteTitle("Test title2", 2)
	d.Write("Testing 123")

	return fmt.Sprintf("%s", d)
}
