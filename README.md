# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `google`: The [google/btree](https://github.com/google/btree) package
- `go-arr`: A simple Go array

The following benchmarks were run on my 2019 Macbook Pro (2.4 GHz 8-Core Intel Core i9) 
using Go 1.17. The items are simple 8-byte ints. BTrees are created in non-concurrent mode.

```
** sequential set **
google:  set-seq        1,000,000 ops in 162ms, 6,183,760/sec, 161 ns/op, 38.6 MB, 40 bytes/op
tidwall: set-seq        1,000,000 ops in 148ms, 6,755,156/sec, 148 ns/op, 44.3 MB, 46 bytes/op
tidwall: set-seq-hint   1,000,000 ops in 66ms, 15,100,485/sec, 66 ns/op, 44.3 MB, 46 bytes/op
tidwall: load-seq       1,000,000 ops in 39ms, 25,510,441/sec, 39 ns/op, 44.3 MB, 46 bytes/op
go-arr:  append         1,000,000 ops in 69ms, 14,454,379/sec, 69 ns/op, 424 bytes, 0 bytes/op

** random set **
google:  set-rand       1,000,000 ops in 653ms, 1,530,252/sec, 653 ns/op, 29.1 MB, 30 bytes/op
tidwall: set-rand       1,000,000 ops in 610ms, 1,638,596/sec, 610 ns/op, 34.3 MB, 35 bytes/op
tidwall: set-rand-hint  1,000,000 ops in 615ms, 1,624,795/sec, 615 ns/op, 34.0 MB, 35 bytes/op
tidwall: set-again      1,000,000 ops in 786ms, 1,272,125/sec, 786 ns/op
tidwall: set-after-copy 1,000,000 ops in 791ms, 1,263,949/sec, 791 ns/op
tidwall: load-rand      1,000,000 ops in 542ms, 1,844,838/sec, 542 ns/op, 33.7 MB, 35 bytes/op

** sequential get **
google:  get-seq        1,000,000 ops in 161ms, 6,216,649/sec, 160 ns/op
tidwall: get-seq        1,000,000 ops in 132ms, 7,548,647/sec, 132 ns/op
tidwall: get-seq-hint   1,000,000 ops in 64ms, 15,716,405/sec, 63 ns/op

** random get **
google:  get-rand       1,000,000 ops in 662ms, 1,511,374/sec, 661 ns/op
tidwall: get-rand       1,000,000 ops in 702ms, 1,424,161/sec, 702 ns/op
tidwall: get-rand-hint  1,000,000 ops in 792ms, 1,263,192/sec, 791 ns/op
```
