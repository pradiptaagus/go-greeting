package go_greeting

import "testing"

func TestSayHello(t *testing.T) {
	result := SayHello("John Doe")
	if result != "Hello John Doe" {
		t.Errorf("Expected to be \"Hello Johh Doe\", got \"%s\"", result)
	}
}
