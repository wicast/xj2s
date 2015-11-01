package xj2s

import (
	"github.com/clbanning/mxj"
	// "fmt"
	// "regexp"
	// "strings"
)

type StructNode struct {
	Name string
	Type string
	Path string
}

func Xml2Struct(xdata []byte) string {
	m, err := mxj.NewMapXml(xdata)
	if err != nil {
		panic(err)
	}
	paths := m.LeafPaths()
	RootName, RootStruct, RestStructs := Path2SrtructLines(paths)
	return RootXmlDatas2Struct(RootName, RootStruct, RestStructs)
}
