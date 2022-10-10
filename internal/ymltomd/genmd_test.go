package ymltomd

import "testing"

func TestGen1(t *testing.T) {
	y := NewYtm(YtmCfg{})
	t.Log(y.GenMd())
}
