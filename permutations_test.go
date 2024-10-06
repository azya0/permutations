package permutations

import "testing"

func CompareArrray[T Compared](array1, array2 []T) bool {
	if len(array1) != len(array2) {
		return false
	}

	for index := range len(array1) {
		if array1[index] != array2[index] {
			return false
		}
	}

	return true
}

func CompareMatrix[T Compared](matrix1, matrix2 [][]T) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}

	for index := range len(matrix1) {
		if !CompareArrray(matrix1[index], matrix2[index]) {
			return false
		}
	}

	return true
}

func Comparison[T Compared](t *testing.T, got, want [][]T) {
	if !CompareMatrix(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestString1(t *testing.T) {
	got := Permutations([]string{"a", "b", "c"})

	want := [][]string{{"a", "b", "c"}, {"b", "a", "c"}, {"c", "a", "b"}, {"a", "c", "b"}, {"b", "c", "a"}, {"c", "b", "a"}}

	Comparison(t, got, want)
}

func TestString2(t *testing.T) {
	got := SortedPermutations([]string{"a", "b", "c"})

	want := [][]string{{"a", "b", "c"}, {"a", "c", "b"}, {"b", "a", "c"}, {"b", "c", "a"}, {"c", "a", "b"}, {"c", "b", "a"}}

	Comparison(t, got, want)
}

func TestString3(t *testing.T) {
	got := SortedPermutations([]string{"c", "a", "b"})

	want := [][]string{{"a", "b", "c"}, {"a", "c", "b"}, {"b", "a", "c"}, {"b", "c", "a"}, {"c", "a", "b"}, {"c", "b", "a"}}

	Comparison(t, got, want)
}

func TestInt1(t *testing.T) {
	got := Permutations([]int{1, 2, 3})

	want := [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2,}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}}

	Comparison(t, got, want)
}

func TestInt2(t *testing.T) {
	got := SortedPermutations([]int{1, 2, 3})
	
	want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}

	Comparison(t, got, want)
}

func TestInt3(t *testing.T) {
	got := SortedPermutations([]int{3, 2, 1})
	
	want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}

	Comparison(t, got, want)
}
