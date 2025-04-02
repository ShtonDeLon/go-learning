// Полностью решена
package ozontechpoint

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Bank struct {
	TO_FROM [6]float32
}

func ThreeBanks() {
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

	var B [3]Bank
	for range t {
		for i := range 3 {
			var from, to int
			for j := range 6 {
				_, err = fmt.Fscan(in, &from, &to)
				if err != nil {
					fmt.Println(err)
				}
				B[i].TO_FROM[j] = float32(to) / float32(from)
			}
		}

		s := fmt.Sprintf("%f", bestStrategy(B))
		fmt.Fprintln(out, s)
	}
}

func bestStrategy(B [3]Bank) float32 {
	var strategy [27]float32

	for i := range 3 {
		strategy[i*9] = B[0].TO_FROM[0]
		strategy[i*9+1] = B[0].TO_FROM[1] * B[1].TO_FROM[5]
		strategy[i*9+2] = B[0].TO_FROM[1] * B[2].TO_FROM[5]
		strategy[i*9+3] = B[0].TO_FROM[0] * B[1].TO_FROM[3] * B[2].TO_FROM[5]
		strategy[i*9+4] = B[0].TO_FROM[0] * B[2].TO_FROM[3] * B[1].TO_FROM[5]
		strategy[i*9+5] = B[0].TO_FROM[0] * B[1].TO_FROM[2] * B[2].TO_FROM[0]
		strategy[i*9+6] = B[0].TO_FROM[0] * B[2].TO_FROM[2] * B[1].TO_FROM[0]
		strategy[i*9+7] = B[0].TO_FROM[1] * B[1].TO_FROM[4] * B[2].TO_FROM[0]
		strategy[i*9+8] = B[0].TO_FROM[1] * B[2].TO_FROM[4] * B[1].TO_FROM[0]

		tmp := B[0]
		B[0] = B[1]
		B[1], B[2] = B[2], tmp
	}

	slices.SortStableFunc(strategy[:], func(a, b float32) int {
		if a == b {
			return 0
		} else {
			if a > b {
				return -1
			} else {
				return 1
			}
		}
	})
	return strategy[0]
}
