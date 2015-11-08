package tireTree

import (
	"errors"
	"strings"
)

type TrieTreeNode struct {
	Name     NodeName
	Parent   *TrieTreeNode
	Children map[NodeName]*TrieTreeNode
	Value    interface{}
}

type NodeName struct {
	Name string
}

type TreePath []NodeName

func NewPaths(pathS string, spliter string) (TreePath, error) {
	splitedPath := strings.Split(pathS, spliter)
	var pN TreePath
	for _, Node := range splitedPath {
		pN = append(pN, NodeName{Name: Node})
	}
	return pN, nil
}

func NewNode(Name NodeName) TrieTreeNode {
	return TrieTreeNode{Name: Name, Children: make(map[NodeName]*TrieTreeNode)}
}

func (TT *TrieTreeNode) InsertNode(path TreePath, value interface{}) (*TrieTreeNode, error) {
	if Next, exist := TT.Children[path[0]]; !exist {
		if len(path) == 1 {
			ALeafNode := TT.insertSingleNode(path[0], value)
			return ALeafNode, nil
		} else {
			NewRouteNode := TT.insertSingleNode(path[0], nil)
			NewRouteNode.Name = path[0]
			return NewRouteNode.InsertNode(path[1:], value)
		}
	} else {
		if len(path) != 1 {
			return Next.InsertNode(path[1:], value)
		} else {
			return nil, errors.New("Insert Node Failed.")
		}
	}

}

func (TT *TrieTreeNode) insertSingleNode(name NodeName, value interface{}) *TrieTreeNode {
	NewRouteNode := &TrieTreeNode{Name: name, Children: make(map[NodeName]*TrieTreeNode), Value: value}
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

func (TT *TrieTreeNode) GetSingleNode(path NodeName) (*TrieTreeNode, error) {
	if Next, exist := TT.Children[path]; exist {
		return Next, nil
	} else {
		return nil, errors.New("No such node")
	}
}

func (TT *TrieTreeNode) GetNodeValue() interface{} {
	return TT.Value
}

func (TT *TrieTreeNode) SetNodeValue(path TreePath, value interface{}) error {
	t, err := TT.GetNode(path)
	t.Value = value
	return err
}

func (TT *TrieTreeNode) SetSingleNodeValue(name NodeName, value interface{}) error {
	t, err := TT.GetSingleNode(name)
	t.Value = value
	return err
}

func (TT *TrieTreeNode) SetSelfValue(value interface{}) {
	TT.Value = value
}

func (TT *TrieTreeNode) DeleteNode(nodePath TreePath) (*TrieTreeNode, error) {
	if Next, exist := TT.Children[nodePath[0]]; exist {
		if len(nodePath) == 1 {
			DyingNode, err := TT.DeleteSingleNode(nodePath[0])
			return DyingNode, err
		} else {
			return Next.DeleteNode(nodePath[1:])
		}
	}
	return nil, errors.New("No such route.")
}

func (TT *TrieTreeNode) DeleteSingleNode(nodename NodeName) (*TrieTreeNode, error) {
	if Dying, exist := TT.Children[nodename]; exist {
		delete(TT.Children, nodename)
		return Dying, nil
	} else {
		return nil, errors.New("No such node.")
	}
}
