package main

import (
	"context"
	"testing"
	"time"
)

func Test_FanProcessResult(t *testing.T) {
	services := []Service{}
	services = append(services,
		Service{urlId: "url1", delay: 2 * time.Second, timeout: 3 * time.Second},
		Service{urlId: "url2", delay: 2 * time.Second, timeout: 3 * time.Second},
		Service{urlId: "url3", delay: 2 * time.Second, timeout: 3 * time.Second},
	)
	out := make(chan IData, len(services))
	ctx, cancel := context.WithTimeout(context.Background(), 13*time.Second)
	defer cancel()
	StartServices(ctx, services, out)
	result := AggregateHandler(ctx, out)
	if len(result) == 2 {
		t.Log("Partial processing worked correctly")
	} else {
		t.Fatal("Partial processing did not work correctly")
	}
}
