# Throttler
[Golang](http://golang.org/) [HTTP](http://golang.org/pkg/net/http/) Middleware that limits number of currently processed requests at the same time.

[![GoDoc](https://godoc.org/github.com/goware/throttler?status.png)](https://godoc.org/github.com/goware/throttler)
[![Travis](https://travis-ci.org/goware/throttler.svg?branch=master)](https://travis-ci.org/goware/throttler)

## Usage

### DefaultServeMux

See [example/mux/](./example/mux/).

```go
waitHandler := http.HandlerFunc(wait)
http.Handle("/", throttler.Limit(1)(waitHandler))

log.Fatal(http.ListenAndServe(":8000", nil))
```

### Goji

See [example/goji/](./example/goji/).

```go
goji.Use(throttler.Limit(2))
goji.Get("/", wait)

goji.Serve()
```

## License
Throttler is licensed under the [MIT License](./LICENSE).
