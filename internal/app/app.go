package app

import (
	"fmt"
	"wbsort/internal/parser"
	"wbsort/internal/sorter"
)

func Run(flags *parser.Flags, lines []string) {
	sorter := sorter.New(flags, lines)
	rows := sorter.Sort()

	for _, row := range rows {
		fmt.Println(row)
	}
}
