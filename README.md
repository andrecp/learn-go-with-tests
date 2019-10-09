# learn-go-with-tests

* Each folder represents a link in the tutorial

 ## Godoc

Install with `go get golang.org/x/tools/cmd/godoc`, run with `godoc -http :8000` and access by going to http://localhost:8000/pkg/testing/

## Example

Examples are run as part of the test suite and are properly formatted in `godoc`, for example

```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```

## Benchmark

`go test -bench=.` executes functions like

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```
