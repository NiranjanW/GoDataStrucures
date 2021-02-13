package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// list := []int{12, 4, 2, 59, 100}
	slice := generateSlice(20)
	fmt.Println("\n --Unsorted---\n\n", slice)
	mergeSort(slice)
	fmt.Println("\n --Sorted --\n\n", mergeSort(slice))
	// fmt.Println("\n --Sorted --\n\n", mergeSort(list))
}

func bubbleSort(arry []int) []int {
	if len(arry) < 2 {
		return arry
	}

	for i := len(arry); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arry[j-1] > arry[j] {
				arry[j-1], arry[j] = arry[j], arry[j-1]
			}
		}
	}
	return arry
}

func bubbleSort1(arry []int) []int {
	if len(arry) < 2 {
		return arry
	}
	swap := true

	for swap {
		swap = false
		for i := 1; i < len(arry); i++ {
			if arry[i-1] > arry[i] {
				arry[i-1], arry[i] = arry[i], arry[i-1]
				swap = true
			}

		}
	}
	return arry
}

func mergeSort(arry []int) []int {
	if len(arry) < 2 {
		return arry
	}
	mid := int(len(arry) / 2)
	left, right := arry[mid:], arry[:mid]

	return merge(mergeSort(left), mergeSort(right))

}

func merge(a, b []int) []int {
	size, i, j := len(a)+len(b), 0, 0
	c := make([]int, size)

	for k := 0; k < size; k++ {
		switch true {
		case i == len(a):
			c[k] = b[j]
			j++

		case j == len(b):
			c[k] = a[i]
			i++

		case a[i] > b[j]:
			c[k] = b[j]
			j++

		case a[i] < b[j]:
			c[k] = a[i]
			i++

		case a[i] == b[j]:
			c[k] = a[i]
			i++

		}

	}
	return c
}

// func merge(a, b []int) []int {
// 	size, i, j := len(a)+len(b), 0, 0
// 	result := make([]int, size)
// 	for k := 0; k < size; k++ {
// 		switch true {
// 		case i == len(a):
// 			//all the elements of a already been judged
// 			//assuming a and b both are sorted, this happens because
// 			//some cases will have not equal a and b, so it might
// 			// be a possibility that one array got finished earlier.
// 			result[k] = b[j]
// 			j += 1
// 		case j == len(b):
// 			//alll the elements of a already been judged
// 			//assuming a nd b both are sorted
// 			result[k] = a[i]
// 			i += 1
// 		case a[i] > b[j]:
// 			result[k] = b[j]
// 			//increment the index of b array because its present index
// 			//is already appended to the result array
// 			j += 1
// 		case a[i] < b[j]:
// 			//increment the index of a array because its present index
// 			//element is already appended to the result array
// 			result[k] = a[i]
// 			i += 1
// 		case a[i] == b[j]:
// 			//in case of equality
// 			result[k] = b[j]
// 			j += 1
// 		}
// 	}
// 	return result
// }

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999)
	}

	return slice
}
