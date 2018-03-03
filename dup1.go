package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("please filename")
		os.Exit(1)
	}

	for _, filename := range files {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println(filename, " error.")
			continue
		}

		countLines(f, counts)
		f.Close()
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Println(v, '\t', k)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()[1:4]]++
	}
}
