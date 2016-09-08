package test

import (
	"Escargot/structures/caches"
	"testing"
)

func TestNewCache(t *testing.T) {
	caches.NewLRU()
}

func TestEmptyCache(t *testing.T) {
	lru := caches.NewLRU()
	if result := lru.Get("nothing"); result != nil {
		t.Errorf("Expecte nil on empty get, got %v", result)
	}
}

func TestPutGet(t *testing.T) {
	lru := caches.NewLRU()
	key := "five"
	value := 5
	lru.Put(key, value)
	if result := lru.Get(key); result != value {
		t.Errorf("Expected value of %v, got %v", value, result)
	}
}

func TestKeyCollision(t *testing.T) {
	lru := caches.NewLRU()
	key := "five"
	value := 5
	secondValue := 5
	lru.Put(key, value)
	lru.Put(key, secondValue)
	if result := lru.Get(key); result != secondValue {
		t.Errorf("Expected value of %v, got %v", secondValue, result)
	}
}
