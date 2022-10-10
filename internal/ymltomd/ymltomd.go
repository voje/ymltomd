// Package ymltomd defines a ymltomd (short ytm) class
// ymltomd implements Reader
package ymltomd

import (
	"fmt"

	yaml "gopkg.in/yaml.v3"
)

type YtmCfg struct{}

type Ytm struct {
	strBuffer string
	bYml      []byte
	data      map[string]interface{}
	// docTree *doctree.DocTree
	tree yaml.Node
}

func NewYtm(cfg YtmCfg) *Ytm {
	ytm := Ytm{}
	return &ytm
}

func (y *Ytm) String() string {
	s := ""

	s += "\ny.bYml\n---\n"
	s += string(y.bYml)

	s += "\ny.data\n---\n"
	s += fmt.Sprintf("%+v\n", y.data)

	s += "\ny.tree\n---\n"
	s += fmt.Sprintf("%+v\n", y.tree)

	s += "\ny.strBuffer\n---\n"
	s += fmt.Sprintf("%+v\n", y.strBuffer)

	/*
		s += "\ny.docTree\n---\n"
		s += fmt.Sprintf("%s\n", y.docTree)
	*/

	return s
}
