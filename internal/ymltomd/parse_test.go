package ymltomd

import (
	"testing"

	"github.com/voje/ymltomd/internal/testdocs"
)

func TestComment(t *testing.T) {
	yc := YtmCfg{}
	y := NewYtm(yc)

	y.Parse(testdocs.Doc1)
	y.TreeWalk()
	t.Log(y)
}
