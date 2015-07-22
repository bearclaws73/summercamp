package example

import "testing"

// func TestHello(t *testing.T) {
// 	if Hello() != "Hello World" {
// 		t.Error("Expected 'Hello World' but got ", Hello())
// 	}
// }
//
func TestHello(t *testing.T) {
	s := Hello()
	if s != "Hello World" {
		t.Error("Expected 'Hello World' but got ", s)
	}
}
