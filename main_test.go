package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/daved/pythatchallenge/internal/daveddev"
	"github.com/daved/pythatchallenge/internal/djherbis"
	"github.com/daved/pythatchallenge/internal/jasonmoo"
	"github.com/daved/pythatchallenge/internal/srtkkou"
)

type IOTest struct {
	in  []string
	out []string
}

func testData() []*IOTest {
	return []*IOTest{
		{
			[]string{"one", "two", `"three`, "four", `five"`},
			[]string{"one", "two", "three four five"},
		},
		{
			[]string{`"one`, `two"`, "three"},
			[]string{"one two", "three"},
		},
		{
			[]string{"one", `"two"`},
			[]string{"one", "two"},
		},
		{
			[]string{"one", `"two`, "three", "four", "five", `six"`},
			[]string{"one", "two three four five six"},
		},
		{
			[]string{"cp", `"file`, "name", "with", `spaces"`, `"file`, "name", `backup"`},
			[]string{"cp", "file name with spaces", "file name backup"},
		},
	}
}

func TestDavedDev(t *testing.T) {
	for k, v := range testData() {
		p, err := ddpkg.Parse(v.in)
		if err != nil {
			fmt.Printf("dd: test index %d: %s\n", k, err)

			continue
		}

		if !isEqual(p, v.out) {
			fmt.Printf("dd: want %s, got %s\n", pp(v.out), pp(p))
		}
	}
}

func TestDJHerbis(t *testing.T) {
	for _, v := range testData() {
		p := djhpkg.Parse(v.in)

		if !isEqual(p, v.out) {
			fmt.Printf("djh: want %s, got %s\n", pp(v.out), pp(p))
		}
	}
}

func TestSRT(t *testing.T) {
	for _, v := range testData() {
		p := srtpkg.Parse(v.in)

		if !isEqual(p, v.out) {
			fmt.Printf("srt: want %s, got %s\n", pp(v.out), pp(p))
		}
	}
}

func TestJasonMoo(t *testing.T) {
	for _, v := range testData() {
		p := jmpkg.Parse(v.in)

		if !isEqual(p, v.out) {
			fmt.Printf("jm: want %s, got %s\n", pp(v.out), pp(p))
		}
	}
}

func isEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}

func pp(s []string) string {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "borked"
	}

	return string(b)
}
