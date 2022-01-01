package main

import (
	"fmt"
	"sort"
)

func main() {

	// https://pkg.go.dev/sort#Float64s
	example_Float64s()

	// https://pkg.go.dev/sort#Float64sAreSorted
	example_Float64sAreSorted()

	// https://pkg.go.dev/sort#Ints
	example_Ints()
}

func example_Float64s() {
	var list = []float64{5.0, 2.3, 1.4, 0.5, 3.2}
	// before sort
	fmt.Println(list)

	// sort
	sort.Float64s(list)

	// after sort
	fmt.Println(list)
}

func example_Float64sAreSorted() {
	// unsorted list.
	var list = []float64{5.0, 2.3, 1.4, 0.5, 3.2}
	result := sort.Float64sAreSorted(list)
	// expect false
	fmt.Println(result)

	// unsorted list.
	var list2 = []float64{1.1, 1.2, 1.3}
	result2 := sort.Float64sAreSorted(list2)
	// expect false
	fmt.Println(result2)

}

func example_Ints() {
	// unsorted list
	var list = []int{5,4,3,2,1}
	// before sort.
	fmt.Println(list)

	// sort the list.
	sort.Ints(list)

	// sorted list.
	fmt.Println(list)
}

