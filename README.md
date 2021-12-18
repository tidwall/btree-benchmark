# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `google`: The [google/btree](https://github.com/google/btree) package (without generics)
- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package (without generics)
- `tidwall(G)`: The [tidwall/btree](https://github.com/tidwall/btree) package (with generics)
- `go-arr`: A simple Go array

The following benchmarks were run on my 2019 Macbook Pro (2.4 GHz 8-Core Intel Core i9) 
using Go Development version 1.18 (beta1).
The items are simple 8-byte ints. 

```
** sequential set **
google:     set-seq        1,000,000 ops in 151ms, 6,635,490/sec, 150 ns/op, 39.0 MB, 40 bytes/op
tidwall:    set-seq        1,000,000 ops in 137ms, 7,279,809/sec, 137 ns/op, 23.5 MB, 24 bytes/op
tidwall(G): set-seq        1,000,000 ops in 79ms, 12,621,967/sec, 79 ns/op, 8.2 MB, 8 bytes/op
tidwall:    set-seq-hint   1,000,000 ops in 71ms, 13,998,690/sec, 71 ns/op, 23.5 MB, 24 bytes/op
tidwall(G): set-seq-hint   1,000,000 ops in 46ms, 21,537,424/sec, 46 ns/op, 8.2 MB, 8 bytes/op
tidwall:    load-seq       1,000,000 ops in 43ms, 23,488,462/sec, 42 ns/op, 23.5 MB, 24 bytes/op
tidwall(G): load-seq       1,000,000 ops in 21ms, 48,613,969/sec, 20 ns/op, 8.2 MB, 8 bytes/op
go-arr:     append         1,000,000 ops in 23ms, 43,374,512/sec, 23 ns/op, 8.1 MB, 8 bytes/op

** random set **
google:     set-rand       1,000,000 ops in 525ms, 1,903,748/sec, 525 ns/op, 29.7 MB, 31 bytes/op
tidwall:    set-rand       1,000,000 ops in 519ms, 1,926,287/sec, 519 ns/op, 29.6 MB, 31 bytes/op
tidwall(G): set-rand       1,000,000 ops in 216ms, 4,630,150/sec, 215 ns/op, 11.2 MB, 11 bytes/op
tidwall:    set-rand-hint  1,000,000 ops in 562ms, 1,778,253/sec, 562 ns/op, 29.6 MB, 31 bytes/op
tidwall(G): set-rand-hint  1,000,000 ops in 252ms, 3,961,275/sec, 252 ns/op, 11.2 MB, 11 bytes/op
tidwall:    set-again      1,000,000 ops in 645ms, 1,550,044/sec, 645 ns/op
tidwall(G): set-again      1,000,000 ops in 208ms, 4,802,954/sec, 208 ns/op
tidwall(:   set-after-copy 1,000,000 ops in 650ms, 1,538,587/sec, 649 ns/op
tidwall(G): set-after-copy 1,000,000 ops in 218ms, 4,584,576/sec, 218 ns/op
tidwall:    load-rand      1,000,000 ops in 532ms, 1,878,363/sec, 532 ns/op, 29.6 MB, 31 bytes/op
tidwall(G): load-rand      1,000,000 ops in 224ms, 4,472,595/sec, 223 ns/op, 11.2 MB, 11 bytes/op

** sequential get **
google:     get-seq        1,000,000 ops in 161ms, 6,219,591/sec, 160 ns/op
tidwall:    get-seq        1,000,000 ops in 137ms, 7,325,738/sec, 136 ns/op
tidwall(G): get-seq        1,000,000 ops in 77ms, 13,034,899/sec, 76 ns/op
tidwall:    get-seq-hint   1,000,000 ops in 68ms, 14,794,017/sec, 67 ns/op
tidwall(G): get-seq-hint   1,000,000 ops in 34ms, 29,327,591/sec, 34 ns/op

** random get **
google:     get-rand       1,000,000 ops in 675ms, 1,481,742/sec, 674 ns/op
tidwall:    get-rand       1,000,000 ops in 650ms, 1,537,988/sec, 650 ns/op
tidwall(G): get-rand       1,000,000 ops in 204ms, 4,913,388/sec, 203 ns/op
tidwall:    get-rand-hint  1,000,000 ops in 741ms, 1,350,287/sec, 740 ns/op
tidwall(G): get-rand-hint  1,000,000 ops in 274ms, 3,643,459/sec, 274 ns/op

** range **
google:     ascend        1,000,000 ops in 17ms, 59,269,441/sec, 16 ns/op
tidwall:    ascend        1,000,000 ops in 12ms, 80,983,194/sec, 12 ns/op
tidwall(G): iter          1,000,000 ops in 7ms, 147,151,593/sec, 6 ns/op
tidwall(G): scan          1,000,000 ops in 5ms, 201,207,526/sec, 4 ns/op
tidwall(G): walk          1,000,000 ops in 4ms, 229,385,050/sec, 4 ns/op
go-arr:     for-loop      1,000,000 ops in 3ms, 348,544,512/sec, 2 ns/op
```
