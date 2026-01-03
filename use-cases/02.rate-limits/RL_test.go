package main

import (
	"testing"
)

func TestRateLimiter_AllowedLimit(t *testing.T) {

	userId := "test_user"
	store := NewCacheStore()

	for i := 0; i < MAX_LIMIT; i++ {
		isAllowed := isRequestAllowed(userId, store)
		if !isAllowed {
			t.Fatal("Request should be allowed")
		}
	}

	t.Log("Request should be allowed")

}

func TestRateLimiter_DisallowedLimit(t *testing.T) {
	userId := "test_user"
	store := NewCacheStore()

	for i := 0; i < 6; i++ {
		isAllowed := isRequestAllowed(userId, store)
		if !isAllowed {
			t.Log("Request should not be allowed ")
		}
	}
}
