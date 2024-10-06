package permutations

func Swap[T any](data []T, from, to int) {
	data[from], data[to] = data[to], data[from]
}

// One way list structure
type OWList[T any] struct {
	next *OWList[T];
	value T;
}

// Type of types that can be compared
type Compared interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64 | ~string
}

// By logic of func Permutations, len(value1) eq len(value2)
// So I won't add length comparing
func CustorCompare[T Compared](value1, value2 []T, strictly bool) bool {
	for index := range len(value1) {
		if value1[index] < value2[index] {
			return true
		} else if value1[index] > value2[index] {
			return false
		}
	}

	return !strictly
}

func AppendOWList[T Compared](root *OWList[[]T], value *OWList[[]T]) *OWList[[]T] {
	if root == nil {
		return value
	}

	if CustorCompare(value.value, root.value, true) {
		value.next = root

		return value
	}

	var current *OWList[[]T] = root;
	for (current.next != nil && CustorCompare(current.next.value, value.value, false)) {
		current = current.next
	}

	value.next = current.next
	current.next = value

	return root
}

func Permutations[T Compared](input []T) [][]T {
	var result [][]T = [][]T{input}

	data := make([]T, len(input))
	indexes := make([]int, len(input) + 1)

	for index, value := range input {
		indexes[index] = index
		data[index] = value
	}

	indexes[len(input)] = len(input)

	var index int = 1

	for index < len(input) {
		indexes[index] = indexes[index] - 1

		if index % 2 != 0 {
			Swap(data, indexes[index], index)
		} else {
			Swap(data, 0, index)
		}

		toSave := make([]T, len(data))
		copy(toSave, data)

		result = append(result, toSave)
		index = 1

		for indexes[index] == 0 {
			indexes[index] = index
			index++
		}
	}
	
	return result
}

func SortedPermutations[T Compared](input []T) [][]T {
	var result *OWList[[]T] = nil;
	result = AppendOWList(result, &OWList[[]T]{nil, input})
	var length int = 1;

	data := make([]T, len(input))
	indexes := make([]int, len(input) + 1)

	for index, value := range input {
		indexes[index] = index
		data[index] = value
	}

	indexes[len(input)] = len(input)

	var index int = 1

	for index < len(input) {
		indexes[index] = indexes[index] - 1

		if index % 2 != 0 {
			Swap(data, indexes[index], index)
		} else {
			Swap(data, 0, index)
		}

		toSave := make([]T, len(data))
		copy(toSave, data)

		result = AppendOWList(result, &OWList[[]T]{nil, toSave})
		length++
		index = 1

		for indexes[index] == 0 {
			indexes[index] = index
			index++
		}
	}

	toReturn := make([][]T, length)

	for index := range length {
		toReturn[index] = result.value;
		result = result.next;
	}

	return toReturn
}