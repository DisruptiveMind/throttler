package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/goware/throttler"
)

func wait(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is going to be legend... (wait for it)\n")
	if fl, ok := w.(http.Flusher); ok {
		fl.Flush()
	}
	time.Sleep(2 * time.Second)
	io.WriteString(w, "...dary! Legendary!\n")
}

func main() {
	waitHandler := http.HandlerFunc(wait)

	http.Handle("/", throttler.Limit(2)(waitHandler))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
