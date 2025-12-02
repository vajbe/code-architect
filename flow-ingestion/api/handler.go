package main

import (
	"fmt"
	"net/http"
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte{}
	n, _ := r.Body.Read(data)
	fmt.Println(string(data[:n]))
	w.WriteHeader(http.StatusAccepted)
}
