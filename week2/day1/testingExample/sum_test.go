package testingExample

import "testing"

func TestAverage(t *testing.T) {
	var sum int
	sum = Sum(2, 3)
	if sum != 5 {
		t.Error("Expected 5 but got", sum)
	}
}
