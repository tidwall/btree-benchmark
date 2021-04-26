# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `google`: The [google/btree](https://github.com/google/btree) package
- `go-arr`: A simple Go array

The following benchmarks were run on my 2019 Macbook Pro (2.4 GHz 8-Core Intel Core i9) 
using Go 1.16.2. The items are simple 8-byte ints.

```
** sequential set **
google:  set-seq        1,000,000 ops in 158ms, 6,321,982/sec, 158 ns/op, 30.9 MB, 32 bytes/op
tidwall: set-seq        1,000,000 ops in 142ms, 7,066,602/sec, 141 ns/op, 36.6 MB, 38 bytes/op
tidwall: set-seq-hint   1,000,000 ops in 77ms, 13,032,533/sec, 76 ns/op, 36.6 MB, 38 bytes/op
tidwall: load-seq       1,000,000 ops in 38ms, 26,151,358/sec, 38 ns/op, 36.6 MB, 38 bytes/op
go-arr:  append         1,000,000 ops in 49ms, 20,479,942/sec, 48 ns/op

** random set **
google:  set-rand       1,000,000 ops in 629ms, 1,589,978/sec, 628 ns/op, 21.5 MB, 22 bytes/op
tidwall: set-rand       1,000,000 ops in 569ms, 1,757,597/sec, 568 ns/op, 26.7 MB, 27 bytes/op
tidwall: set-rand-hint  1,000,000 ops in 565ms, 1,769,211/sec, 565 ns/op, 26.4 MB, 27 bytes/op
tidwall: set-again      1,000,000 ops in 390ms, 2,561,129/sec, 390 ns/op, 27.1 MB, 28 bytes/op
tidwall: set-after-copy 1,000,000 ops in 401ms, 2,496,375/sec, 400 ns/op, 27.9 MB, 29 bytes/op
tidwall: load-rand      1,000,000 ops in 520ms, 1,922,016/sec, 520 ns/op, 26.1 MB, 27 bytes/op

** sequential get **
google:  get-seq        1,000,000 ops in 119ms, 8,391,178/sec, 119 ns/op
tidwall: get-seq        1,000,000 ops in 118ms, 8,442,502/sec, 118 ns/op
tidwall: get-seq-hint   1,000,000 ops in 39ms, 25,949,759/sec, 38 ns/op

** random get **
google:  get-rand       1,000,000 ops in 151ms, 6,641,486/sec, 150 ns/op
tidwall: get-rand       1,000,000 ops in 125ms, 7,984,195/sec, 125 ns/op
tidwall: get-rand-hint  1,000,000 ops in 142ms, 7,022,156/sec, 142 ns/op
```
