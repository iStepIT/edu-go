package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	genNumsChan := make(chan int)
	squareChan := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			num := rand.Intn(10)
			genNumsChan <- num
		}
		close(genNumsChan)
	}()

	go func() {
		defer wg.Done()
		for num := range genNumsChan {
			squareChan <- num * num
		}
	}()

	go func() {
		wg.Wait()
		close(squareChan)
	}()

	var squares []int
	for square := range squareChan {
		squares = append(squares, square)
	}

	fmt.Println(squares)
}
