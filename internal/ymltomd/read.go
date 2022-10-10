package ymltomd

import (
	yaml "gopkg.in/yaml.v3"
)

func (y *Ytm) Read(p []byte) (n int, err error) {
	y.bYml = p

	y.data = make(map[string]interface{})
	err = yaml.Unmarshal(y.bYml, y.data)
	if err != nil {
		return 0, err
	}

	// y.docTree = doctree.NewDocTree(y.data)

	return 0, nil
}
