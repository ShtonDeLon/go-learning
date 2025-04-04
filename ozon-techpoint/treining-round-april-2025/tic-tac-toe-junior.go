// Частично решена
// Скорее всего плохо подобран алгоритм,
// так как все операции со строками сведены к байтовым,
// что обеспечивает выигрыш и по памяти и по времени
package ozontechpoint

import (
	"bufio"
	"fmt"
	"os"
)

func TicTacToe() {
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
		var k int
		_, err = fmt.Fscan(in, &k)
		if err != nil {
			fmt.Println(err)
		}

		var n, m int
		_, err = fmt.Fscan(in, &n, &m)
		if err != nil {
			fmt.Println(err)
		}

		board := make([][]byte, n)
		yesFlag := false
		winFlag := false

		for i := range n {
			board[i] = make([]byte, m)

			var s string
			_, err = fmt.Fscan(in, &s)
			if err != nil {
				fmt.Println(err)
			}

			board[i] = []byte(s)
		}

		for i := range n {
			count := m - k + 1
			for j := range count {
				switch xCheck(getHorizontalLine(board, k, i, j), k) {
				case "WIN":
					winFlag = true
				case "YES":
					yesFlag = true
				}
				if winFlag {
					break
				}
			}
			if winFlag {
				break
			}
		}
		if winFlag {
			fmt.Fprintln(out, "NO")
			continue
		}

		for j := range m {
			count := n - k + 1
			for i := range count {
				switch xCheck(getVerticalLine(board, k, i, j), k) {
				case "WIN":
					winFlag = true
				case "YES":
					yesFlag = true
				}
				if winFlag {
					break
				}
			}
			if winFlag {
				break
			}
		}
		if winFlag {
			fmt.Fprintln(out, "NO")
			continue
		}

		for x := range n - k + 1 {
			for y := range m - k + 1 {
				switch xCheck(getMainDiagonal(board, k, x, y), k) {
				case "WIN":
					winFlag = true
				case "YES":
					yesFlag = true
				}
				if winFlag {
					break
				}
			}
			if winFlag {
				break
			}
		}
		if winFlag {
			fmt.Fprintln(out, "NO")
			continue
		}

		for x := range n - k + 1 {
			for y := k - 1; y < m; y++ {
				switch xCheck(getSecondaryDiagonal(board, k, x, y), k) {
				case "WIN":
					winFlag = true
				case "YES":
					yesFlag = true
				}
				if winFlag {
					break
				}
			}
			if winFlag {
				break
			}
		}
		if winFlag {
			fmt.Fprintln(out, "NO")
			continue
		}

		if yesFlag {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func getHorizontalLine(board [][]byte, k, x, y int) (tmp []byte) {
	for j := 0; j < k; j++ {
		tmp = append(tmp, board[x][y+j])
	}
	return tmp
}

func getVerticalLine(board [][]byte, k, x, y int) (tmp []byte) {
	for j := 0; j < k; j++ {
		tmp = append(tmp, board[j+x][y])
	}
	return tmp
}

func getMainDiagonal(board [][]byte, k, x, y int) (tmp []byte) {
	for range k {
		tmp = append(tmp, board[x][y])
		y++
		x++
	}
	return tmp
}

func getSecondaryDiagonal(board [][]byte, k, x, y int) (tmp []byte) {
	for range k {
		tmp = append(tmp, board[x][y])
		y--
		x++
	}
	return tmp
}

func countByteInArray(arr []byte, c byte) (count int) {
	for _, a := range arr {
		if a == c {
			count++
		}
	}
	return count
}

func xCheck(line []byte, k int) string {
	countX := countByteInArray(line, byte('X'))
	countO := countByteInArray(line, byte('O'))
	if countX == k || countO == k {
		return "WIN"
	} else if countX == k-1 && countO == 0 {
		return "YES"
	} else {
		return "NO"
	}
}
