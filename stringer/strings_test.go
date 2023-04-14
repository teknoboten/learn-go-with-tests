package stringer

import (
	"fmt"
	"testing"
)

// stringer takes in a string, a word to replace and what to replace it with
// converts it to lower case
// replaces any occurances of the old word with the new word
// then returns the new string

func ExampleStringer() {
	result := Stringer("THIS is a dang fine, albeit SOMEWHAT POINTLESS, string!", "dang", "very")
	fmt.Println(result)
	// Output: this is a very fine, albeit somewhat pointless, string!
}

func TestStringer(t *testing.T) {
	t.Run("replace all the cuss words", func(t *testing.T) {
		got	:= Stringer("this is flippin awesome!", "flippin", "frackin")
		want := "this is frackin awesome!"
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("stop all the yelling", func(t *testing.T) {
		got	:= Stringer("HEY STOP ALL THE flippin YELLING!", "flippin", "frackin")
		want := "hey stop all the frackin yelling!"
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})


}


// func TestStringer(t *testing.T) {
// 	got	:= Stringer("Boop", "bOOp")
// 	want := true

// 	if got != want {
// 		t.Errorf("got %t want %t", got, want)
// 	}
// }

// func TestStringer(t *testing.T) {
// 	got	:= Stringer("Boop", "bOOp")
// 	want := true

// 	if got != want {
// 		t.Errorf("got %t want %t", got, want)
// 	}
// }


// func TestStringer(t *testing.T) {
// 	t.Run("compare two matching strings", func(t *testing.T) {
// 		got	:= Stringer("boop", "boop")
// 		want := "match!"
// 			if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// })
// t.Run("compare two non-matching strings", func(t *testing.T) {
// 		got	:= Stringer("boop", "zap")
// 		want := "not a match!"
	 	
// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// })
// }