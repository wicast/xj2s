package main

import (
	"flag"
	"fmt"
	"github.com/wicast/xj2s"
	"io/ioutil"
	"os"
)

func parseXml(s []byte) string {
	return xj2s.Xml2Struct(s)
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " [-flags] [file]")
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func main() {
	ParseType := flag.String("t", "xml", "Type for parse\n\tavaliable type:xml,json")
	FileName := flag.String("f", "", "Parse from a file given a name")
	Stdin := flag.Bool("stdin", false, "Parse from stdin")
	flag.Parse()
	if *FileName != "" {
		f, err := ioutil.ReadFile(*FileName)
		if err != nil {
			panic(err)
		}
		if *ParseType == "xml" {
			fmt.Println(parseXml(f))
		} else if *ParseType == "json" {
			fmt.Println("Not implemented yet...")
		} else {
			usage()
		}
	} else if *Stdin == true {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		if *ParseType == "xml" {
			fmt.Println(parseXml(bytes))
		} else if *ParseType == "json" {
			fmt.Println("Not implemented yet...")
		} else {
			usage()
		}
	} else if flag.NArg() != 0 {
		bytes, err := ioutil.ReadFile(flag.Args()[0])
		if err != nil {
			panic(err)
		}
		if *ParseType == "xml" {
			fmt.Println(parseXml(bytes))
		} else if *ParseType == "json" {
			fmt.Println("Not implemented yet...")
		} else {
			usage()
		}

	} else {
		usage()
	}
}
