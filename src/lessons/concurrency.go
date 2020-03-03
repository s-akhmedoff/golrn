package lessons

import (
	"fmt"
	"strconv"
	"sync"
)

func count(str string, w *sync.WaitGroup) {
	defer w.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(str, ":", i)
	}
}

func echo(str []string, ch *chan string) {
	for _, v := range str {
		*ch <- v
	}
}

func Conc() {

	var wg sync.WaitGroup

	wg.Add(1)

	go count("string2", &wg)

	wg.Wait()

	fmt.Println("Go Concurrency")

	ch := make(chan string)

	go func() {ch <- "ping"}()

	ms := <- ch

	fmt.Println(ms)

	strs := make([]string, 5)
	for i := 0; i < len(strs); i++ {
		strs[i] = "it's number: " + strconv.Itoa(i)
	}

	ch1 := make(chan string, 5)
	go echo(strs, &ch1)

	for i := 0; i < len(strs); i++ {
		fmt.Println(<-ch1)
	}

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
