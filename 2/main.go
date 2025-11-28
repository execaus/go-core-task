package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

const (
	randomSize          = 10
	randomMaxValue      = 99
	addedElement        = 999
	removedElementIndex = 2
)

func main() {
	originalSlice, err := randomSlice(randomSize, randomMaxValue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Original slice:", originalSlice)

	evenSlice, err := sliceExample(originalSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Even slice:", evenSlice)

	addedSlice, err := addElements(originalSlice, addedElement)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Slice after adding %v: %v,", addedElement, addedSlice))

	copiedSlice, err := copySlice(originalSlice)
	if err != nil {
		log.Fatal(err)
	}
	originalSlice[0] = originalSlice[0] + addedElement
	fmt.Println("Modified original slice:", originalSlice)
	fmt.Println("Copied slice:", copiedSlice)

	removedSlice, err := removeElement(originalSlice, removedElementIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Slice after removing index %v: %v", removedElementIndex, removedSlice))
}

// randomSlice создает срез случайных целых чисел заданного размера.
func randomSlice(size, maxValue int) ([]int, error) {
	if size < 0 {
		return nil, errors.New("size cannot be negative")
	}
	if size == 0 {
		return nil, errors.New("size cannot be zero")
	}

	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(maxValue)
	}

	return slice, nil
}

// sliceExample возвращает новый срез, содержащий только четные элементы из входного среза.
func sliceExample(in []int) ([]int, error) {
	if in == nil {
		return nil, errors.New("input slice cannot be nil")
	}

	var evens []int
	for _, v := range in {
		if v&1 == 0 {
			evens = append(evens, v)
		}
	}

	return evens, nil
}

// addElements добавляет элемент num в конец среза in.
func addElements(in []int, num int) ([]int, error) {
	if in == nil {
		return nil, errors.New("input slice cannot be nil")
	}

	return append(in, num), nil
}

// copySlice создает копию входного среза in.
func copySlice(in []int) ([]int, error) {
	if in == nil {
		return nil, errors.New("input slice cannot be nil")
	}

	out := make([]int, len(in))
	copy(out, in)

	return out, nil
}

// removeElement удаляет элемент по индексу index из среза in.
func removeElement(in []int, index int) ([]int, error) {
	if in == nil {
		return nil, errors.New("input slice cannot be nil")
	}
	if index < 0 || index >= len(in) {
		return nil, errors.New("index out of range")
	}

	return append(in[:index], in[index+1:]...), nil
}
