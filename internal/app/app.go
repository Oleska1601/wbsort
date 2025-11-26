package app

import (
	"fmt"

	"github.com/Oleska1601/wbsort/internal/parser"
	"github.com/Oleska1601/wbsort/internal/sorter"
)

// Run starts the sorting application with command line arguments
func Run(flags *parser.Flags, lines []string) {
	sorter := sorter.New(flags, lines)
	rows := sorter.Sort()

	for _, row := range rows {
		fmt.Println(row)
	}
}
