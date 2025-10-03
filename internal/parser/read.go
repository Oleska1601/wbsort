package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func ParseArgs(args []string) ([]string, error) {
	if len(args) == 0 {
		lines, err := readLines(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("error readLines: %w", err)
		}
		return lines, nil
	}

	lines, err := readFiles(args)
	if err != nil {
		return nil, fmt.Errorf("error readFiles: %w", err)
	}

	return lines, nil
}

func readFiles(files []string) ([]string, error) {
	var result []string
	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			return nil, fmt.Errorf("error os.Open: %w", err)
		}
		defer file.Close()
		lines, err := readLines(file)
		if err != nil {
			return nil, fmt.Errorf("error readLines: %w", err)
		}
		result = append(result, lines...)
	}
	return result, nil
}

// aaaa
// aaa aaaaa aaaaaaa

func readLines(file *os.File) ([]string, error) {
	reader := bufio.NewReader(file)
	var lines []string
	var buf []byte
	for {
		line, isPrefix, err := reader.ReadLine() //isPrefix = true -> неполная строка
		if errors.Is(err, io.EOF) {
			if len(buf) > 0 {
				lines = append(lines, string(buf))
			}

			break
		}
		if err != nil {
			return nil, fmt.Errorf("readLine error: %w", err)
		}
		buf = append(buf, line...)
		if !isPrefix { // если полная строка
			lines = append(lines, string(buf))
			buf = buf[:0]
		}
	}
	return lines, nil
}
