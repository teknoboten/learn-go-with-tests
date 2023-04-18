package main

import (
	"testing"
)

func TestDNAtoRNA(t *testing.T) {

	got := DNAtoRNA("GCAT")
	want := "GCAU"
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
