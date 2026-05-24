package internal

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Alice")
	want := "Hello, Alice!"

	if got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}
