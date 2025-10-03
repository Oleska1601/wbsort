package parser

import (
	"errors"
	"flag"
	"fmt"
)

type Flags struct {
	FlagK int
	FlagN bool
	FlagR bool
	FlagU bool
	FlagM bool
	FlagB bool
	FlagC bool
	FlagH bool
}

func Parse() (*Flags, []string, error) {
	flags := &Flags{}
	flag.IntVar(&flags.FlagK, "k", 1, "sort by column №N") // нумерация с 1, если пользователь передаст 0, то сортировка по всей строке
	flag.BoolVar(&flags.FlagN, "n", false, "sort by numeric value")
	flag.BoolVar(&flags.FlagR, "r", false, "sort by reverse order")
	flag.BoolVar(&flags.FlagU, "u", false, "show only unique values")
	flag.BoolVar(&flags.FlagM, "M", false, "sort by month name")
	flag.BoolVar(&flags.FlagB, "b", false, "ignore trailing blanks")
	flag.BoolVar(&flags.FlagC, "c", false, "check if data is already sorted")
	flag.BoolVar(&flags.FlagH, "h", false, "sort by numeric human suffixes")

	flag.Parse()

	if err := validateFlags(flags); err != nil {
		return nil, nil, fmt.Errorf("validateFlags :%w", err)
	}

	args := flag.Args()
	return flags, args, nil

}

func validateFlags(flags *Flags) error {
	if flags.FlagK < 0 {
		return errors.New("incorrect -k")
	}
	if flags.FlagN && flags.FlagM {
		return errors.New("impossible to use -n -M at the same time")
	}
	if flags.FlagN && flags.FlagH {
		return errors.New("impossible to use -n -h at the same time")
	}
	if flags.FlagM && flags.FlagH {
		return errors.New("impossible to use -M -h at the same time")
	}
	return nil
}

/*

// "-" - указатель читать ищ stdin
// "--" - разделитель (что дальше идут только имена команд, а не опции)

func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "-") && arg != "-" && arg != "--"
}

// -k=2
// -k2
// -k42

// -k 2
func parseKFlag(arg string) {

}

func isCombinedFlag(arg string) bool {
	return strings.HasPrefix(arg, "-") && len(arg[1:]) >=2
}

func parseInput(args []string) {

	var data []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !isFlag(arg) { // если это не флаг а данные для обработки -> добавляем как есть
			data = append(data, arg)
			continue
		}



		// если это флаг

		// флаг -k с присоединенным значением вида -k42 или -k=42
		if strings.HasPrefix(arg, "-k") && len(arg) > 2 {
			n := arg[2:]
			if strings.HasPrefix(n, "=") {
				n = n[1:]
			}
			data = append(data, "-k", n)
			continue
		}

		// флаг -k
		if arg == "-k" {
			data = append(data, "-k")
			for i + 1 < len(args) { //пропускаем пробелы
				nextArg := args[i]
				if

			}
		}


	}
}
*/
