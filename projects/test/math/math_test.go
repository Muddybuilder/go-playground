package math_test

import (
	"test-proj/math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Skip()
	total := math.Add(2, 2)
	if total != 4 {
		t.Errorf("Sum was incorrect, Actual: %d, Expected:%d", total, 4)
	}
	t.Log("Running sum test")

}
