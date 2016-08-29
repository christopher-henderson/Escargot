package main

import (
	"Escargot/structures/caches"
	"fmt"
)

func main() {
	lru := caches.NewLRU()
	lru.Put("bob", 5)
	lru.Put("bob1", 6)
	lru.Put("bob2", 7)
	lru.Put("bob3", 8)
	lru.Put("bob4", 9)
	lru.Put("bob4", "bilbo baggins")
	lru.Put("bob5", [...]int{1.0, 2.0, 3.0})
	fmt.Println(lru.Slice())
	lru.Get("bob")
	fmt.Println(lru.Slice())
}
