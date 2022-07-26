package sorting

import "testing"

func Test1(t *testing.T) {
	x := []uint8{2, 1, 5, 4, 3}
	InsertionSort(x)
	if !IsSorted(x) {
		t.Fatalf("failed to sort")
	}
}

func Test2(t *testing.T) {
	x := []uint8{3, 244, 8, 33, 3}
	InsertionSort(x)
	if !IsSorted(x) {
		t.Fatalf("failed to sort")
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
