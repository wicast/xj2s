package tireTree

import (
	"errors"
	"regexp"
	"strings"
)

type TrieTreeNode struct {
	NodeName PathName
	Parent   *TrieTreeNode
	Children map[PathName]*TrieTreeNode
	Value    interface{}
}

type PathName struct {
	Name string
	// IsFlod bool
}

type TreePath []PathName

func NewPaths(pathS string, spliter string) (TreePath, error) {
	removeNum := regexp.MustCompile(`\[(\d+)\]`)
	pathS = removeNum.ReplaceAllString(pathS, "[]")
	splitedPath := strings.Split(pathS, spliter)
	var pN TreePath
	for _, Node := range splitedPath {
		// IsFlod := false
		// if strings.Contains(Node, "[]") {
		// 	IsFlod = true
		// }
		// pN = append(pN, PathName{Name: Node, IsFlod: IsFlod})
		pN = append(pN, PathName{Name: Node})
	}
	return pN, nil
}

func NewNode(NodeName PathName, Value interface{}) TrieTreeNode {
	return TrieTreeNode{NodeName: NodeName,
		Children: make(map[PathName]*TrieTreeNode),
		Value:    Value}
}

func (TT *TrieTreeNode) InsertNode(path TreePath, value interface{}) {
	if Next, exist := TT.Children[path[0]]; !exist {
		if len(path) == 1 {
			ALeafNode := TT.insertSingleNode(path[0])
			ALeafNode.Value = value
		} else {
			NewRouteNode := TT.insertSingleNode(path[0])
			NewRouteNode.InsertNode(path[1:], value)
		}
	} else {
		if len(path) != 1 {
			Next.InsertNode(path[1:], value)
		}
	}

}

func (TT *TrieTreeNode) insertSingleNode(name PathName) *TrieTreeNode {
	NewRouteNode := &TrieTreeNode{NodeName: name, Children: make(map[PathName]*TrieTreeNode)}
	TT.Children[name] = NewRouteNode
	NewRouteNode.Parent = TT
	return NewRouteNode
}

func (TT *TrieTreeNode) deleteSingleNode(nodename PathName) (*TrieTreeNode, error) {
	if Dying, exist := TT.Children[nodename]; exist {
		delete(TT.Children, nodename)
		return Dying, nil
	} else {
		return nil, errors.New("No such node.")
	}
}
