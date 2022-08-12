# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `google`: The [google/btree](https://github.com/google/btree) package (without generics)
- `google(G)`: The [google/btree](https://github.com/google/btree) package (generics)
- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package (without generics)
- `tidwall(G)`: The [tidwall/btree](https://github.com/tidwall/btree) package (generics)
- `tidwall(M)`: The [tidwall/btree](https://github.com/tidwall/btree) package (generics using the `btree.Map` type)
- `go-arr`: A simple Go array

The following benchmarks were run on my 2021 Macbook Pro M1 Max 
using Go version 1.18.2.
The degrees is 128. The items are simple 8-byte ints. 

```
** sequential set **
google:     set-seq        1,000,000 ops in 144ms, 6,947,904/sec, 143 ns/op, 39.0 MB, 40 bytes/op
google(G):  set-seq        1,000,000 ops in 83ms, 12,109,255/sec, 82 ns/op, 16.0 MB, 16 bytes/op
tidwall:    set-seq        1,000,000 ops in 131ms, 7,619,018/sec, 131 ns/op, 38.9 MB, 40 bytes/op
tidwall(G): set-seq        1,000,000 ops in 82ms, 12,232,671/sec, 81 ns/op, 15.9 MB, 16 bytes/op
tidwall(M): set-seq        1,000,000 ops in 29ms, 34,080,355/sec, 29 ns/op, 15.9 MB, 16 bytes/op
tidwall:    load-seq       1,000,000 ops in 30ms, 33,647,800/sec, 29 ns/op, 38.9 MB, 40 bytes/op
tidwall(G): load-seq       1,000,000 ops in 16ms, 64,174,554/sec, 15 ns/op, 15.9 MB, 16 bytes/op
tidwall(M): load-seq       1,000,000 ops in 13ms, 74,798,981/sec, 13 ns/op, 15.9 MB, 16 bytes/op
go-arr:     append         1,000,000 ops in 6ms, 160,630,726/sec, 6 ns/op, 8.1 MB, 8 bytes/op

** sequential get **
google:     get-seq        1,000,000 ops in 135ms, 7,382,798/sec, 135 ns/op
google(G):  get-seq        1,000,000 ops in 80ms, 12,542,815/sec, 79 ns/op
tidwall:    get-seq        1,000,000 ops in 132ms, 7,565,131/sec, 132 ns/op
tidwall(G): get-seq        1,000,000 ops in 77ms, 12,920,939/sec, 77 ns/op
tidwall(M): get-seq        1,000,000 ops in 43ms, 23,115,196/sec, 43 ns/op

** random set **
google:     set-rand       1,000,000 ops in 458ms, 2,182,319/sec, 458 ns/op, 29.7 MB, 31 bytes/op
google(G):  set-rand       1,000,000 ops in 161ms, 6,196,735/sec, 161 ns/op, 11.2 MB, 11 bytes/op
tidwall:    set-rand       1,000,000 ops in 473ms, 2,115,070/sec, 472 ns/op, 29.7 MB, 31 bytes/op
tidwall(G): set-rand       1,000,000 ops in 158ms, 6,344,256/sec, 157 ns/op, 11.2 MB, 11 bytes/op
tidwall(M): set-rand       1,000,000 ops in 150ms, 6,687,662/sec, 149 ns/op, 11.2 MB, 11 bytes/op
tidwall:    set-after-copy 1,000,000 ops in 559ms, 1,787,460/sec, 559 ns/op, 88 bytes, 0 bytes/op
tidwall(G): set-after-copy 1,000,000 ops in 182ms, 5,493,477/sec, 182 ns/op, 88 bytes, 0 bytes/op
tidwall:    load-rand      1,000,000 ops in 483ms, 2,072,016/sec, 482 ns/op, 29.7 MB, 31 bytes/op
tidwall(G): load-rand      1,000,000 ops in 168ms, 5,969,528/sec, 167 ns/op, 11.2 MB, 11 bytes/op
tidwall(M): load-rand      1,000,000 ops in 162ms, 6,184,374/sec, 161 ns/op, 11.2 MB, 11 bytes/op

** random get **
google:     get-rand       1,000,000 ops in 790ms, 1,266,378/sec, 789 ns/op, 464 bytes, 0 bytes/op
google(G):  get-rand       1,000,000 ops in 162ms, 6,165,913/sec, 162 ns/op
tidwall:    get-rand       1,000,000 ops in 805ms, 1,242,231/sec, 805 ns/op
tidwall(G): get-rand       1,000,000 ops in 166ms, 6,032,594/sec, 165 ns/op
tidwall(M): get-rand       1,000,000 ops in 166ms, 6,007,088/sec, 166 ns/op

** range **
google:     ascend        1,000,000 ops in 11ms, 87,465,150/sec, 11 ns/op
google(G):  ascend        1,000,000 ops in 5ms, 193,182,287/sec, 5 ns/op
tidwall:    ascend        1,000,000 ops in 10ms, 101,164,658/sec, 9 ns/op
tidwall(G): iter          1,000,000 ops in 6ms, 171,719,277/sec, 5 ns/op
tidwall(G): scan          1,000,000 ops in 5ms, 212,826,363/sec, 4 ns/op
tidwall(G): walk          1,000,000 ops in 4ms, 231,249,536/sec, 4 ns/op
go-arr:     for-loop      1,000,000 ops in 2ms, 608,041,346/sec, 1 ns/op
```
