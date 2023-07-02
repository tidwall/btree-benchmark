# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `google`: The [google/btree](https://github.com/google/btree) package (without generics)
- `google(G)`: The [google/btree](https://github.com/google/btree) package (generics)
- `pile(M)`: The [pascaldekloe/pile](https://github.com/pascaldekloe/pile) package (generics)
- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package (without generics)
- `tidwall(G)`: The [tidwall/btree](https://github.com/tidwall/btree) package (generics)
- `tidwall(M)`: The [tidwall/btree](https://github.com/tidwall/btree) package (generics using the `btree.Map` type)
- `go-arr`: A simple Go array

The following benchmarks were run on my 2021 Macbook Pro M1 Max 
using Go version 1.19.3.
All items are key/value pairs where the key is a string filled with 16 random digits such as `5204379379828236`, and the value is the int64 representation of the key.
The degree is 32.  

```
degree=32, key=string (16 bytes), val=int64, count=1000000

** sequential set **
google:     set-seq        1,000,000 ops in 223ms, 4,476,918/sec, 223 ns/op, 56.9 MB, 59.6 bytes/op
google(G):  set-seq        1,000,000 ops in 166ms, 6,034,933/sec, 165 ns/op, 49.7 MB, 52.1 bytes/op
tidwall:    set-seq        1,000,000 ops in 148ms, 6,755,073/sec, 148 ns/op, 56.4 MB, 59.1 bytes/op
tidwall(G): set-seq        1,000,000 ops in 116ms, 8,655,897/sec, 115 ns/op, 49.2 MB, 51.6 bytes/op
tidwall(M): set-seq        1,000,000 ops in 88ms, 11,342,680/sec, 88 ns/op, 49.2 MB, 51.6 bytes/op
tidwall:    set-seq-hint   1,000,000 ops in 90ms, 11,151,385/sec, 89 ns/op, 56.4 MB, 59.1 bytes/op
tidwall(G): set-seq-hint   1,000,000 ops in 60ms, 16,783,122/sec, 59 ns/op, 49.2 MB, 51.6 bytes/op
tidwall:    load-seq       1,000,000 ops in 50ms, 20,011,289/sec, 49 ns/op, 56.4 MB, 59.1 bytes/op
tidwall(G): load-seq       1,000,000 ops in 31ms, 31,979,831/sec, 31 ns/op, 49.2 MB, 51.6 bytes/op
tidwall(M): load-seq       1,000,000 ops in 27ms, 37,222,055/sec, 26 ns/op, 49.2 MB, 51.6 bytes/op
go-arr:     append         1,000,000 ops in 19ms, 52,924,636/sec, 18 ns/op, 26.5 MB, 27.7 bytes/op

** sequential get **
google:     get-seq        1,000,000 ops in 199ms, 5,025,759/sec, 198 ns/op
google(G):  get-seq        1,000,000 ops in 160ms, 6,255,822/sec, 159 ns/op
tidwall:    get-seq        1,000,000 ops in 155ms, 6,446,902/sec, 155 ns/op
tidwall(G): get-seq        1,000,000 ops in 110ms, 9,069,557/sec, 110 ns/op
tidwall(M): get-seq        1,000,000 ops in 93ms, 10,722,998/sec, 93 ns/op
tidwall:    get-seq-hint   1,000,000 ops in 81ms, 12,357,674/sec, 80 ns/op
tidwall(G): get-seq-hint   1,000,000 ops in 49ms, 20,362,037/sec, 49 ns/op

** random set **
google:     set-rand       1,000,000 ops in 1077ms, 928,915/sec, 1076 ns/op, 46.4 MB, 48.7 bytes/op
google(G):  set-rand       1,000,000 ops in 735ms, 1,359,985/sec, 735 ns/op, 34.5 MB, 36.2 bytes/op
tidwall:    set-rand       1,000,000 ops in 830ms, 1,204,150/sec, 830 ns/op, 46.4 MB, 48.6 bytes/op
tidwall(G): set-rand       1,000,000 ops in 558ms, 1,792,369/sec, 557 ns/op, 34.6 MB, 36.3 bytes/op
tidwall(M): set-rand       1,000,000 ops in 529ms, 1,891,806/sec, 528 ns/op, 34.6 MB, 36.3 bytes/op
tidwall:    set-rand-hint  1,000,000 ops in 870ms, 1,148,923/sec, 870 ns/op, 46.4 MB, 48.6 bytes/op
tidwall(G): set-rand-hint  1,000,000 ops in 614ms, 1,629,389/sec, 613 ns/op, 34.6 MB, 36.3 bytes/op
tidwall:    set-after-copy 1,000,000 ops in 972ms, 1,028,466/sec, 972 ns/op
tidwall(G): set-after-copy 1,000,000 ops in 606ms, 1,650,820/sec, 605 ns/op
tidwall:    load-rand      1,000,000 ops in 860ms, 1,162,836/sec, 859 ns/op, 46.4 MB, 48.6 bytes/op
tidwall(G): load-rand      1,000,000 ops in 583ms, 1,714,875/sec, 583 ns/op, 34.6 MB, 36.3 bytes/op
tidwall(M): load-rand      1,000,000 ops in 535ms, 1,868,764/sec, 535 ns/op, 34.6 MB, 36.3 bytes/op

** random get **
google:     get-rand       1,000,000 ops in 1508ms, 663,149/sec, 1507 ns/op, 32 bytes, 0.0 bytes/op
google(G):  get-rand       1,000,000 ops in 835ms, 1,197,256/sec, 835 ns/op
tidwall:    get-rand       1,000,000 ops in 1087ms, 920,025/sec, 1086 ns/op
tidwall(G): get-rand       1,000,000 ops in 643ms, 1,554,417/sec, 643 ns/op
tidwall(M): get-rand       1,000,000 ops in 608ms, 1,644,485/sec, 608 ns/op
tidwall:    get-rand-hint  1,000,000 ops in 1126ms, 888,007/sec, 1126 ns/op
tidwall(G): get-rand-hint  1,000,000 ops in 709ms, 1,410,174/sec, 709 ns/op

** range **
google:     ascend        1,000,000 ops in 9ms, 113,213,421/sec, 8 ns/op
google(G):  ascend        1,000,000 ops in 11ms, 93,153,965/sec, 10 ns/op
tidwall:    ascend        1,000,000 ops in 7ms, 134,611,387/sec, 7 ns/op
tidwall(G): iter          1,000,000 ops in 10ms, 102,630,322/sec, 9 ns/op
tidwall(G): scan          1,000,000 ops in 8ms, 122,300,476/sec, 8 ns/op
tidwall(G): walk          1,000,000 ops in 3ms, 325,578,275/sec, 3 ns/op
go-arr:     for-loop      1,000,000 ops in 2ms, 637,704,264/sec, 1 ns/op
```
