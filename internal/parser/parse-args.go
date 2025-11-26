package parser

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ParseArgs processes command line arguments as files or os.Stdin and returns lines
func ParseArgs(args []string) ([]string, error) {
	if len(args) == 0 {
		return readLines(os.Stdin)
	}

	var totalLines []string
	if len(args) > 0 {
		for _, file := range args {
			file, err := os.Open(file)
			if err != nil {
				return nil, fmt.Errorf("open file: %w", err)
			}

			lines, err := readLines(file)
			if err != nil {
				return nil, fmt.Errorf("readLines: %w", err)
			}

			file.Close()
			totalLines = append(totalLines, lines...)
		}
	}

	return totalLines, nil
}

func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return lines, nil
}
