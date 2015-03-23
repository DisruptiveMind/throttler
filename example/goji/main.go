package main

import (
	"io"
	"net/http"
	"time"

	"github.com/goware/throttler"

	"github.com/zenazn/goji"
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
	goji.Use(throttler.Limit(2))
	goji.Get("/", wait)
	goji.Serve()
}
