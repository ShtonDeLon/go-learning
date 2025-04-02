package chapter1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup1() {
	count := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "eof" {
			count[scanner.Text()]++
		} else {
			break
		}
	}
	for k, v := range count {
		if v > 1 {
			fmt.Printf("%s\n", k)
		}
	}
}
