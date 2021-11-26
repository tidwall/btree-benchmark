# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `google`: The [google/btree](https://github.com/google/btree) package
- `go-arr`: A simple Go array

The following benchmarks were run on my 2019 Macbook Pro (2.4 GHz 8-Core Intel Core i9) 
using Go 1.17.3. The items are simple 8-byte ints. BTrees are created in non-concurrent mode.

```
** sequential set **
google:  set-seq        1,000,000 ops in 178ms, 5,618,049/sec, 177 ns/op, 39.0 MB, 40 bytes/op
tidwall: set-seq        1,000,000 ops in 156ms, 6,389,837/sec, 156 ns/op, 23.5 MB, 24 bytes/op
tidwall: set-seq-hint   1,000,000 ops in 78ms, 12,895,355/sec, 77 ns/op, 23.5 MB, 24 bytes/op
tidwall: load-seq       1,000,000 ops in 53ms, 18,937,400/sec, 52 ns/op, 23.5 MB, 24 bytes/op
go-arr:  append         1,000,000 ops in 78ms, 12,843,432/sec, 77 ns/op

** random set **
google:  set-rand       1,000,000 ops in 555ms, 1,803,133/sec, 554 ns/op, 29.7 MB, 31 bytes/op
tidwall: set-rand       1,000,000 ops in 545ms, 1,835,818/sec, 544 ns/op, 29.6 MB, 31 bytes/op
tidwall: set-rand-hint  1,000,000 ops in 670ms, 1,493,473/sec, 669 ns/op, 29.6 MB, 31 bytes/op
tidwall: set-again      1,000,000 ops in 681ms, 1,469,038/sec, 680 ns/op
tidwall: set-after-copy 1,000,000 ops in 670ms, 1,493,230/sec, 669 ns/op
tidwall: load-rand      1,000,000 ops in 569ms, 1,756,187/sec, 569 ns/op, 29.6 MB, 31 bytes/op

** sequential get **
google:  get-seq        1,000,000 ops in 165ms, 6,048,307/sec, 165 ns/op
tidwall: get-seq        1,000,000 ops in 144ms, 6,940,120/sec, 144 ns/op
tidwall: get-seq-hint   1,000,000 ops in 78ms, 12,815,243/sec, 78 ns/op

** random get **
google:  get-rand       1,000,000 ops in 701ms, 1,427,507/sec, 700 ns/op
tidwall: get-rand       1,000,000 ops in 679ms, 1,473,531/sec, 678 ns/op
tidwall: get-rand-hint  1,000,000 ops in 824ms, 1,213,805/sec, 823 ns/op
```
