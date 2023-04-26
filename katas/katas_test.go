package main

import (
	"reflect"
	"testing"
)

func TestRemoveChar(t *testing.T){
	got := RemoveChar("country")
	want := "ountr"
	assertString(t, got, want)
}

func TestSquareSum(t *testing.T) {
	got := SquareSum([]int{1,2,2})
	want := 9
	assertInt(t, got, want)
}

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

func TestInvert(t *testing.T) {
	got := Invert([]int{1, 2, 3, 4, 5})
	want := []int{-1, -2, -3, -4, -5}
	assertSlice(t, got, want)
}

func assertSlice(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) { // <- use reflect.deepEqual when output is a slice / array
		t.Errorf("got %v want %v", got, want)
	}
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
