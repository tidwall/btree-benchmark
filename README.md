# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `google`: The [google/btree](https://github.com/google/btree) package (without generics)
- `google(G)`: The [google/btree](https://github.com/google/btree) package (generics)
- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package (without generics)
- `tidwall(G)`: The [tidwall/btree](https://github.com/tidwall/btree) package (generics)
- `tidwall(M)`: The [tidwall/btree](https://github.com/tidwall/btree) package (generics using the `btree.Map` type)
- `go-arr`: A simple Go array

The following benchmarks were run on my 2021 Macbook Pro M1 Max 
using Go version 1.19.3.  
The keys are strings filled with 16 random digits such as `5204379379828236`.  
The degrees is 32.  

```
degree=32, key=string (16 bytes), count=1000000

** sequential set **
google:     set-seq        1,000,000 ops in 227ms, 4,403,863/sec, 227 ns/op, 49.2 MB, 51 bytes/op
google(G):  set-seq        1,000,000 ops in 156ms, 6,411,038/sec, 155 ns/op, 34.0 MB, 35 bytes/op
tidwall:    set-seq        1,000,000 ops in 146ms, 6,867,149/sec, 145 ns/op, 48.7 MB, 51 bytes/op
tidwall(G): set-seq        1,000,000 ops in 115ms, 8,681,032/sec, 115 ns/op, 33.5 MB, 35 bytes/op
tidwall(M): set-seq        1,000,000 ops in 89ms, 11,240,649/sec, 88 ns/op, 33.5 MB, 35 bytes/op
tidwall:    set-seq-hint   1,000,000 ops in 76ms, 13,143,295/sec, 76 ns/op, 48.7 MB, 51 bytes/op
tidwall(G): set-seq-hint   1,000,000 ops in 56ms, 17,827,496/sec, 56 ns/op, 33.5 MB, 35 bytes/op
tidwall:    load-seq       1,000,000 ops in 47ms, 21,349,635/sec, 46 ns/op, 48.7 MB, 51 bytes/op
tidwall(G): load-seq       1,000,000 ops in 29ms, 34,432,890/sec, 29 ns/op, 33.5 MB, 35 bytes/op
tidwall(M): load-seq       1,000,000 ops in 26ms, 38,911,577/sec, 25 ns/op, 33.5 MB, 35 bytes/op
go-arr:     append         1,000,000 ops in 13ms, 74,212,879/sec, 13 ns/op, 16.9 MB, 17 bytes/op

** sequential get **
google:     get-seq        1,000,000 ops in 208ms, 4,818,820/sec, 207 ns/op
google(G):  get-seq        1,000,000 ops in 153ms, 6,539,750/sec, 152 ns/op
tidwall:    get-seq        1,000,000 ops in 142ms, 7,038,170/sec, 142 ns/op
tidwall(G): get-seq        1,000,000 ops in 108ms, 9,228,181/sec, 108 ns/op
tidwall(M): get-seq        1,000,000 ops in 88ms, 11,328,791/sec, 88 ns/op
tidwall:    get-seq-hint   1,000,000 ops in 68ms, 14,768,458/sec, 67 ns/op
tidwall(G): get-seq-hint   1,000,000 ops in 46ms, 21,667,826/sec, 46 ns/op

** random set **
google:     set-rand       1,000,000 ops in 1040ms, 961,554/sec, 1039 ns/op, 38.8 MB, 40 bytes/op
google(G):  set-rand       1,000,000 ops in 656ms, 1,524,490/sec, 655 ns/op, 23.5 MB, 24 bytes/op
tidwall:    set-rand       1,000,000 ops in 778ms, 1,285,492/sec, 777 ns/op, 38.8 MB, 40 bytes/op
tidwall(G): set-rand       1,000,000 ops in 511ms, 1,957,928/sec, 510 ns/op, 23.5 MB, 24 bytes/op
tidwall(M): set-rand       1,000,000 ops in 482ms, 2,073,218/sec, 482 ns/op, 23.5 MB, 24 bytes/op
tidwall:    set-rand-hint  1,000,000 ops in 905ms, 1,105,298/sec, 904 ns/op, 38.8 MB, 40 bytes/op
tidwall(G): set-rand-hint  1,000,000 ops in 591ms, 1,692,256/sec, 590 ns/op, 23.5 MB, 24 bytes/op
tidwall:    set-after-copy 1,000,000 ops in 949ms, 1,053,795/sec, 948 ns/op
tidwall(G): set-after-copy 1,000,000 ops in 577ms, 1,733,692/sec, 576 ns/op
tidwall:    load-rand      1,000,000 ops in 833ms, 1,200,349/sec, 833 ns/op, 38.8 MB, 40 bytes/op
tidwall(G): load-rand      1,000,000 ops in 566ms, 1,765,325/sec, 566 ns/op, 23.5 MB, 24 bytes/op
tidwall(M): load-rand      1,000,000 ops in 524ms, 1,907,006/sec, 524 ns/op, 23.5 MB, 24 bytes/op

** random get **
google:     get-rand       1,000,000 ops in 1459ms, 685,468/sec, 1458 ns/op
google(G):  get-rand       1,000,000 ops in 785ms, 1,273,144/sec, 785 ns/op
tidwall:    get-rand       1,000,000 ops in 1049ms, 952,865/sec, 1049 ns/op
tidwall(G): get-rand       1,000,000 ops in 577ms, 1,734,046/sec, 576 ns/op
tidwall(M): get-rand       1,000,000 ops in 538ms, 1,858,632/sec, 538 ns/op
tidwall:    get-rand-hint  1,000,000 ops in 1084ms, 922,717/sec, 1083 ns/op
tidwall(G): get-rand-hint  1,000,000 ops in 643ms, 1,555,637/sec, 642 ns/op

** range **
google:     ascend        1,000,000 ops in 9ms, 106,277,921/sec, 9 ns/op
google(G):  ascend        1,000,000 ops in 9ms, 107,128,513/sec, 9 ns/op
tidwall:    ascend        1,000,000 ops in 7ms, 133,345,921/sec, 7 ns/op
tidwall(G): iter          1,000,000 ops in 9ms, 110,793,601/sec, 9 ns/op
tidwall(G): scan          1,000,000 ops in 8ms, 132,896,246/sec, 7 ns/op
tidwall(G): walk          1,000,000 ops in 3ms, 323,001,693/sec, 3 ns/op
go-arr:     for-loop      1,000,000 ops in 2ms, 623,506,312/sec, 1 ns/op
```
