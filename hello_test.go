package main

import "testing"

func TestHello(t *testing.T) {
	var s = hello()
	if s != "Hello World!" {
		t.Error("Incorrect string value ", s)
	}
}
