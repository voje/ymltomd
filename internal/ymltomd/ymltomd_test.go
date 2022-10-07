package ymltomd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {

	doc := []byte(`---
dict_struct: 
  level1:
    level2: "hhey22"`)

	ytm := NewYtm(YtmCfg{})
	ytm.Read(doc)

	assert.Equal(
		t,
		"hhey22",
		ytm.docTree.Root.Children[0].Children[0].Children[0].Children[0].Name,
	)

	t.Log(ytm)
}
