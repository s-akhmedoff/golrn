package internship

import (
	"fmt"
	"github.com/srfrog/dict"
)

func GoMOD() {
	dict1 := dict.New(map[string]int{"one": 1})

	fmt.Println(dict1.Keys(), dict1.Values())
}
