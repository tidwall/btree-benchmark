# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package
- `google`: The [google/btree](https://github.com/google/btree) package
- `go-arr`: A simple Go array

```
** sequential set **
google:  set-seq       1,000,000 ops in 175ms, 5,700,922/sec, 175 ns/op, 33.1 MB, 34 bytes/op
tidwall: set-seq       1,000,000 ops in 143ms, 6,996,275/sec, 142 ns/op, 30.9 MB, 32 bytes/op
tidwall: set-seq-hint  1,000,000 ops in 65ms, 15,441,082/sec, 64 ns/op, 30.9 MB, 32 bytes/op
tidwall: load-seq      1,000,000 ops in 19ms, 53,242,398/sec, 18 ns/op, 30.9 MB, 32 bytes/op
go-arr:  append        1,000,000 ops in 52ms, 19,153,714/sec, 52 ns/op, 41.3 MB, 43 bytes/op

** random set **
google:  set-rand      1,000,000 ops in 662ms, 1,509,924/sec, 662 ns/op, 32.1 MB, 33 bytes/op
tidwall: set-rand      1,000,000 ops in 589ms, 1,697,471/sec, 589 ns/op, 22.5 MB, 23 bytes/op
tidwall: set-rand-hint 1,000,000 ops in 592ms, 1,688,184/sec, 592 ns/op, 22.2 MB, 23 bytes/op
tidwall: load-rand     1,000,000 ops in 578ms, 1,728,932/sec, 578 ns/op, 22.3 MB, 23 bytes/op

** sequential get **
google:  get-seq       1,000,000 ops in 135ms, 7,414,046/sec, 134 ns/op
tidwall: get-seq       1,000,000 ops in 111ms, 8,995,090/sec, 111 ns/op
tidwall: get-seq-hint  1,000,000 ops in 56ms, 18,017,397/sec, 55 ns/op

** random get **
google:  get-rand      1,000,000 ops in 161ms, 6,199,818/sec, 161 ns/op
tidwall: get-rand      1,000,000 ops in 139ms, 7,214,017/sec, 138 ns/op
tidwall: get-rand-hint 1,000,000 ops in 191ms, 5,243,833/sec, 190 ns/op
```
