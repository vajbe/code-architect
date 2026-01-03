package main

import "testing"

func TestPut_Cache(t *testing.T) {

	cs := NewLRUCache(10)

	cs.Put("Vivek", "Vivek")

	val, ok := cs.Get("Vivek")
	if ok {
		t.Logf("Key 'Vivek' found with value: %s", val)
	} else {
		t.Error("Key 'Vivek' should have persisted but was not found")
	}

}

func TestEvict_Cache(t *testing.T) {
	cs := NewLRUCache(0)

	if cs == nil {
		t.Log("Cache was not created due to zero-size")
		return
	}

	cs.Put("Vivek", "Vivek")
	cs.Put("Cyberhaven", "Cyberhaven")

	val, ok := cs.Get("Vivek")
	if !ok {
		t.Log("Key 'Vivek' was correctly evicted")
	} else {
		t.Errorf("Key 'Vivek' should have been evicted but was found with value: %s", val)
	}

}
