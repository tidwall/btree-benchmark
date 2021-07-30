# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `google`: The [google/btree](https://github.com/google/btree) package
- `go-arr`: A simple Go array

The following benchmarks were run on my 2019 Macbook Pro (2.4 GHz 8-Core Intel Core i9) 
using Go 1.16.5. The items are simple 8-byte ints.

```
** sequential set **
google:  set-seq        1,000,000 ops in 163ms, 6,140,597/sec, 162 ns/op, 30.9 MB, 32 bytes/op
tidwall: set-seq        1,000,000 ops in 141ms, 7,075,240/sec, 141 ns/op, 36.6 MB, 38 bytes/op
tidwall: set-seq-hint   1,000,000 ops in 79ms, 12,673,902/sec, 78 ns/op, 36.6 MB, 38 bytes/op
tidwall: load-seq       1,000,000 ops in 40ms, 24,887,293/sec, 40 ns/op, 36.6 MB, 38 bytes/op
go-arr:  append         1,000,000 ops in 51ms, 19,617,269/sec, 50 ns/op

** random set **
google:  set-rand       1,000,000 ops in 666ms, 1,501,583/sec, 665 ns/op, 21.5 MB, 22 bytes/op
tidwall: set-rand       1,000,000 ops in 569ms, 1,756,845/sec, 569 ns/op, 26.7 MB, 27 bytes/op
tidwall: set-rand-hint  1,000,000 ops in 670ms, 1,491,637/sec, 670 ns/op, 26.4 MB, 27 bytes/op
tidwall: set-again      1,000,000 ops in 488ms, 2,050,667/sec, 487 ns/op, 27.1 MB, 28 bytes/op
tidwall: set-after-copy 1,000,000 ops in 494ms, 2,022,980/sec, 494 ns/op, 27.9 MB, 29 bytes/op
tidwall: load-rand      1,000,000 ops in 594ms, 1,682,937/sec, 594 ns/op, 26.1 MB, 27 bytes/op

** sequential get **
google:  get-seq        1,000,000 ops in 141ms, 7,078,690/sec, 141 ns/op
tidwall: get-seq        1,000,000 ops in 124ms, 8,075,925/sec, 123 ns/op
tidwall: get-seq-hint   1,000,000 ops in 40ms, 25,142,979/sec, 39 ns/op

** random get **
google:  get-rand       1,000,000 ops in 152ms, 6,593,518/sec, 151 ns/op
tidwall: get-rand       1,000,000 ops in 128ms, 7,783,293/sec, 128 ns/op
tidwall: get-rand-hint  1,000,000 ops in 135ms, 7,403,823/sec, 135 ns/op
```
