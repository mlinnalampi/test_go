package main

import "testing"

func TestOrdinal(t *testing.T) {
	ord := ordinal(1)
	exp := "1st"
	if ord != exp {
		t.Error("Expected")
	}
	ord = ordinal(4)
	exp = "4th"
	if ord != exp {
		t.Error("Expected")
	}
	ord = ordinal(2)
	exp = "2nd"
	if ord != exp {
		t.Error("Expected")
	}
}
