package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mx sync.Mutex
	m  map[string]int
}

// конструктор для мапа
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mx.Lock()
	defer s.mx.Unlock()
	val, ok := s.m[key]
	return val, ok
}

func (s *SafeMap) Set(key string, value int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.m[key] = value

}
func main() {
	m := NewSafeMap()
	var wg sync.WaitGroup
	numGoroutines := 20
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(ind int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", ind)
			m.Set(key, ind)
		}(i)
	}
	wg.Wait()

	wg.Add(numGoroutines * 2)
	for i := 0; i < numGoroutines; i++ {
		go func(ind int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", ind)
			if val, ok := m.Get(key); ok {
				fmt.Printf("Read: %s -> %d\n", key, val)
			}
		}(i)

		go func(ind int) {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", ind+1000)
			m.Set(key, ind)
		}(i)
	}

	wg.Wait()
}
