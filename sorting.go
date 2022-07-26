package sorting

import "golang.org/x/exp/constraints"

func swap[T any](array []T, x int, y int) {
	temp := array[x]
	array[x] = array[y]
	array[y] = temp
}

func IsSorted[T constraints.Ordered](array []T) bool {
	n := len(array)
	for i := 0; i < n-1; i++ {

		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}

//	Perform *insertion* sorting method within the provided array with *comparable* (`constraints.Ordered`) elements
func InsertionSort[T constraints.Ordered](array []T) {
	n := len(array) //	Length of `array`

	var j int //	Initialize `j` indexing variable
	var k T   //	Initialize `k` comparable variable

	//	Begin loop at the first index with `j` behind value of `k`
	for i := 1; i < n; i++ {
		k = array[i] //	Set `k` to current value

		//	Continue to forward any values smaller than `k`
		for j = i - 1; j >= 0 && array[j] > k; j-- {
			array[j+1] = array[j]
		}
		//	Set the value of the previous index of `j` to `k` if `j` has changed
		if j < i-1 {
			array[j+1] = k
		}
	}
}
