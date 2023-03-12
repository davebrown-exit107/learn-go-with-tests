package arrays_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := arrays.Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	checkArraySums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("two full arrays", func(t *testing.T) {
		got := arrays.SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkArraySums(t, got, want)
	})

	t.Run("one full array and one empty array", func(t *testing.T) {
		got := arrays.SumAll([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkArraySums(t, got, want)
	})

	t.Run("two empty arrays", func(t *testing.T) {
		got := arrays.SumAll([]int{}, []int{})
		want := []int{0, 0}

		checkArraySums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	checkArraySums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("two full arrays", func(t *testing.T) {
		got := arrays.SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkArraySums(t, got, want)
	})

	t.Run("one full array and one empty array", func(t *testing.T) {
		got := arrays.SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkArraySums(t, got, want)
	})

	t.Run("two empty arrays", func(t *testing.T) {
		got := arrays.SumAllTails([]int{}, []int{})
		want := []int{0, 0}

		checkArraySums(t, got, want)
	})
}

func ExampleSum() {
	// Sum the contents of an array of integers
	numbers := []int{1, 2, 3, 4, 5}

	sum := arrays.Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func ExampleSumAll() {
	// For each array, calculate the sum of the members. Return an array of the sums
	numbers1 := []int{1, 2}
	numbers2 := []int{0, 9}

	sum := arrays.SumAll(numbers1, numbers2)
	fmt.Println(sum)
	// Output: [3 9]
}

func ExampleSumAllTails() {
	// For each array, calculate the sum of the members, excluding the head. Return an array of the sums
	numbers1 := []int{1, 2}
	numbers2 := []int{0, 9}

	sum := arrays.SumAllTails(numbers1, numbers2)
	fmt.Println(sum)
	// Output: [2 9]
}
