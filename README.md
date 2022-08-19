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
The degrees is 128. 
The keys are strings filled with 16 random digits such as `5204379379828236`.

```
** sequential set **
google:     set-seq        1,000,000 ops in 218ms, 4,595,330/sec, 217 ns/op, 46.6 MB, 48 bytes/op
google(G):  set-seq        1,000,000 ops in 157ms, 6,358,466/sec, 157 ns/op, 31.4 MB, 32 bytes/op
tidwall:    set-seq        1,000,000 ops in 118ms, 8,443,331/sec, 118 ns/op, 46.5 MB, 48 bytes/op
tidwall(G): set-seq        1,000,000 ops in 96ms, 10,446,173/sec, 95 ns/op, 31.2 MB, 32 bytes/op
tidwall(M): set-seq        1,000,000 ops in 76ms, 13,224,699/sec, 75 ns/op, 31.2 MB, 32 bytes/op
tidwall:    set-seq-hint   1,000,000 ops in 63ms, 15,770,310/sec, 63 ns/op, 46.5 MB, 48 bytes/op
tidwall(G): set-seq-hint   1,000,000 ops in 44ms, 22,915,566/sec, 43 ns/op, 31.2 MB, 32 bytes/op
tidwall:    load-seq       1,000,000 ops in 38ms, 26,441,240/sec, 37 ns/op, 46.5 MB, 48 bytes/op
tidwall(G): load-seq       1,000,000 ops in 21ms, 48,546,439/sec, 20 ns/op, 31.2 MB, 32 bytes/op
tidwall(M): load-seq       1,000,000 ops in 17ms, 58,555,867/sec, 17 ns/op, 31.2 MB, 32 bytes/op
go-arr:     append         1,000,000 ops in 13ms, 76,110,741/sec, 13 ns/op, 16.9 MB, 17 bytes/op

** sequential get **
google:     get-seq        1,000,000 ops in 194ms, 5,155,076/sec, 193 ns/op
google(G):  get-seq        1,000,000 ops in 143ms, 6,991,464/sec, 143 ns/op
tidwall:    get-seq        1,000,000 ops in 132ms, 7,574,612/sec, 132 ns/op
tidwall(G): get-seq        1,000,000 ops in 98ms, 10,190,200/sec, 98 ns/op
tidwall(M): get-seq        1,000,000 ops in 87ms, 11,452,654/sec, 87 ns/op
tidwall:    get-seq-hint   1,000,000 ops in 57ms, 17,559,737/sec, 56 ns/op
tidwall(G): get-seq-hint   1,000,000 ops in 38ms, 26,214,610/sec, 38 ns/op

** random set **
google:     set-rand       1,000,000 ops in 1128ms, 886,399/sec, 1128 ns/op, 37.6 MB, 39 bytes/op
google(G):  set-rand       1,000,000 ops in 731ms, 1,368,387/sec, 730 ns/op, 22.3 MB, 23 bytes/op
tidwall:    set-rand       1,000,000 ops in 826ms, 1,210,910/sec, 825 ns/op, 37.6 MB, 39 bytes/op
tidwall(G): set-rand       1,000,000 ops in 532ms, 1,879,893/sec, 531 ns/op, 22.3 MB, 23 bytes/op
tidwall(M): set-rand       1,000,000 ops in 505ms, 1,978,932/sec, 505 ns/op, 22.3 MB, 23 bytes/op
tidwall:    set-rand-hint  1,000,000 ops in 853ms, 1,171,910/sec, 853 ns/op, 37.6 MB, 39 bytes/op
tidwall(G): set-rand-hint  1,000,000 ops in 632ms, 1,582,204/sec, 632 ns/op, 22.3 MB, 23 bytes/op
tidwall:    set-after-copy 1,000,000 ops in 902ms, 1,108,986/sec, 901 ns/op
tidwall(G): set-after-copy 1,000,000 ops in 527ms, 1,896,504/sec, 527 ns/op
tidwall:    load-rand      1,000,000 ops in 824ms, 1,213,410/sec, 824 ns/op, 37.6 MB, 39 bytes/op
tidwall(G): load-rand      1,000,000 ops in 556ms, 1,799,942/sec, 555 ns/op, 22.3 MB, 23 bytes/op
tidwall(M): load-rand      1,000,000 ops in 517ms, 1,933,136/sec, 517 ns/op, 22.3 MB, 23 bytes/op

** random get **
google:     get-rand       1,000,000 ops in 1479ms, 676,068/sec, 1479 ns/op
google(G):  get-rand       1,000,000 ops in 715ms, 1,398,291/sec, 715 ns/op
tidwall:    get-rand       1,000,000 ops in 1041ms, 961,027/sec, 1040 ns/op
tidwall(G): get-rand       1,000,000 ops in 564ms, 1,774,469/sec, 563 ns/op
tidwall(M): get-rand       1,000,000 ops in 508ms, 1,967,797/sec, 508 ns/op
tidwall:    get-rand-hint  1,000,000 ops in 1068ms, 935,954/sec, 1068 ns/op
tidwall(G): get-rand-hint  1,000,000 ops in 604ms, 1,656,866/sec, 603 ns/op

** range **
google:     ascend        1,000,000 ops in 6ms, 179,013,624/sec, 5 ns/op
google(G):  ascend        1,000,000 ops in 5ms, 184,666,511/sec, 5 ns/op
tidwall:    ascend        1,000,000 ops in 5ms, 217,367,677/sec, 4 ns/op
tidwall(G): iter          1,000,000 ops in 6ms, 178,942,894/sec, 5 ns/op
tidwall(G): scan          1,000,000 ops in 4ms, 230,403,651/sec, 4 ns/op
tidwall(G): walk          1,000,000 ops in 3ms, 387,390,441/sec, 2 ns/op
go-arr:     for-loop      1,000,000 ops in 2ms, 641,008,383/sec, 1 ns/op
```
