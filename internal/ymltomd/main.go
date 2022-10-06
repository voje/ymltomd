// Package ymltomd defines a ymltomd (short ytm) class
// ymltomd implements Reader
package ymltomd

import (
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type YtmCfg struct{}

type Ytm struct {
	bYml []byte
	data map[string]interface{}
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

	y.traverse(y.data, 0)

	return 0, nil
}

func (y *Ytm) traverse(node map[string]interface{}, depth int) {
	padding := strings.Repeat(" ", depth*2)
	for k, v := range node {
		fmt.Printf("%s[%v]: ", padding, k)
		switch v.(type) {
		case map[string]interface{}:
			fmt.Println()
			y.traverse(v.(map[string]interface{}), depth+1)
		case string:
			fmt.Printf("%v\n", v)
		}
	}
}

func (y *Ytm) String() string {
	s := ""

	s += "\ny.bYml\n---\n"
	s += string(y.bYml)

	s += "\ny.data\n---\n"
	s += fmt.Sprintf("%+v\n", y.data)

	return s
}
