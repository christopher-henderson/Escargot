package caches

// Cache is the generic interface for an Escargot cache.
// at minimum, regardless of the eviction stratgey, a cache
// must honor retrieval, emplacement, and eviction.
type Cache interface {
	Get(key interface{}) (interface{}, error)
	Put(key interface{}) error
	evict() error
}
