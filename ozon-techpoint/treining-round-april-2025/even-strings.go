// Полностью решена
package ozontechpoint

import (
	"bufio"
	"fmt"
	"os"
)

func EvenStrings() {
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
		var arr []string
		dict := make(map[string]int)
		dictParts := make(map[string][2]string)
		var n int
		_, err = fmt.Fscan(in, &n)
		if err != nil {
			fmt.Println(err)
		}

		for range n {
			var s string
			_, err = fmt.Fscan(in, &s)
			if err != nil {
				fmt.Println(err)
			}

			if _, ok := dict[s]; ok {
				dict[s] += 1
			} else {
				dict[s] = 1
				var tmp [2]string
				tmp[0], tmp[1] = stringSplit(s)
				dictParts[s] = tmp
				arr = append(arr, s)
			}
		}

		count := 0
		if len(arr) > 1 {
			for i := 0; i < len(arr)-1; i++ {
				strCount := 0
				for j := i + 1; j < len(arr); j++ {
					if (len(dictParts[arr[i]][0]) > 0 && dictParts[arr[i]][0] == dictParts[arr[j]][0]) ||
						(len(dictParts[arr[i]][1]) > 0 && dictParts[arr[i]][1] == dictParts[arr[j]][1]) {
						strCount += dict[arr[j]]
					}
				}
				count += strCount * dict[arr[i]]
			}
		}

		for _, v := range dict {
			count += v * (v - 1) / 2
		}

		fmt.Fprintln(out, count)
	}
}

func stringSplit(s string) (string, string) {
	bytes := []byte(s)
	evenBytes := make([]byte, (len(bytes)+1)/2)
	oddBytes := make([]byte, len(bytes)/2)

	evenIdx, oddIdx := 0, 0
	for i := 0; i < len(bytes); i++ {
		if i%2 == 0 {
			evenBytes[evenIdx] = bytes[i]
			evenIdx++
		} else {
			oddBytes[oddIdx] = bytes[i]
			oddIdx++
		}
	}
	return string(evenBytes), string(oddBytes)
}

// Конкатенация строк гораздо медленнее побайтовых операций.
// Также, так как память в массивах байт предвыделена,
// это позволяет избежать трат времени на переопределение массивов

// Разница, продемонстрирована на скришоте.
// Верхний вариант через массив байт,
// Нижний - через конкатенацию строк.

// Простой вариант с конкатенация строк:
// func longStringSplit(s string) (even string, odd string) {
// 	for i := 0; i < len(s); i += 2 {
// 		even += string(s[i])
// 		if i+1 < len(s) {
// 			odd += string(s[i+1])
// 		}
// 	}
// 	return even, odd
// }
