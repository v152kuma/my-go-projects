package main

import (
	"fmt"
	"slices"
)

func main() {

	mapOfSlices := map[string][]string{
		"coffee": {"1", "2", "3"},
	}

	fmt.Println(mapOfSlices)

	mapOfSlices["tea"] = []string{"", "", ""}

	fmt.Println(mapOfSlices)

	arr := [3]int{1, 2, 3}

	fmt.Println(arr)

	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)
	slice = append(slice, 7, 8)
	fmt.Println(slice)
	slices.Delete(slice, 8, 9)

}
