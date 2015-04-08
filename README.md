# Throttler
[Golang](http://golang.org/) [HTTP](http://golang.org/pkg/net/http/) middleware to throttle the number of requests processed at a time.

[![GoDoc](https://godoc.org/github.com/goware/throttler?status.png)](https://godoc.org/github.com/goware/throttler)
[![Travis](https://travis-ci.org/goware/throttler.svg?branch=master)](https://travis-ci.org/goware/throttler)

## Usage

1. [Goji](#goji)
2. [DefaultServeMux (net/http)](#defaultservemux-nethttp)
3. ...don't see your favorite router/framework? We accept [Pull Requests](https://github.com/goware/throttler/pulls)!

### [Goji](https://github.com/zenazn/goji)

```go
// Limit to 5 requests globally.
goji.Use(throttler.Limit(5))

// Limit /admin route to 2 requests.
admin := web.New()
admin.Use(throttler.Limit(2))
admin.Get("/*", handler)
```

See [full example](./example/goji/main.go).

### [DefaultServeMux (net/http)](http://golang.org/pkg/net/http/#ServeMux)

```go
// Limit /admin route to 2 requests.
limit := throttler.Limit(2)
handlerFunc := http.HandlerFunc(handler)
http.Handle("/admin", limit(handlerFunc))

```

See [full example](./example/mux/main.go).

## License
Throttler is licensed under the [MIT License](./LICENSE).
