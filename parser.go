package main

import (
	"flag"
)

type ParseVar struct {
	Value string
	Name  string
	Usage string
}

type ParseVarList []ParseVar

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseReader(parseList ParseVarList) {
	for _, parseVar := range parseList {
		flag.StringVar(&parseVar.Value, parseVar.Name, "ABC", parseVar.Usage)
	}
	flag.Parse()
}
