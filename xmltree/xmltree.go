package xmltree

import (
	"github.com/wicast/xj2s/tireTree"
	"regexp"
)

type XmlTree struct {
	tireTree.TrieTreeNode
}

type XmlNodeData struct {
	IsTop   bool
	HasLeaf bool
	IsFlood bool
	Data
}

type Data struct {
	Type     string
	Path     string
	FullPath tireTree.PathName
}

func (XT *XmlTree) RawInsertXmlNode(path tireTree.TreePath, value XmlNodeData) (*tireTree.TrieTreeNode, error) {
	return XT.InsertNode(path, value)
}

func (XT *XmlTree) SetXmlNodeValue(path tireTree.TreePath, newValue XmlNodeData) error {
	err := XT.SetNodeValue(path, newValue)
	return err
}

func CleanPath(path string) string {
	removeNum := regexp.MustCompile(`\[(\d+)\]`)
	return removeNum.ReplaceAllString(path, "[]")
}
