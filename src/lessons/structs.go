package lessons

import (
	"errors"
)

const (
	TEXT = iota
	FILE
	BINARY
)

type data struct {
	size   int
	typeof int
}

type data_arr []data

func (d data_arr) create(size int) error {
	if d == nil {
		d = make([]data, 10)
		return nil
	}

	return errors.New("error: already initialized")
}

func (d data_arr) addByValue() () {

}

func (d data_arr) addByPointer() () {

}

func Structures() {

}
