package main

import (
	"log"

	"github.com/Oleska1601/wbsort/internal/app"
	"github.com/Oleska1601/wbsort/internal/parser"
)

func main() {
	flags, files := parser.Parse()

	lines, err := parser.ParseArgs(files)
	if err != nil {
		log.Fatalln("parser.ParseArgs:", err)
	}

	app.Run(flags, lines)
}
