package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)      // Initialize new map
	input := bufio.NewScanner(os.Stdin) // Scan inoput

	for input.Scan() { // Iterate input
		counts[input.Text()]++
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}
