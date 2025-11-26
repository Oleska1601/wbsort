package parser

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	FlagK      int
	FlagN      bool
	FlagR      bool
	FlagU      bool
	FlagM      bool
	FlagB      bool
	FlagC      bool
	FlagH      bool
	InputFiles []string
}

func Parse() (*Flags, []string) {
	flags := &Flags{}
	pflag.IntVarP(&flags.FlagK, "key", "k", 0, "sort by column №N")
	pflag.BoolVarP(&flags.FlagN, "numeric-sort", "n", false, "sort by numeric value")
	pflag.BoolVarP(&flags.FlagR, "reverse", "r", false, "sort by reverse order")
	pflag.BoolVarP(&flags.FlagU, "unique", "u", false, "show only unique values")
	pflag.BoolVarP(&flags.FlagM, "month-sort", "M", false, "sort by month name")
	pflag.BoolVarP(&flags.FlagB, "ignore-trailing-blanks", "b", false, "ignore trailing blanks")
	pflag.BoolVarP(&flags.FlagC, "check", "c", false, "check if data is already sorted")
	pflag.BoolVarP(&flags.FlagH, "human-numeric-sort", "h", false, "sort by numeric human suffixes")

	pflag.Parse()

	args := pflag.Args()
	return flags, args
}
