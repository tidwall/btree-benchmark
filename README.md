# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `google`: The [google/btree](https://github.com/google/btree) package
- `go-arr`: A simple Go array

The following benchmarks were run on my 2019 Macbook Pro (2.4 GHz 8-Core Intel Core i9) 
using Go 1.17. The items are simple 8-byte ints. BTrees are created in non-concurrent mode.

```
** sequential set **
google:  set-seq        1,000,000 ops in 129ms, 7,761,884/sec, 128 ns/op, 31.0 MB, 32 bytes/op
tidwall: set-seq        1,000,000 ops in 116ms, 8,655,931/sec, 115 ns/op, 36.6 MB, 38 bytes/op
tidwall: set-seq-hint   1,000,000 ops in 52ms, 19,219,654/sec, 52 ns/op, 36.6 MB, 38 bytes/op
tidwall: load-seq       1,000,000 ops in 22ms, 45,096,800/sec, 22 ns/op, 36.6 MB, 38 bytes/op
go-arr:  append         1,000,000 ops in 48ms, 20,860,238/sec, 47 ns/op

** random set **
google:  set-rand       1,000,000 ops in 533ms, 1,876,341/sec, 532 ns/op, 21.5 MB, 22 bytes/op
tidwall: set-rand       1,000,000 ops in 495ms, 2,020,118/sec, 495 ns/op, 26.7 MB, 27 bytes/op
tidwall: set-rand-hint  1,000,000 ops in 537ms, 1,863,372/sec, 536 ns/op, 26.4 MB, 27 bytes/op
tidwall: set-again      1,000,000 ops in 350ms, 2,857,997/sec, 349 ns/op, 27.1 MB, 28 bytes/op
tidwall: set-after-copy 1,000,000 ops in 373ms, 2,682,891/sec, 372 ns/op, 27.9 MB, 29 bytes/op
tidwall: load-rand      1,000,000 ops in 504ms, 1,984,558/sec, 503 ns/op, 26.1 MB, 27 bytes/op

** sequential get **
google:  get-seq        1,000,000 ops in 92ms, 10,851,246/sec, 92 ns/op
tidwall: get-seq        1,000,000 ops in 82ms, 12,224,334/sec, 81 ns/op
tidwall: get-seq-hint   1,000,000 ops in 29ms, 34,086,961/sec, 29 ns/op

** random get **
google:  get-rand       1,000,000 ops in 106ms, 9,426,080/sec, 106 ns/op
tidwall: get-rand       1,000,000 ops in 104ms, 9,641,568/sec, 103 ns/op
tidwall: get-rand-hint  1,000,000 ops in 113ms, 8,819,336/sec, 113 ns/op
```
