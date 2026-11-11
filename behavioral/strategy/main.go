package main

import "fmt"

type SortStrategy interface {
	Sort(array []int) []int
}

type BubbleSortStrategy struct{}

var _ SortStrategy = &BubbleSortStrategy{}

func (bs *BubbleSortStrategy) Sort(array []int) []int {
	fmt.Println("Sorting using bubble sort")
	// 执行排序
	return array
}

type QuickSortStrategy struct{}

var _ SortStrategy = &QuickSortStrategy{}

func (qs *QuickSortStrategy) Sort(array []int) []int {
	fmt.Println("Sorting using quick sort")
	// 执行排序
	return array
}

type Sorter struct {
	sorterSmall SortStrategy
	sorterBig   SortStrategy
}

func NewSorter(sorterSmall, sorterBig SortStrategy) *Sorter {
	return &Sorter{
		sorterSmall: sorterSmall,
		sorterBig:   sorterBig,
	}
}

func (s *Sorter) Sort(array []int) []int {
	if len(array) > 5 {
		return s.sorterBig.Sort(array)
	}
	return s.sorterSmall.Sort(array)
}

func main() {
	small := []int{1, 3, 4, 2}
	big := []int{1, 4, 3, 2, 8, 10, 5, 6, 9, 7}

	sorter := NewSorter(&BubbleSortStrategy{}, &QuickSortStrategy{})

	sorter.Sort(small) // 输出 : Sorting using bubble sort
	sorter.Sort(big)   // 输出 : Sorting using quick sort
}
