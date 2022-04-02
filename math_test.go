package main
import "testing"

func TestSum(t *testing.T) {
	if Sum(10, 10) != 20 {
		t.Error("Result should be 20")
	}
}