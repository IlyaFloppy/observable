# Observable

[![CI](https://github.com/Antonboom/nilnil/actions/workflows/ci.yml/badge.svg)](https://github.com/IlyaFloppy/observable/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/IlyaFloppy/observable)](https://goreportcard.com/report/github.com/IlyaFloppy/observable)
[![Coverage](https://coveralls.io/repos/github/IlyaFloppy/observable/badge.svg?branch=master)](https://coveralls.io/github/IlyaFloppy/observable?branch=master)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

This package solves the problem of broadcasting data from multiple publishers to multiple observers.
This package requires go 1.18.

- `go get github.com/IlyaFloppy/observable`
- Create observable object with `obj := observable.New[type](value)`
- Subscribe for changes with `ch := obj.Subscribe(ctx, options...)`
- Broadcast new value with `obj.Set(newValue)`
# Example

```golang
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

obj := observable.New[string]("initial")
ch := obj.Subscribe(ctx, observable.WithCurrent(true))
var results []string

var wg sync.WaitGroup
wg.Add(4)
go func() {
    for r := range ch {
        results = append(results, r)
        wg.Done()
    }
}()

obj.Set("value")
obj.Set("is")
obj.Set("updated")

wg.Wait()
cancel()

fmt.Println(results) // [initial value is updated]
```
