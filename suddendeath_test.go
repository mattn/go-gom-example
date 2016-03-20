package main

import (
	"github.com/mattn/gover"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestSimple(t *testing.T) {
	f, err := ioutil.TempFile("", "gom")
	if err != nil {
		t.Fatal(err)
	}
	oldstdout := os.Stdout
	os.Stdout = f
	suddenDeath("こんにちわ世界")
	os.Stdout = oldstdout
	f.Close()
	b, err := ioutil.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	value := string(b)
	expected := `
＿人人人人人人人人人＿
＞　こんにちわ世界　＜
￣ＹＹＹＹＹＹＹＹＹ￣
`[1:]
	expected = regexp.MustCompile("\x1b[[^m]*m").ReplaceAllString(expected, "")

	if value != expected {
		t.Fatalf("Expected %v, but %v:", value, expected)
	}
}

func TestDependency(t *testing.T) {
	v := gover.Version()
	if gover.Version() == "" {
		t.Fatalf("Expected empty string, but %v:", v)
	}
}
