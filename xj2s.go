package xj2s

import (
	// "github.com/clbanning/mxj"
	// "log"
	"regexp"
	"strings"
)

// type StructRest struct {
// 	Leafs map[string]StructNode
// }
type StructNode struct {
	Name string
	Type string
	Path string
}

// func Xml2Struct(xdata []byte) (string, error) {
// 	m, err := mxj.NewMapXml(xdata)
// 	if err != nil {
// 		panic(err)
// 	}
// 	paths := m.LeafPaths()
// 	RootName, StructLines, LeafStructs := Path2SrtructLines(paths)
// 	return FormatStruct(RootName, StructLines, LeafStructs), err
// }

func Path2SrtructLines(paths []string) (string, map[string]StructNode, map[string]map[string]StructNode) {
	var RootName string
	StructLines := make(map[string]StructNode)
	RestStructs := make(map[string]map[string]StructNode)

	RootName = strings.Split(paths[0], ".")[0]

	removeNum := regexp.MustCompile(`\[(\d+)\]`)
	for _, path := range paths {
		path = removeNum.ReplaceAllString(path, "[]")
		Flods := strings.Count(path, "[")
		path = strings.Replace(path, "[]", "", -1)
		splitPath := strings.Split(path, ".")
		last := splitPath[len(splitPath)-1]
		if strings.Index(last, "-") == 0 { //Attr
			if RootName == splitPath[len(splitPath)-2] { //RootAttr
				NodeName := strings.Title(last[1:])
				xmlRoute := "`xml:" + `"` + last[1:] + `,attr"` + "`"
				StructLineAppend := StructNode{Name: NodeName, Type: "string", Path: xmlRoute}
				StructLines[xmlRoute] = StructLineAppend
			} else { //NoneRootAttr
				NodeName := strings.Title(splitPath[len(splitPath)-2])
				xmlRoute := strings.Join(splitPath[1:len(splitPath)-1], ">")
				xmlPath := "`xml:" + `"` + xmlRoute + `"` + "`"
				Stype := "[]" + NodeName
				for i := 0; i < Flods; i++ {
					Stype = "[]" + Stype
				}
				StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: xmlPath}
				StructLines[xmlRoute] = StructLineAppend

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
			if RootName == splitPath[len(splitPath)-2] { //RootChartata
				NodeName := strings.Title(last[1:])
				xmlRoute := "`xml:" + `",chardata"` + "`"
				StructLineAppend := StructNode{Name: NodeName, Type: "string", Path: xmlRoute}
				StructLines[xmlRoute] = StructLineAppend
			} else { //NonRootChardata
				NodeName := strings.Title(splitPath[len(splitPath)-2])
				xmlRoute := strings.Join(splitPath[1:len(splitPath)-1], ">")
				xmlPath := "`xml:" + `"` + xmlRoute + `"` + "`"
				Stype := "[]" + NodeName
				for i := 0; i < Flods; i++ {
					Stype = "[]" + Stype
				}
				StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: xmlPath}
				StructLines[xmlRoute] = StructLineAppend

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
		} else {
			NodeName := strings.Title(splitPath[len(splitPath)-1])
			xmlRoute := strings.Join(splitPath[1:], ">")
			xmlPath := "`xml:" + `"` + xmlRoute + `"` + "`"
			Stype := "[]" + "string"
			for i := 0; i < Flods; i++ {
				Stype = "[]" + Stype
			}
			StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: xmlPath}
			StructLines[xmlRoute] = StructLineAppend
		}
	}
	return RootName, StructLines, RestStructs
}
