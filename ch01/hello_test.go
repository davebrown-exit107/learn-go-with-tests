package hello_test

import (
	"fmt"
	"testing"

	hello "github.com/davebrown-exit107/learn-go-with-tests/ch01"
)

func TestHello(t *testing.T) {
	t.Run("empty string defaults to world", func(t *testing.T) {
		got := hello.Hello("", hello.EN)
		want := "hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with a name", func(t *testing.T) {
		got := hello.Hello("whole world", hello.EN)
		want := "hello, whole world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("use a string instead of the included constants", func(t *testing.T) {
		got := hello.Hello("world", "spanish")
		want := "hola, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := hello.Hello("world", hello.ES)
		want := "hola, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got := hello.Hello("world", hello.FR)
		want := "bonjour, world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q got %q", want, got)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello.Hello("world", hello.EN)
	}
}

func ExampleHello() {
	msg := hello.Hello("World", hello.EN)
	fmt.Println(msg)
	// Output: hello, World
}
