package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Peter", "")
		want := "Hello, Peter"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got := Hello("Emanuelle", "French")
		want := "Bonjour, Emanuelle"
		assertCorrectMessage(t, got, want)
	})

}
