package xj2s

import (
	"github.com/clbanning/mxj"
	// "fmt"
	"regexp"
	"strings"
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

func Path2SrtructLines(paths []string) (string, map[string]StructNode, map[string]map[string]StructNode) {
	var RootName string
	RootStruct := make(map[string]StructNode)
	RestStructs := make(map[string]map[string]StructNode)

	RootName = strings.Split(paths[0], ".")[0]

	deDuplicateMap := make(map[string]string)

	removeNum := regexp.MustCompile(`\[(\d+)\]`)
	for _, path := range paths {
		path = removeNum.ReplaceAllString(path, "[]")
		Flods := strings.Count(path, "[")
		path = strings.Replace(path, "[]", "", -1)
		splitedPath := strings.Split(path, ".")
		last := splitedPath[len(splitedPath)-1]
		if strings.Index(last, "-") == 0 { //Attr
			if RootName == splitedPath[len(splitedPath)-2] { //RootAttr
				NodeName := strings.Title(last[1:])
				xmlRoute := "`xml:" + `"` + last[1:] + `,attr"` + "`"
				if _, exist := deDuplicateMap[NodeName]; exist {
					if deDuplicateMap[NodeName] != xmlRoute {
						NodeName = "Rss" + NodeName
						deDuplicateMap[NodeName] = xmlRoute
					}
				} else {
					deDuplicateMap[NodeName] = xmlRoute
				}
				StructLineAppend := StructNode{Name: NodeName, Type: "string", Path: xmlRoute}
				RootStruct[xmlRoute] = StructLineAppend
			} else { //NoneRootAttr
				NodeName := strings.Title(splitedPath[len(splitedPath)-2])
				xmlRoute := strings.Join(splitedPath[1:len(splitedPath)-1], ">")
				xmlPath := "`xml:" + `"` + xmlRoute + `"` + "`"
				Stype := NodeName
				for i := 0; i < Flods; i++ {
					Stype = "[]" + Stype
				}
				if _, exist := deDuplicateMap[NodeName]; exist {
					if deDuplicateMap[NodeName] != xmlRoute {
						NodeName = ""
						for _, v := range strings.Split(xmlRoute, ">") {
							NodeName = strings.Title(v) + NodeName
						}
						deDuplicateMap[NodeName] = xmlRoute
					}
				} else {
					deDuplicateMap[NodeName] = xmlRoute
				}
				StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: xmlPath}
				RootStruct[xmlRoute] = StructLineAppend

				LeafName := strings.Title(last[1:])
				RsetStructLineAppend := StructNode{Name: LeafName, Type: "string", Path: "`xml:" + `"` + last[1:] + `,attr"` + "`"}

				// log.Println(NodeName, LeafName)
				if _, exist := RestStructs[NodeName]; exist {
					RestStructs[NodeName][LeafName] = RsetStructLineAppend
				} else {
					NewLeafStruct := make(map[string]StructNode)
					NewLeafStruct[LeafName] = RsetStructLineAppend
					RestStructs[NodeName] = NewLeafStruct
				}

			}
		} else if strings.Index(last, "#") == 0 { //chardata
			if RootName == splitedPath[len(splitedPath)-2] { //RootChartata
				NodeName := strings.Title(last[1:])
				xmlRoute := "`xml:" + `",chardata"` + "`"
				if _, exist := deDuplicateMap[NodeName]; exist {
					if deDuplicateMap[NodeName] != xmlRoute {
						NodeName = "Rss" + NodeName
						deDuplicateMap[NodeName] = xmlRoute
					}
				} else {
					deDuplicateMap[NodeName] = xmlRoute
				}
				StructLineAppend := StructNode{Name: NodeName, Type: "string", Path: xmlRoute}
				RootStruct[xmlRoute] = StructLineAppend
			} else { //NonRootChardata
				NodeName := strings.Title(splitedPath[len(splitedPath)-2])
				xmlRoute := strings.Join(splitedPath[1:len(splitedPath)-1], ">")
				xmlPath := "`xml:" + `"` + xmlRoute + `"` + "`"
				Stype := NodeName
				for i := 0; i < Flods; i++ {
					Stype = "[]" + Stype
				}
				if _, exist := deDuplicateMap[NodeName]; exist {
					if deDuplicateMap[NodeName] != xmlRoute {
						NodeName = ""
						for _, v := range strings.Split(xmlRoute, ">") {
							NodeName = strings.Title(v) + NodeName
						}
						deDuplicateMap[NodeName] = xmlRoute
					}
				} else {
					deDuplicateMap[NodeName] = xmlRoute
				}
				StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: xmlPath}
				RootStruct[xmlRoute] = StructLineAppend

				LeafName := strings.Title(last[1:])
				RsetStructLineAppend := StructNode{Name: LeafName, Type: "string", Path: "`xml:" + `",chardata"` + "`"}

				if _, exist := RestStructs[NodeName]; exist {
					RestStructs[NodeName][LeafName] = RsetStructLineAppend
				} else {
					NewLeafStruct := make(map[string]StructNode)
					NewLeafStruct[LeafName] = RsetStructLineAppend
					RestStructs[NodeName] = NewLeafStruct
				}
			}
		} else { //NormalString
			NodeName := strings.Title(splitedPath[len(splitedPath)-1])
			xmlRoute := strings.Join(splitedPath[1:], ">")
			xmlPath := "`xml:" + `"` + xmlRoute + `"` + "`"
			Stype := "string"
			for i := 0; i < Flods; i++ {
				Stype = "[]" + Stype
			}
			if _, exist := deDuplicateMap[NodeName]; exist {
				if deDuplicateMap[NodeName] != xmlRoute {
					NodeName = ""
					for _, v := range strings.Split(xmlRoute, ">") {
						NodeName = strings.Title(v) + NodeName
					}
					deDuplicateMap[NodeName] = xmlRoute
				}
			} else {
				deDuplicateMap[NodeName] = xmlRoute
			}
			StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: xmlPath}
			RootStruct[xmlRoute] = StructLineAppend
		}
	}
	return strings.Title(RootName), RootStruct, RestStructs
}

func RootXmlDatas2Struct(RootName string, RootLines map[string]StructNode, RestStructs map[string]map[string]StructNode) string {
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
