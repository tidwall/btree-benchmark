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
google:     set-seq        1,000,000 ops in 164ms, 6,098,705/sec, 163 ns/op, 39.0 MB, 40 bytes/op
tidwall:    set-seq        1,000,000 ops in 155ms, 6,469,268/sec, 154 ns/op, 23.5 MB, 24 bytes/op
tidwall(G): set-seq        1,000,000 ops in 79ms, 12,647,065/sec, 79 ns/op, 8.2 MB, 8 bytes/op
tidwall:    set-seq-hint   1,000,000 ops in 71ms, 14,074,636/sec, 71 ns/op, 23.5 MB, 24 bytes/op
tidwall(G): set-seq-hint   1,000,000 ops in 46ms, 21,646,961/sec, 46 ns/op, 8.2 MB, 8 bytes/op
tidwall:    load-seq       1,000,000 ops in 41ms, 24,112,048/sec, 41 ns/op, 23.5 MB, 24 bytes/op
tidwall(G): load-seq       1,000,000 ops in 22ms, 44,875,579/sec, 22 ns/op, 8.2 MB, 8 bytes/op
go-arr:     append         1,000,000 ops in 22ms, 44,709,206/sec, 22 ns/op, 8.1 MB, 8 bytes/op

** sequential get **
google:     get-seq        1,000,000 ops in 124ms, 8,093,243/sec, 123 ns/op
tidwall:    get-seq        1,000,000 ops in 109ms, 9,186,378/sec, 108 ns/op
tidwall(G): get-seq        1,000,000 ops in 80ms, 12,461,917/sec, 80 ns/op
tidwall:    get-seq-hint   1,000,000 ops in 53ms, 18,989,592/sec, 52 ns/op
tidwall(G): get-seq-hint   1,000,000 ops in 38ms, 26,405,028/sec, 37 ns/op

** random set **
google:     set-rand       1,000,000 ops in 564ms, 1,772,822/sec, 564 ns/op, 29.7 MB, 31 bytes/op
tidwall:    set-rand       1,000,000 ops in 614ms, 1,629,402/sec, 613 ns/op, 29.6 MB, 31 bytes/op
tidwall(G): set-rand       1,000,000 ops in 219ms, 4,566,151/sec, 219 ns/op, 11.2 MB, 11 bytes/op
tidwall:    set-rand-hint  1,000,000 ops in 671ms, 1,489,527/sec, 671 ns/op, 29.6 MB, 31 bytes/op
tidwall(G): set-rand-hint  1,000,000 ops in 317ms, 3,158,224/sec, 316 ns/op, 11.2 MB, 11 bytes/op
tidwall:    set-again      1,000,000 ops in 669ms, 1,494,282/sec, 669 ns/op
tidwall(G): set-again      1,000,000 ops in 215ms, 4,647,206/sec, 215 ns/op
tidwall(:   set-after-copy 1,000,000 ops in 700ms, 1,428,075/sec, 700 ns/op
tidwall(G): set-after-copy 1,000,000 ops in 253ms, 3,951,027/sec, 253 ns/op
tidwall:    load-rand      1,000,000 ops in 579ms, 1,726,277/sec, 579 ns/op, 29.6 MB, 31 bytes/op
tidwall(G): load-rand      1,000,000 ops in 240ms, 4,169,993/sec, 239 ns/op, 11.2 MB, 11 bytes/op

** random get **
google:     get-rand       1,000,000 ops in 696ms, 1,435,926/sec, 696 ns/op
tidwall:    get-rand       1,000,000 ops in 661ms, 1,512,866/sec, 660 ns/op
tidwall(G): get-rand       1,000,000 ops in 218ms, 4,587,040/sec, 218 ns/op
tidwall:    get-rand-hint  1,000,000 ops in 744ms, 1,343,339/sec, 744 ns/op
tidwall(G): get-rand-hint  1,000,000 ops in 281ms, 3,557,355/sec, 281 ns/op

** range **
google:     ascend        1,000,000 ops in 22ms, 46,375,987/sec, 21 ns/op
tidwall:    ascend        1,000,000 ops in 13ms, 79,519,239/sec, 12 ns/op
tidwall(G): iter          1,000,000 ops in 7ms, 142,601,763/sec, 7 ns/op
tidwall(G): scan          1,000,000 ops in 5ms, 197,217,264/sec, 5 ns/op
tidwall(G): walk          1,000,000 ops in 4ms, 228,085,749/sec, 4 ns/op
go-arr:     for-loop      1,000,000 ops in 3ms, 399,110,940/sec, 2 ns/op
```
