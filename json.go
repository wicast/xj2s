package xj2s

import (
	"regexp"
	"strings"
)

func JsonPath2SrtructLinesNoNesting(paths []string) (map[string]StructNode, map[string]map[string]StructNode) {
	RootStruct := make(map[string]StructNode)
	RestStructs := make(map[string]map[string]StructNode)

	deDuplicateMap := make(map[string]string)
	removeNum := regexp.MustCompile(`\[(\d+)\]`)

	for _, path := range paths {
		path = removeNum.ReplaceAllString(path, "[]")
		Flods := strings.Count(path, "[")
		path = strings.Replace(path, "[]", "", -1)
		splitedPath := strings.Split(path, ".")
		last := splitedPath[len(splitedPath)-1]
		NodeName := strings.Title(last)
		jsonRoute := strings.Join(splitedPath, ">")
		jsonPath := "`json:" + `"` + jsonRoute + `"` + "`"
		Stype := "string"
		for i := 0; i < Flods; i++ {
			Stype = "[]" + Stype
		}
		if _, exist := deDuplicateMap[NodeName]; exist {
			if deDuplicateMap[NodeName] != jsonRoute {
				NodeName = ""
				for _, v := range strings.Split(jsonRoute, ">") {
					NodeName = strings.Title(v) + NodeName
				}
				deDuplicateMap[NodeName] = jsonRoute
			}
		} else {
			deDuplicateMap[NodeName] = jsonRoute
		}
		StructLineAppend := StructNode{Name: NodeName, Type: Stype, Path: jsonPath}
		RootStruct[jsonRoute] = StructLineAppend
	}
	return RootStruct, RestStructs

}
