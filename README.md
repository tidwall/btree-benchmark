# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `google`: The [google/btree](https://github.com/google/btree) package
- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `tidwall(G)`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `go-arr`: A simple Go array

The following benchmarks were run on my 2019 Macbook Pro (2.4 GHz 8-Core Intel Core i9) 
using Go Development version 1.18-8ff254e3 (gotip).
The items are simple 8-byte ints. 

```
** sequential set **
google:     set-seq        1,000,000 ops in 157ms, 6,352,361/sec, 157 ns/op, 39.0 MB, 40 bytes/op
tidwall:    set-seq        1,000,000 ops in 144ms, 6,934,619/sec, 144 ns/op, 23.5 MB, 24 bytes/op
tidwall(G): set-seq        1,000,000 ops in 82ms, 12,205,826/sec, 81 ns/op, 8.2 MB, 8 bytes/op
tidwall:    set-seq-hint   1,000,000 ops in 75ms, 13,256,139/sec, 75 ns/op, 23.5 MB, 24 bytes/op
tidwall(G): set-seq-hint   1,000,000 ops in 47ms, 21,073,524/sec, 47 ns/op, 8.2 MB, 8 bytes/op
tidwall:    load-seq       1,000,000 ops in 47ms, 21,438,912/sec, 46 ns/op, 23.5 MB, 24 bytes/op
tidwall(G): load-seq       1,000,000 ops in 23ms, 43,405,780/sec, 23 ns/op, 8.2 MB, 8 bytes/op
go-arr:     append         1,000,000 ops in 71ms, 14,104,431/sec, 70 ns/op

** random set **
google:     set-rand       1,000,000 ops in 593ms, 1,685,165/sec, 593 ns/op, 29.7 MB, 31 bytes/op
tidwall:    set-rand       1,000,000 ops in 571ms, 1,751,559/sec, 570 ns/op, 29.6 MB, 31 bytes/op
tidwall(G): set-rand       1,000,000 ops in 232ms, 4,307,484/sec, 232 ns/op, 11.2 MB, 11 bytes/op
tidwall:    set-rand-hint  1,000,000 ops in 632ms, 1,582,840/sec, 631 ns/op, 29.6 MB, 31 bytes/op
tidwall(G): set-rand-hint  1,000,000 ops in 267ms, 3,752,093/sec, 266 ns/op, 11.2 MB, 11 bytes/op
tidwall:    set-again      1,000,000 ops in 682ms, 1,466,345/sec, 681 ns/op
tidwall(G): set-again      1,000,000 ops in 244ms, 4,092,881/sec, 244 ns/op
tidwall(:   set-after-copy 1,000,000 ops in 679ms, 1,472,038/sec, 679 ns/op
tidwall(G): set-after-copy 1,000,000 ops in 242ms, 4,139,555/sec, 241 ns/op
tidwall:    load-rand      1,000,000 ops in 591ms, 1,691,903/sec, 591 ns/op, 29.6 MB, 31 bytes/op
tidwall(G): load-rand      1,000,000 ops in 257ms, 3,893,858/sec, 256 ns/op, 11.2 MB, 11 bytes/op

** sequential get **
google:     get-seq        1,000,000 ops in 173ms, 5,772,752/sec, 173 ns/op
tidwall:    get-seq        1,000,000 ops in 164ms, 6,099,748/sec, 163 ns/op
tidwall(G): get-seq        1,000,000 ops in 82ms, 12,212,668/sec, 81 ns/op
tidwall:    get-seq-hint   1,000,000 ops in 84ms, 11,959,334/sec, 83 ns/op
tidwall(G): get-seq-hint   1,000,000 ops in 36ms, 27,856,690/sec, 35 ns/op

** random get **
google:     get-rand       1,000,000 ops in 708ms, 1,412,173/sec, 708 ns/op
tidwall:    get-rand       1,000,000 ops in 694ms, 1,441,070/sec, 693 ns/op
tidwall(G): get-rand       1,000,000 ops in 246ms, 4,064,635/sec, 246 ns/op
tidwall:    get-rand-hint  1,000,000 ops in 775ms, 1,290,195/sec, 775 ns/op
tidwall(G): get-rand-hint  1,000,000 ops in 280ms, 3,575,213/sec, 279 ns/op
```
