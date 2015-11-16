package xmltree

import (
	"regexp"

	"github.com/wicast/xj2s/tireTree"
)

type XmlNodeData struct {
	IsTop    bool
	HasLeaf  bool
	IsFlood  bool
	Type     string
	Path     tireTree.NodeName
	FullPath tireTree.TreePath
}

func NewXMLNodeData() XmlNodeData {
	return XmlNodeData{}
}

func NewPath(s string) (tireTree.TreePath, error) {
	s = cleanPath(s)
	return tireTree.NewPaths(s, ".")
}

func NewXMLTree(name string) tireTree.TrieTreeNode {
	return tireTree.NewNode(tireTree.NodeName(name))
}

func cleanPath(path string) string {
	removeNum := regexp.MustCompile(`\[(\d+)\]`)
	return removeNum.ReplaceAllString(path, "[]")
}

func InsertXML(tree *tireTree.TrieTreeNode,
	path tireTree.TreePath,
	value XmlNodeData) (*tireTree.TrieTreeNode, error) {
	ALeafNode, err := tree.InsertNode(path, value)

	if t, ok := tree.Value.(XmlNodeData); ok {
		t.IsTop = true
		tree.SetSelfValue(t)
	} else {
		N := NewXMLNodeData()
		N.IsTop = true
		tree.SetSelfValue(N)
	}
	if V, ok := ALeafNode.Parent.Value.(XmlNodeData); ok {
		V.HasLeaf = true
		ALeafNode.Parent.SetSelfValue(V)
	} else {
		N := NewXMLNodeData()
		N.HasLeaf = true
		ALeafNode.Parent.SetSelfValue(N)
	}
	return ALeafNode, err
}
