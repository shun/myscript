package main

import (
	"bufio"
	"os"
)

func main() {
	lines := readlines()
	newlines := removeDuplicate(lines)
	writelines(newlines)
}

func readlines() []string {
	var fp *os.File
	var err error

	var lines []string
	fp, err = os.Open(os.Getenv("HOME") + "/.bash_history")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func writelines(lines []string) {
	var fp *os.File
	var err error

	fp, err = os.Create(os.Getenv("HOME") + "/.bash_history")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	writer := bufio.NewWriter(fp)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}

func removeDuplicate(args []string) []string {
	results := make([]string, 0, len(args))
	encountered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}
