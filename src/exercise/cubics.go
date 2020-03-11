package exercise

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cubic struct {
	words []byte
}

func (c cubic) isIn(b byte) bool {
	for _, v := range c.words {
		if b == v {
			return true
		}
	}
	return false
}

func fillUp(source []cubic, strs []string) {
	for i := range source {
		source[i].words = []byte(strings.Replace(strs[i], " ", "", -1))
	}
}

func getWord() string {
	fmt.Println("Write a word: ")
	temp, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Replace(temp, "\n", "", -1)
}

func getCubicNumbers() (int, error) {
	fmt.Println("Write cubic numbers: ")
	temp, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strconv.Atoi(strings.Replace(temp, "\n", "", -1))
}

func readLinesToCubic(cubcs []cubic, count_cubics int) {
	strs := make([]string, count_cubics)
	fmt.Println("Write lines: ")
	for i := range strs {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		strs[i] = text
	}

	fillUp(cubcs, strs)
}

func do(wrd []byte, cbcs []cubic, n int) (bool, []int) {
	i, j := len(wrd)-1, 0
	var arr []int

	for ; i != 0; {
		if cbcs[j].isIn(wrd[i]) {
			arr = append(arr, j)

		}
	}

	return false, nil
}

func Cubics() {

	word := getWord()
	n, _ := getCubicNumbers()

	word_latters := []byte(word)
	cubics := make([]cubic, n)

	readLinesToCubic(cubics, n)



	fmt.Println(word)
	fmt.Println(word_latters)
	fmt.Println(cubics)

}
