# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `google`: The [google/btree](https://github.com/google/btree) package
- `go-arr`: A simple Go array

The following benchmarks were run on my 2019 Macbook Pro (2.4 GHz 8-Core Intel Core i9) 
using Go 1.15.8. The items are simple 8-byte ints.

```
** sequential set **
google:  set-seq        1,000,000 ops in 160ms, 6,262,097/sec, 159 ns/op, 31.0 MB, 32 bytes/op
tidwall: set-seq        1,000,000 ops in 142ms, 7,020,721/sec, 142 ns/op, 36.6 MB, 38 bytes/op
tidwall: set-seq-hint   1,000,000 ops in 87ms, 11,503,315/sec, 86 ns/op, 36.6 MB, 38 bytes/op
tidwall: load-seq       1,000,000 ops in 37ms, 27,177,242/sec, 36 ns/op, 36.6 MB, 38 bytes/op
go-arr:  append         1,000,000 ops in 49ms, 20,574,760/sec, 48 ns/op

** random set **
google:  set-rand       1,000,000 ops in 606ms, 1,649,921/sec, 606 ns/op, 21.5 MB, 22 bytes/op
tidwall: set-rand       1,000,000 ops in 543ms, 1,841,590/sec, 543 ns/op, 26.7 MB, 27 bytes/op
tidwall: set-rand-hint  1,000,000 ops in 573ms, 1,745,624/sec, 572 ns/op, 26.4 MB, 27 bytes/op
tidwall: set-again      1,000,000 ops in 452ms, 2,212,581/sec, 451 ns/op, 27.1 MB, 28 bytes/op
tidwall: set-after-copy 1,000,000 ops in 472ms, 2,117,457/sec, 472 ns/op, 27.9 MB, 29 bytes/op
tidwall: load-rand      1,000,000 ops in 551ms, 1,816,498/sec, 550 ns/op, 26.1 MB, 27 bytes/op

** sequential get **
google:  get-seq        1,000,000 ops in 133ms, 7,497,604/sec, 133 ns/op
tidwall: get-seq        1,000,000 ops in 110ms, 9,082,972/sec, 110 ns/op
tidwall: get-seq-hint   1,000,000 ops in 55ms, 18,289,945/sec, 54 ns/op

** random get **
google:  get-rand       1,000,000 ops in 149ms, 6,704,337/sec, 149 ns/op
tidwall: get-rand       1,000,000 ops in 131ms, 7,616,296/sec, 131 ns/op
tidwall: get-rand-hint  1,000,000 ops in 216ms, 4,632,532/sec, 215 ns/op
```
