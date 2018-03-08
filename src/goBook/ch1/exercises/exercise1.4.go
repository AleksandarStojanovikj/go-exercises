package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	filesWithDup := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesWithDup)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "exercise1.4: %v\n", err)
				continue
			}
			countLines(f, counts, filesWithDup)
			f.Close()
		}
	}
	for line := range filesWithDup {
		fmt.Printf("File with dup:\t%s\n", line)
	}
}
func countLines(f *os.File, counts map[string]int, filesWithDup map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			filesWithDup[f.Name()]++
		}
	}
}
