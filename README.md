# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `google`: The [google/btree](https://github.com/google/btree) package
- `go-arr`: A simple Go array

The following benchmarks were run on my 2019 Macbook Pro (2.4 GHz 8-Core Intel Core i9) 
using Go 1.16.5. The items are simple 8-byte ints. BTrees are created in non-concurrent mode.

```
** sequential set **
google:  set-seq        1,000,000 ops in 185ms, 5,417,843/sec, 184 ns/op, 30.9 MB, 32 bytes/op
tidwall: set-seq        1,000,000 ops in 127ms, 7,899,165/sec, 126 ns/op, 36.6 MB, 38 bytes/op
tidwall: set-seq-hint   1,000,000 ops in 62ms, 16,125,537/sec, 62 ns/op, 36.6 MB, 38 bytes/op
tidwall: load-seq       1,000,000 ops in 24ms, 42,112,100/sec, 23 ns/op, 36.6 MB, 38 bytes/op
go-arr:  append         1,000,000 ops in 67ms, 14,824,910/sec, 67 ns/op

** random set **
google:  set-rand       1,000,000 ops in 639ms, 1,564,308/sec, 639 ns/op, 21.5 MB, 22 bytes/op
tidwall: set-rand       1,000,000 ops in 544ms, 1,839,254/sec, 543 ns/op, 26.7 MB, 27 bytes/op
tidwall: set-rand-hint  1,000,000 ops in 611ms, 1,636,143/sec, 611 ns/op, 26.4 MB, 27 bytes/op
tidwall: set-again      1,000,000 ops in 399ms, 2,504,433/sec, 399 ns/op, 27.1 MB, 28 bytes/op
tidwall: set-after-copy 1,000,000 ops in 423ms, 2,362,738/sec, 423 ns/op, 27.9 MB, 29 bytes/op
tidwall: load-rand      1,000,000 ops in 551ms, 1,814,937/sec, 550 ns/op, 26.1 MB, 27 bytes/op

** sequential get **
google:  get-seq        1,000,000 ops in 121ms, 8,263,261/sec, 121 ns/op
tidwall: get-seq        1,000,000 ops in 96ms, 10,379,981/sec, 96 ns/op
tidwall: get-seq-hint   1,000,000 ops in 34ms, 29,231,769/sec, 34 ns/op

** random get **
google:  get-rand       1,000,000 ops in 152ms, 6,572,261/sec, 152 ns/op
tidwall: get-rand       1,000,000 ops in 126ms, 7,952,534/sec, 125 ns/op
tidwall: get-rand-hint  1,000,000 ops in 138ms, 7,255,665/sec, 137 ns/op
```
