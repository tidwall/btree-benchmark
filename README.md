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
The degree is 32.  

```
degree=32, key=string (16 bytes), count=1000000

** sequential set **
google:     set-seq        1,000,000 ops in 207ms, 4,825,592/sec, 207 ns/op, 49.2 MB, 51.6 bytes/op
google(G):  set-seq        1,000,000 ops in 159ms, 6,301,736/sec, 158 ns/op, 34.0 MB, 35.6 bytes/op
tidwall:    set-seq        1,000,000 ops in 134ms, 7,465,291/sec, 133 ns/op, 48.7 MB, 51.1 bytes/op
tidwall(G): set-seq        1,000,000 ops in 108ms, 9,244,301/sec, 108 ns/op, 33.5 MB, 35.1 bytes/op
tidwall(M): set-seq        1,000,000 ops in 87ms, 11,454,912/sec, 87 ns/op, 33.5 MB, 35.1 bytes/op
tidwall:    set-seq-hint   1,000,000 ops in 72ms, 13,801,161/sec, 72 ns/op, 48.7 MB, 51.1 bytes/op
tidwall(G): set-seq-hint   1,000,000 ops in 54ms, 18,357,409/sec, 54 ns/op, 33.5 MB, 35.1 bytes/op
tidwall:    load-seq       1,000,000 ops in 45ms, 22,339,567/sec, 44 ns/op, 48.7 MB, 51.1 bytes/op
tidwall(G): load-seq       1,000,000 ops in 29ms, 34,439,017/sec, 29 ns/op, 33.5 MB, 35.1 bytes/op
tidwall(M): load-seq       1,000,000 ops in 26ms, 38,972,364/sec, 25 ns/op, 33.5 MB, 35.1 bytes/op
go-arr:     append         1,000,000 ops in 13ms, 76,438,474/sec, 13 ns/op, 16.9 MB, 17.8 bytes/op

** sequential get **
google:     get-seq        1,000,000 ops in 191ms, 5,227,249/sec, 191 ns/op
google(G):  get-seq        1,000,000 ops in 148ms, 6,776,054/sec, 147 ns/op
tidwall:    get-seq        1,000,000 ops in 135ms, 7,397,461/sec, 135 ns/op
tidwall(G): get-seq        1,000,000 ops in 107ms, 9,382,590/sec, 106 ns/op
tidwall(M): get-seq        1,000,000 ops in 88ms, 11,416,214/sec, 87 ns/op
tidwall:    get-seq-hint   1,000,000 ops in 66ms, 15,059,343/sec, 66 ns/op
tidwall(G): get-seq-hint   1,000,000 ops in 46ms, 21,673,364/sec, 46 ns/op

** random set **
google:     set-rand       1,000,000 ops in 968ms, 1,032,568/sec, 968 ns/op, 38.8 MB, 40.7 bytes/op
google(G):  set-rand       1,000,000 ops in 625ms, 1,601,237/sec, 624 ns/op, 23.5 MB, 24.7 bytes/op
tidwall:    set-rand       1,000,000 ops in 755ms, 1,323,978/sec, 755 ns/op, 38.8 MB, 40.6 bytes/op
tidwall(G): set-rand       1,000,000 ops in 513ms, 1,947,800/sec, 513 ns/op, 23.5 MB, 24.6 bytes/op
tidwall(M): set-rand       1,000,000 ops in 476ms, 2,099,771/sec, 476 ns/op, 23.5 MB, 24.6 bytes/op
tidwall:    set-rand-hint  1,000,000 ops in 950ms, 1,052,387/sec, 950 ns/op, 38.8 MB, 40.6 bytes/op
tidwall(G): set-rand-hint  1,000,000 ops in 556ms, 1,796,971/sec, 556 ns/op, 23.5 MB, 24.6 bytes/op
tidwall:    set-after-copy 1,000,000 ops in 893ms, 1,120,163/sec, 892 ns/op
tidwall(G): set-after-copy 1,000,000 ops in 547ms, 1,828,304/sec, 546 ns/op
tidwall:    load-rand      1,000,000 ops in 751ms, 1,331,323/sec, 751 ns/op, 38.8 MB, 40.6 bytes/op
tidwall(G): load-rand      1,000,000 ops in 512ms, 1,951,348/sec, 512 ns/op, 23.5 MB, 24.6 bytes/op
tidwall(M): load-rand      1,000,000 ops in 504ms, 1,985,926/sec, 503 ns/op, 23.5 MB, 24.6 bytes/op

** random get **
google:     get-rand       1,000,000 ops in 1486ms, 672,937/sec, 1486 ns/op
google(G):  get-rand       1,000,000 ops in 743ms, 1,345,497/sec, 743 ns/op
tidwall:    get-rand       1,000,000 ops in 1068ms, 936,318/sec, 1068 ns/op
tidwall(G): get-rand       1,000,000 ops in 559ms, 1,788,164/sec, 559 ns/op
tidwall(M): get-rand       1,000,000 ops in 548ms, 1,825,106/sec, 547 ns/op
tidwall:    get-rand-hint  1,000,000 ops in 1070ms, 934,912/sec, 1069 ns/op
tidwall(G): get-rand-hint  1,000,000 ops in 632ms, 1,582,180/sec, 632 ns/op

** range **
google:     ascend        1,000,000 ops in 9ms, 117,320,397/sec, 8 ns/op
google(G):  ascend        1,000,000 ops in 8ms, 125,864,009/sec, 7 ns/op
tidwall:    ascend        1,000,000 ops in 8ms, 127,609,339/sec, 7 ns/op
tidwall(G): iter          1,000,000 ops in 8ms, 125,250,501/sec, 7 ns/op
tidwall(G): scan          1,000,000 ops in 6ms, 156,598,676/sec, 6 ns/op
tidwall(G): walk          1,000,000 ops in 3ms, 335,664,298/sec, 2 ns/op
go-arr:     for-loop      1,000,000 ops in 2ms, 621,874,304/sec, 1 ns/op
```
