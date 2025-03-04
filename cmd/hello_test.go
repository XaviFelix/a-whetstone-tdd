package main

import (
	"testing"
)

// Creating subtets withing a test.
// Grouping tests around a 'thing' having subsets describing different scenarios
func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Marco", "French")
		want := "Bonjour, Marco"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Goku", "Japanese")
		want := "Kon'nichiwa, Goku"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// tells test suite that this method is a helper function
	// when it fails, the line number reported will be in our function call rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
