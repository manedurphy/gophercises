package goByExamples

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScannerLines() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func ScannerWords() {
	const input = "Now here is a string with many words,\nAnd here is another string on a new line"

	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	count := 0 // The number of words in out input string

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("%d\n", count)
}
