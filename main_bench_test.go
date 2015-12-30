package main

import (
	"testing"

	"github.com/daved/pythatchallenge/internal/daveddev"
	"github.com/daved/pythatchallenge/internal/djherbis"
	"github.com/daved/pythatchallenge/internal/jasonmoo"
	"github.com/daved/pythatchallenge/internal/srtkkou"
)

func benchDataA() []string {
	return []string{"cp", `"file`, "name", "with", `spaces"`, `"file`, "name", `backup"`}
}
func benchDataB() []string {
	return []string{"one", "two", `"three`, "four", `five"`}
}

func BenchmarkA_DavedDev(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = ddpkg.Parse(benchDataA())
	}
}

func BenchmarkA_DJHerbis(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = djhpkg.Parse(benchDataA())
	}
}

func BenchmarkA_JasonMoo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = jmpkg.Parse(benchDataA())
	}
}

func BenchmarkA_SRT(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = srtpkg.Parse(benchDataA())
	}
}

func BenchmarkB_DavedDev(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = ddpkg.Parse(benchDataB())
	}
}

func BenchmarkB_DJHerbis(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = djhpkg.Parse(benchDataB())
	}
}

func BenchmarkB_JasonMoo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = jmpkg.Parse(benchDataB())
	}
}

func BenchmarkB_SRT(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = srtpkg.Parse(benchDataB())
	}
}
