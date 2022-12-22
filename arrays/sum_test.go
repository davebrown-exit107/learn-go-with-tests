package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSum() {
	// Sum the contents of an array of integers
	numbers := []int{1, 2, 3, 4, 5}

	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func ExampleSumAll() {
	// For each array, calculate the sum of the members. Return an array of the sums
	numbers1 := []int{1, 2}
	numbers2 := []int{0, 9}

	sum := SumAll(numbers1, numbers2)
	fmt.Println(sum)
	// Output: [3 9]
}

func ExampleSumAllTails() {
	// For each array, calculate the sum of the members, excluding the head. Return an array of the sums
	numbers1 := []int{1, 2}
	numbers2 := []int{0, 9}

	sum := SumAllTails(numbers1, numbers2)
	fmt.Println(sum)
	// Output: [2 9]
}
