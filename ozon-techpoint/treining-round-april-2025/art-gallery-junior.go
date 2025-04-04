// Частично решена
// Проблема в алгоритме, вылезает за временные рамки
package ozontechpoint

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Object struct {
	sizeX, sizeY  int //max, min
	storageNumber int
	picsCount     int
}

func (o *Object) setup() {
	if o.sizeX < o.sizeY {
		o.sizeX, o.sizeY = o.sizeY, o.sizeX
	}
}

func cmpSq(a Object, b Object) int {
	s1 := a.sizeX * a.sizeY
	s2 := b.sizeX * b.sizeY
	if s1 < 0 || s2 < 0 {
		panic("negative square")
	} else {
		if s1 > s2 {
			return -1
		} else {
			if a.sizeX*a.sizeY < b.sizeX*b.sizeY {
				return 1
			} else {
				return 0
			}
		}
	}
}

func ArtGallery() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	_, err := fmt.Fscan(in, &t)
	if err != nil {
		fmt.Fprintln(out, err)
	}

	for range t {
		var n int
		_, err = fmt.Fscan(in, &n)
		if err != nil {
			fmt.Fprintln(out, err)
		}

		inputBoxes := make(map[Object]bool, n)
		for range n {
			var tmp Object
			_, err = fmt.Fscan(in, &tmp.sizeX, &tmp.sizeY)
			if err != nil {
				fmt.Fprintln(out, err)
			}
			tmp.setup()

			if _, ok := inputBoxes[tmp]; !ok {
				inputBoxes[tmp] = true
			}
		}

		counter := 0
		boxes := make([]Object, len(inputBoxes))
		for b := range inputBoxes {
			boxes[counter] = b
			counter++
		}
		slices.SortStableFunc(boxes, cmpSq)

		var m int
		_, err = fmt.Fscan(in, &m)
		if err != nil {
			fmt.Fprintln(out, err)
		}

		pics := make([]Object, m)
		inputPics := make(map[Object]bool, m)
		counter = 0
		for range m {
			var tmp Object
			_, err = fmt.Fscan(in, &tmp.sizeX, &tmp.sizeY)
			if err != nil {
				fmt.Fprintln(out, err)
			}
			tmp.setup()

			if _, ok := inputPics[tmp]; !ok {
				inputPics[tmp] = true
			}
		}

		for p := range inputPics {
			pics[counter] = p
			counter++
		}

		slices.SortStableFunc(pics, cmpSq)
		fmt.Fprintln(out, planSearch(boxes, pics))
	}
}

func planSearch(boxes []Object, pics []Object) int {
	storage := make([]Object, len(boxes))
	boxCount := 0

	err := false
	for ip := 0; ip < len(pics); ip++ {
		inserted := false
		for i := 0; i < boxCount; i++ {
			if storage[i].sizeX >= pics[ip].sizeX && storage[i].sizeY >= pics[ip].sizeY {
				inserted = true
				pics[ip].storageNumber = i
				storage[i].picsCount++
				break
			}
		}
		if inserted {
			continue
		}

		for i := 0; i < len(boxes); i++ {
			if boxes[i].sizeX < pics[ip].sizeX && boxes[i].sizeY < pics[ip].sizeY {
				err = true
				break
			}
			if boxes[i].sizeX >= pics[ip].sizeX && boxes[i].sizeY >= pics[ip].sizeY {
				storage[boxCount] = boxes[i]
				storage[boxCount].picsCount++
				pics[ip].storageNumber = boxCount
				boxCount++
				break
			}
		}
		if err {
			break
		}
	}

	if err {
		return -1
	} else {

		for ip := range pics {
			curStorageNumber := pics[ip].storageNumber
			ln := boxCount - 1
			for i := ln; i > curStorageNumber; i-- {
				if storage[i].sizeX >= pics[ip].sizeX && storage[i].sizeY >= pics[ip].sizeY {
					pics[ip].storageNumber = i
					storage[i].picsCount++
					storage[curStorageNumber].picsCount--
					if storage[curStorageNumber].picsCount == 0 {
						boxCount--
					}
					break
				}
			}
		}

		if boxCount == 0 {
			return -1
		} else {
			return boxCount
		}
	}
}
