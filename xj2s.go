package xj2s

import (
	"github.com/clbanning/mxj"
	// "fmt"
	// "regexp"
	"strings"
)

type StructNode struct {
	Name string
	Type string
	Path string
}

func Xml2Struct(xdata []byte, Nesting bool) string {
	m, err := mxj.NewMapXml(xdata)
	if err != nil {
		panic(err)
	}
	paths := m.LeafPaths()
	if Nesting {
		return "Not implement yet..."
	} else {
		RootName, RootStruct, RestStructs := XmlPath2SrtructLinesNoNesting(paths)
		return RootDatas2Struct(RootName, RootStruct, RestStructs)
	}
}

func Json2Struct(jdata []byte, RootName string, Nesting bool) string {
	m, err := mxj.NewMapJson(jdata)
	if err != nil {
		panic(err)
	}
	paths := m.LeafPaths()
	if Nesting {
		return "Not implement yet..."
	} else {
		RootStruct, RestStructs := JsonPath2SrtructLinesNoNesting(paths)
		return RootDatas2Struct(strings.Title(RootName), RootStruct, RestStructs)
	}
}

func RootDatas2Struct(RootName string, RootLines map[string]StructNode, RestStructs map[string]map[string]StructNode) string {
	Structs := "type " + RootName + " struct {\n"
	for _, v := range RootLines {
		Structs += "\t" + v.Name + "\t" + v.Type + "\t" + v.Path + "\n"
	}
	Structs += "}\n\n"

	for NodeName, v1 := range RestStructs {
		Structs += "type " + NodeName + " struct {\n"
		for _, v2 := range v1 {
			Structs += "\t" + v2.Name + "\t" + v2.Type + "\t" + v2.Path + "\n"
		}
		Structs += "}\n"
	}
	return Structs
}
