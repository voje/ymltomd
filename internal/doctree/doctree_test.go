package doctree

import "testing"

var map1 map[string]interface{}

func init() {
	map1 = make(map[string]interface{})
	map1a := make(map[string]interface{})
	map1a["b"] = 2
	map1a["c"] = 4
	map1["a"] = map1a
	map1["d"] = 5

}

func TestNewTree(t *testing.T) {
	dt := NewDocTree(map1)
	t.Log(dt)
}
