package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2,4)
	expected := 6

	if sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1,3)
	fmt.Println(sum)
	//output: 4
}