package nnnnn

import (
	"fmt"
	"testing"
)

func TestNnnnnn(t *testing.T) {
	var a = []int{2, 3, 5, 7, 11, 13}
	var b = 2
	result := Nnnnnn(a, b)

	if result != 0 {
		fmt.Printf("%d\n", result)
		t.Errorf("this two are not the same")
	}

}
