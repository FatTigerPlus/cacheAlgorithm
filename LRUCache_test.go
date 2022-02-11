package cache

import "testing"

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(5)
	cache.Set(10, 20)
	cache.Set(20, 20)
	cache.Set(30, 20)
	cache.Set(40, 20)
	cache.Set(50, 20)
	cache.Set(60, 20)
	val, err := cache.Get(10)
	if err != nil {
		t.Errorf("get val error is: %s", err)
	}
	if val != nil {
		t.Errorf("cache miss err is: %s", err)
	}
	v, _ := cache.Get(60)
	println(v.(int))
}
