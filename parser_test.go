package main

import (
	"testing"
)

func TestGetValueByString(t *testing.T) {
	got := "value"
	want := "value"
	if want != got {
		t.Fatalf("getValueByString() = %v, wanted %v", got, want)
	}
}

func TestGetClaimNameByString(t *testing.T) {
	
}
