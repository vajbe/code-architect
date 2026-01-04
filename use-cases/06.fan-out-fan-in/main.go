package main

// write all http get put post delete in notes

/* HTTP Aggregator Service (Fan-out / Fan-in)
🔹 Design an HTTP service with endpoint /aggregate that:
Calls 3 downstream services concurrently
Aggregates responses

Respects a global timeout

Allows partial success

Cancels slow requests when timeout is hit */

import (
	"context"
	"log"
	"sync"
	"time"
)

type Service struct {
	urlId   string
	delay   time.Duration
	timeout time.Duration
}

type IData struct {
	Message string
	Id      int
	UrlId   string
}

func callApi(ctx context.Context, serviceURL string, delay time.Duration) (IData, error) {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return IData{
			Message: "Hello " + serviceURL,
			Id:      int(time.Now().Unix()),
			UrlId:   serviceURL,
		}, nil
	case <-ctx.Done():
		return IData{}, ctx.Err()
	}
}

func StartServices(ctx context.Context, services []Service, out chan<- IData) {
	wg := sync.WaitGroup{}

	for _, service := range services {
		wg.Add(1)
		go func(s Service) {
			defer wg.Done()

			svcCtx, cancel := context.WithTimeout(ctx, s.timeout)
			defer cancel()

			data, err := callApi(svcCtx, s.urlId, s.delay)
			if err != nil {
				log.Println("Error processing service; Id:", s.urlId)
			} else {
				out <- data
			}
		}(service)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func AggregateHandler(ctx context.Context, out <-chan IData) []IData {
	result := []IData{}
	for {
		select {
		case data, ok := <-out:
			if !ok {
				return result
			}
			result = append(result, data)
			log.Println("Received response from:", data.UrlId)
		case <-ctx.Done():
			log.Println("Wrapping up goroutine")
			return result
		}
	}
}

func main() {
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
	log.Println("Result after processing; ", result)
}
