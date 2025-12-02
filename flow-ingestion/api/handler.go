package main

import (
	"io"
	"log"
	"net/http"
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	SendMessage(string(bodyBytes), topic)
	w.WriteHeader(http.StatusOK)
}
