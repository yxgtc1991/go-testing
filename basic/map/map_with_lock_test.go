package main

import "testing"

// BenchmarkAddMapWithUnlock 无锁
func BenchmarkAddMapWithUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addMapWithUnlock()
	}
}

func addMapWithUnlock() {
	m := make(map[int]int)
	for i := 0; i < 100000; i++ {
		m[i] = i
	}
}
