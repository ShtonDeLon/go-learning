// Полностью решена
package ozontechpoint

import (
	"bufio"
	"fmt"
	"os"
)

func InsertingChars() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	_, err := fmt.Fscan(in, &t)
	if err != nil {
		fmt.Println(err)
	}

	for range t {
		var s string
		_, err = fmt.Fscan(in, &s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintln(out, check(s))
	}
}

func check(s string) string {
	if len(s) > 1 {
		if s[0] != s[len(s)-1] {
			return "NO"
		}

		src := s[0]
		for i := 0; i < len(s)-1; i++ {
			if s[i] == src || s[i+1] == src {
				continue
			}

			if s[i] != src && s[i+1] != src {
				return "NO"
			}
		}
	}

	return "YES"
}
