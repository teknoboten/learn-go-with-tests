package iterators

import (
	"fmt"
	"testing"
)

//Write ExampleRepeat to document your function
func ExampleRepeat() {
	repeated := Repeat("s", 4)
	fmt.Println(repeated)
	// Output: ssss
}

//Change the test so a caller can specify how many times the
//character is repeated and then fix the code
func TestRepeater(t *testing.T) {
	repeated := Repeat("s", 3)
	expected := "sss"

	if repeated != expected {
		t.Errorf("got %q but wanted %q", repeated, expected)
	}
}

//intial test
// func TestRepeater(t *testing.T) {
// 	repeated := Repeat("s")
// 	expected := "sssss"

// 	if repeated != expected {
// 		t.Errorf("got %q but wanted %q", repeated, expected)
// 	}
// }