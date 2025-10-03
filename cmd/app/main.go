package main

import (
	"fmt"
	"log"
	"wbsort/internal/parser"
	"wbsort/internal/sorter"
)

func main() {
	flags, args, err := parser.Parse()
	if err != nil {
		log.Fatalln("flags.ParseInput:", err)
	}
	lines, err := parser.ParseArgs(args)
	if err != nil {
		log.Fatalln("parser.ParseArgs:", err)
	}
	sorter := sorter.New(flags, lines)
	sorter.Sort()

	for _, line := range sorter.Lines {
		fmt.Println(line)
	}

}
