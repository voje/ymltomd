package doctree

import (
	"fmt"
	"strings"
)

type DocTree struct {
	Root      *Node
	stringRep string
}

func NewDocTree(m map[string]interface{}) *DocTree {
	dt := DocTree{}
	dt.Root = parseMap("root", m)
	return &dt
}

type Node struct {
	Name     string
	Children []*Node
}

func parseLeaf(name string, m interface{}) *Node {
	n := Node{
		Name: name,
	}
	n.Children = append(n.Children, &Node{Name: fmt.Sprintf("%v", m)})
	return &n
}

func parseMap(name string, m map[string]interface{}) *Node {
	n := Node{
		Name: name,
	}
	for k, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			n.Children = append(n.Children, parseMap(k, v.(map[string]interface{})))
		case interface{}:
			n.Children = append(n.Children, parseLeaf(k, v))
		case string:
			n.Children = append(n.Children, &Node{Name: k})
		}
	}
	return &n
}

func (dt *DocTree) stringify(n *Node, depth int) {
	padding := strings.Repeat(" ", depth*2)
	if n == nil {
		return
	}
	dt.stringRep += fmt.Sprintf("%s%v\n", padding, n.Name)
	for _, c := range n.Children {
		dt.stringify(c, depth+1)
	}
}

// String returns a string representation of the document tree
func (dt *DocTree) String() string {
	dt.stringRep = ""
	dt.stringify(dt.Root, 0)
	ret := dt.stringRep
	dt.stringRep = ""
	return ret
}
