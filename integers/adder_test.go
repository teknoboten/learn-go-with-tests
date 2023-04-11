package integers

import (
	"fmt"
	"testing"
)

//this is example function!
//examples are executable snippets of code that help keep docs up to date
//they take no args, begin with Example and live in a packages _test.go file
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2,2)
	expected := 4

	if sum != expected {
		//use %d to print an integer instead of string
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}