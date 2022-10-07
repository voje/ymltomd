// Package ymltomd defines a ymltomd (short ytm) class
// ymltomd implements Reader
package ymltomd

import (
	"fmt"

	"github.com/voje/ymltomd/internal/doctree"
	yaml "gopkg.in/yaml.v3"
)

type YtmCfg struct{}

type Ytm struct {
	bYml    []byte
	data    map[string]interface{}
	docTree *doctree.DocTree
}

func NewYtm(cfg YtmCfg) *Ytm {
	ytm := Ytm{}
	return &ytm
}

func (y *Ytm) Read(p []byte) (n int, err error) {
	y.bYml = p

	y.data = make(map[string]interface{})
	err = yaml.Unmarshal(y.bYml, y.data)
	if err != nil {
		return 0, err
	}

	y.docTree = doctree.NewDocTree(y.data)

	return 0, nil
}

func (y *Ytm) String() string {
	s := ""

	s += "\ny.bYml\n---\n"
	s += string(y.bYml)

	s += "\ny.data\n---\n"
	s += fmt.Sprintf("%+v\n", y.data)

	s += "\ny.docTree\n---\n"
	s += fmt.Sprintf("%s\n", y.docTree)

	return s
}
