package integers_test

import (
	"fmt"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/integers"
)

func TestAdder(t *testing.T) {
	got := integers.Adder(3, 3)
	want := 6

	if want != got {
		t.Errorf("wanted %d got %d", want, got)
	}

}

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		integers.Adder(1, 3)
	}
}

func ExampleAdder() {
	sum := integers.Adder(3, 3)
	fmt.Println(sum)
	// Output: 6
}
