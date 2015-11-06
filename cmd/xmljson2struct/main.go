package main

import (
	"flag"
	"fmt"
	"github.com/wicast/xj2s"
	"io/ioutil"
	"os"
	"runtime"
)

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " [-flags] [file]")
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func main() {
	ParseType := flag.String("t", "xml", "Type to parse\n\tavaliable type:xml,json")
	FileName := flag.String("f", "", "Parse from a file given a name")
	JsonRootName := flag.String("root", "JsonRoot", "For struct root name when using json,Default is JsonRoot")
	Nesting := flag.Bool("n", false, "Generate structs whit nesting style\n\tnotice:json haven't implement the > style yet,so only nesting style is working for json")
	flag.Parse()
	if *FileName != "" {
		bytes, err := ioutil.ReadFile(*FileName)
		if err != nil {
			panic(err)
		}
		if *ParseType == "xml" {
			fmt.Println(xj2s.Xml2Struct(bytes, *Nesting))
		} else if *ParseType == "json" {
			fmt.Println(xj2s.Json2Struct(bytes, *JsonRootName, *Nesting))
		} else {
			usage()
		}
	} else if flag.NArg() != 0 {
		bytes, err := ioutil.ReadFile(flag.Args()[0])
		if err != nil {
			panic(err)
		}
		if *ParseType == "xml" {
			fmt.Println(xj2s.Xml2Struct(bytes, *Nesting))
		} else if *ParseType == "json" {
			fmt.Println(xj2s.Json2Struct(bytes, *JsonRootName, *Nesting))
		} else {
			usage()
		}

	} else if runtime.GOOS != "windows" {
		if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {

			bytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}
			if *ParseType == "xml" {
				fmt.Println(xj2s.Xml2Struct(bytes, *Nesting))
			} else if *ParseType == "json" {
				fmt.Println(xj2s.Json2Struct(bytes, *JsonRootName, *Nesting))
			} else {
				usage()
			}
		}
	} else {
		usage()
	}
}
