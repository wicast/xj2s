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

func NewNode(NodeName PathName) TrieTreeNode {
	return TrieTreeNode{NodeName: NodeName, Children: make(map[PathName]*TrieTreeNode)}
}

func (TT *TrieTreeNode) InsertNode(path TreePath) (*TrieTreeNode, error) {
	if Next, exist := TT.Children[path[0]]; !exist {
		if len(path) == 1 {
			ALeafNode := TT.insertSingleNode(path[0])
			return ALeafNode, nil
		} else {
			NewRouteNode := TT.insertSingleNode(path[0])
			return NewRouteNode.InsertNode(path[1:])
		}
	} else {
		if len(path) != 1 {
			return Next.InsertNode(path[1:])
		} else {
			return nil, errors.New("Insert Node Failed.")
		}
	}

}

func (TT *TrieTreeNode) insertSingleNode(name PathName) *TrieTreeNode {
	NewRouteNode := &TrieTreeNode{NodeName: name, Children: make(map[PathName]*TrieTreeNode)}
	TT.Children[name] = NewRouteNode
	NewRouteNode.Parent = TT
	return NewRouteNode
}

func (TT *TrieTreeNode) GetNode(path TreePath) (*TrieTreeNode, error) {
	if Next, exist := TT.Children[path[0]]; exist {
		if len(path) == 1 {
			return Next, nil
		} else {
			return Next.GetNode(path[1:])
		}
	} else {
		return nil, errors.New("No such node")
	}
}

func (TT *TrieTreeNode) DeleteNode(nodePath TreePath) (*TrieTreeNode, error) {
	if Next, exist := TT.Children[nodePath[0]]; exist {
		if len(nodePath) == 1 {
			DyingNode, err := TT.deleteSingleNode(nodePath[0])
			return DyingNode, err
		} else {
			return Next.DeleteNode(nodePath[1:])
		}
	}
	return nil, errors.New("No such route.")
}

func (TT *TrieTreeNode) deleteSingleNode(nodename PathName) (*TrieTreeNode, error) {
	if Dying, exist := TT.Children[nodename]; exist {
		delete(TT.Children, nodename)
		return Dying, nil
	} else {
		return nil, errors.New("No such node.")
	}
}
