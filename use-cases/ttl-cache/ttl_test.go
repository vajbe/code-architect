package main

import (
	"testing"
	"time"
)

func TestTTLCache_KeyExistence(t *testing.T) {
	// Verify key existence
	ch := NewCacheStore()
	k := "Vivek"
	ch.Set(k, "eyjc1234543", 5*time.Second)

	_, ok := ch.Get(k)

	if !ok {
		t.Fatal("Key not found")
	}

	t.Log("Key exists ", k)

}

func TestTTLCache_KeyRemoval(t *testing.T) {
	// Verify key existence
	ch := NewCacheStore()
	k := "Vivek"
	ch.Set(k, "eyjc1234543", 1*time.Second)
	time.Sleep(1 * time.Second)
	_, ok := ch.Get(k)

	if !ok {
		t.Log("Key not found")
	} else {
		t.Fatal("Key exists ", k)
	}
}
