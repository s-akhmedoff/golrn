package exercise

import (
	"fmt"
)

type cubic struct {
	words [6]byte
}

func (c cubic) isIn(b byte) bool {
	for _, v := range c.words {
		if b == v {
			return true
		}
	}
	return false
}

func Cubics() {

	cbs := make([]cubic, 4)

	cbs[0].words = [6]byte{'a', 'c', 'e', 'f', 'g', 'h'}
	cbs[1].words = [6]byte{'a', 'c', 'c', 'b', 'w', 't'}
	cbs[2].words = [6]byte{'n', 'e', 'o', 'z', 'l', 'l'}
	cbs[3].words = [6]byte{'v', 'a', 'o', 't', 'r', 'e'}

	var bts []byte

	var str string
	str = "hello"
	for i := 0; i < len(str); i++ {
		bts = append(bts, str[i])
	}

	check := func(arr []int, b int) bool {
		for _, v := range arr {
			if b == v {
				return true
			}
		}
		return false
	}

	where := func() (a []int) {
		for i1, v1 := range bts {
			for i := range cbs {
				if cbs[i].isIn(v1) {
					if !check(a, i) {
						a = append(a, i)
					}
					bts = append(bts[:i1], bts[i1+1:]...)
				}
				break
			}
		}
		return a
	}

	fmt.Println(where())

}

// func read
//for i := 0; i < 4; i++ {
//	cbs[i] = cubic{}
//	for j := 0; j < 6; j++ {
//		reader := bufio.NewReader(os.Stdin)
//		t, _ := reader.ReadByte()
//		cbs[i].words[j] = t
//	}
//	fmt.Println("New Cubic")
//}
