package ymltomd

import (
	"testing"

	"github.com/voje/ymltomd/internal/testdocs"
)

func TestRead(t *testing.T) {
	ytm := NewYtm(YtmCfg{})
	ytm.Read(testdocs.Doc1)
	t.Log(ytm)
}
