package problem146

import "testing"

func assertVal(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(3)
	assertVal(t, 0, cache.Len())
	assertVal(t, -1, cache.Get(1))

	cache.Put(1, 10)
	assertVal(t, 1, cache.Len())
	assertVal(t, 10, cache.Get(1))

	cache.Put(2, 20)
	cache.Put(3, 30)
	assertVal(t, 3, cache.Len())
	assertVal(t, 10, cache.Get(1))
	assertVal(t, 20, cache.Get(2))
	assertVal(t, 30, cache.Get(3))

	cache.Put(3, 40)
	assertVal(t, 40, cache.Get(3))
	cache.Put(3, 30)
	assertVal(t, 30, cache.Get(3))
	assertVal(t, 3, cache.Len())

	cache.Put(4, 40)
	assertVal(t, 3, cache.Len())
	assertVal(t, -1, cache.Get(1))
	assertVal(t, 20, cache.Get(2))
	assertVal(t, 30, cache.Get(3))
	assertVal(t, 40, cache.Get(4))
}
