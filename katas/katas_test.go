package main

import (
	"testing"
)

func TestDNAtoRNA(t *testing.T) {
	got := DNAtoRNA("GCAT")
	want := "GCAU"
	assertString(t, got, want)
}

func TestMove(t *testing.T) {
	got := Move(2, 6)
	want := 14
	assertInt(t, got, want)
}

func TestAbbrevName(t *testing.T) {
	got := AbbrevName("fjord winstar")
	want := "F.W"
	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
