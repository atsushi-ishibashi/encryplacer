package main

import "testing"

func Test_Filter_match(t *testing.T) {
	f := Filter{}
	if !f.match("hoge") {
		t.Error("error")
	}

	f.Suffix = "ge"
	if !f.match("hoge") {
		t.Error("error")
	}
	if f.match("hogu") {
		t.Error("error")
	}

	f.Contain = "og"
	if !f.match("hoge") {
		t.Error("error")
	}
	if f.match("huge") {
		t.Error("error")
	}
}
