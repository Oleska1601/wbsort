package main

import (
	"log"
	"wbsort/internal/app"
	"wbsort/internal/parser"
)

func main() {
	flags, files := parser.Parse()

	lines, err := parser.ParseArgs(files)
	if err != nil {
		log.Fatalln("parser.ParseArgs:", err)
	}

	app.Run(flags, lines)
}
