# Permutations module
Go/Golang module

## Add it to your Golang project:
```go get github.com/azya0/permutations```

## Description
This module is used for all possible permutations among the elements of the array

For example, we can represent the array [1 2 3] as:
* [1 2 3]
* [1 3 2]
* [2 1 3]
* [2 3 1]
* [3 1 2]
* [3 2 1]


## Usage
Within the framework of this module, the user is invited to use the following 2 functions:
<br><br>
```func Permutations[T Compared](input []T) [][]T``` <br>

Accepts an input array as input. Returns a matrix of all possible permutations

```func SortedPermutations[T Compared](input []T) [][]T```

Accepts an input array as input. Returns a matrix of all possible permutations in sorted form.

## T Compared
This is the type of data that Golang can compare with each other by default

```
// Type of types that can be compared
type Compared interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64 | ~string
}
```
