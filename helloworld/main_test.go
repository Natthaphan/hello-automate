package main

import "testing"

func Test_hello_world(t *testing.T) {
	expected := "Hello, World!"

	if actual := GetHello(); expected != actual {
		t.Errorf("expected: %s but got %s", expected, actual)
	}
}
