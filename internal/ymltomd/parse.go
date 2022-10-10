package ymltomd

import (
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

func (y *Ytm) Parse(b []byte) error {
	err := yaml.Unmarshal(b, &y.tree)
	if err != nil {
		return err
	}
	return nil
}

func (y *Ytm) TreeWalk() {
	y.strBuffer = ""
	y.treeWalk(&y.tree)
}

func (y *Ytm) treeWalk(n *yaml.Node) {
	padding := strings.Repeat(" ", n.Column)
	y.strBuffer += fmt.Sprintf("%s%+v\n", padding, n)
	for _, c := range n.Content {
		y.treeWalk(c)
	}
}
