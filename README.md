# Observable

This package solves the problem of broadcasting data from multiple publishers to multiple observers.
This package requires go 1.18.

- Create observable value with `observable.New[type](value)`
- Subscribe for changes with `obj.Subscribe(ctx, sendCurrent)`
- Broadcast new value with `obj.Update(newValue)`
# Example

```golang
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

obj := observable.New[string]("initial")
ch := obj.Subscribe(ctx, true)
var results []string

var wg sync.WaitGroup
wg.Add(4)
go func() {
    for r := range ch {
        results = append(results, r)
        wg.Done()
    }
}()

obj.Update("value")
obj.Update("is")
obj.Update("updated")

wg.Wait()
cancel()

fmt.Println(results) // [initial value is updated]
```
