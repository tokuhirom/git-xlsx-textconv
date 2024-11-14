package main

import (
	"bytes"
	"testing"
)

func TestConv(t *testing.T) {
	var buf bytes.Buffer
	err := textconv("testfile.xlsx", &buf)
	if err != nil {
		t.Fatal(err)
	}

	want := "" +
		"[Tabelle1] Foo\tBar\t\n" +
		"[Tabelle1] Baz\tQuuk\tLOWER(B2)\n" +
		"[Tabelle2] \n" +
		"[Tabelle2] \n" +
		"[Tabelle2] \n" +
		"[Tabelle2] \n"

	got := buf.String()
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}
