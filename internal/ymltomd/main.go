// Package ymltomd defines a ymltomd (short ytm) class
// ymltomd implements Reader
package ymltomd

type YtmCfg struct{}

type Ytm struct {
	bYml []byte
}

func NewYtm(cfg YtmCfg) *Ytm {
	ytm := Ytm{}
	return &ytm
}

func (y *Ytm) Read(p []byte) (n int, err error) {
	y.bYml = p
	return 0, nil
}
