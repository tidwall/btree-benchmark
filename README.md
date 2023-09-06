# btree-benchmark

Benchmark utility for the [tidwall/btree](https://github.com/tidwall/btree) Go package

- `google`: The [google/btree](https://github.com/google/btree) package (without generics)
- `google(G)`: The [google/btree](https://github.com/google/btree) package (generics)
- `tidwall`: The [tidwall/btree](https://github.com/tidwall/btree) package (without generics)
- `tidwall(G)`: The [tidwall/btree](https://github.com/tidwall/btree) package (generics)
- `tidwall(M)`: The [tidwall/btree](https://github.com/tidwall/btree) package (generics using the `btree.Map` type)

The following benchmarks were run on my 2021 Macbook Pro M1 Max 
using Go version 1.20.4.  
All items are key/value pairs where the key is a string filled with 16 random digits such as `5204379379828236`, and the value is the int64 representation of the key.
The degree is 32.  

```
degree=32, key=string (16 bytes), val=int64, count=1000000

** sequential set **
google:     set-seq           1,000,000 ops in 219ms, 4,568,211/sec, 218 ns/op, 56.9 MB, 59.6 bytes/op
google(G):  set-seq           1,000,000 ops in 167ms, 5,991,695/sec, 166 ns/op, 49.7 MB, 52.1 bytes/op
tidwall:    set-seq           1,000,000 ops in 139ms, 7,186,487/sec, 139 ns/op, 56.4 MB, 59.1 bytes/op
tidwall(G): set-seq           1,000,000 ops in 119ms, 8,419,681/sec, 118 ns/op, 49.2 MB, 51.6 bytes/op
tidwall(M): set-seq           1,000,000 ops in 102ms, 9,761,406/sec, 102 ns/op, 49.2 MB, 51.6 bytes/op
tidwall:    set-seq-hint      1,000,000 ops in 87ms, 11,538,644/sec, 86 ns/op, 56.4 MB, 59.1 bytes/op
tidwall(G): set-seq-hint      1,000,000 ops in 58ms, 17,338,334/sec, 57 ns/op, 49.2 MB, 51.6 bytes/op
tidwall:    load-seq          1,000,000 ops in 53ms, 18,909,653/sec, 52 ns/op, 56.4 MB, 59.1 bytes/op
tidwall(G): load-seq          1,000,000 ops in 32ms, 31,300,333/sec, 31 ns/op, 49.2 MB, 51.6 bytes/op
tidwall(M): load-seq          1,000,000 ops in 28ms, 36,346,564/sec, 27 ns/op, 49.2 MB, 51.6 bytes/op

** sequential get **
google:     get-seq           1,000,000 ops in 207ms, 4,835,779/sec, 206 ns/op
google(G):  get-seq           1,000,000 ops in 169ms, 5,927,806/sec, 168 ns/op
tidwall:    get-seq           1,000,000 ops in 156ms, 6,405,236/sec, 156 ns/op
tidwall(G): get-seq           1,000,000 ops in 125ms, 8,023,836/sec, 124 ns/op
tidwall(M): get-seq           1,000,000 ops in 102ms, 9,822,836/sec, 101 ns/op
tidwall:    get-seq-hint      1,000,000 ops in 84ms, 11,866,425/sec, 84 ns/op
tidwall(G): get-seq-hint      1,000,000 ops in 53ms, 18,696,903/sec, 53 ns/op

** random set **
google:     set-rand          1,000,000 ops in 1134ms, 881,661/sec, 1134 ns/op, 46.5 MB, 48.8 bytes/op
google(G):  set-rand          1,000,000 ops in 743ms, 1,345,894/sec, 743 ns/op, 34.6 MB, 36.3 bytes/op
tidwall:    set-rand          1,000,000 ops in 838ms, 1,193,269/sec, 838 ns/op, 46.5 MB, 48.8 bytes/op
tidwall(G): set-rand          1,000,000 ops in 571ms, 1,751,498/sec, 570 ns/op, 34.8 MB, 36.5 bytes/op
tidwall(M): set-rand          1,000,000 ops in 560ms, 1,784,123/sec, 560 ns/op, 34.8 MB, 36.5 bytes/op
tidwall:    set-rand-hint     1,000,000 ops in 888ms, 1,125,569/sec, 888 ns/op, 46.5 MB, 48.8 bytes/op
tidwall(G): set-rand-hint     1,000,000 ops in 606ms, 1,648,850/sec, 606 ns/op, 34.8 MB, 36.5 bytes/op
tidwall:    set-after-copy    1,000,000 ops in 974ms, 1,026,808/sec, 973 ns/op
tidwall(G): set-after-copy    1,000,000 ops in 612ms, 1,634,365/sec, 611 ns/op
tidwall:    load-rand         1,000,000 ops in 868ms, 1,151,739/sec, 868 ns/op, 46.5 MB, 48.8 bytes/op
tidwall(G): load-rand         1,000,000 ops in 594ms, 1,682,555/sec, 594 ns/op, 34.8 MB, 36.5 bytes/op
tidwall(M): load-rand         1,000,000 ops in 551ms, 1,816,466/sec, 550 ns/op, 34.8 MB, 36.5 bytes/op

** random get **
google:     get-rand          1,000,000 ops in 1523ms, 656,452/sec, 1523 ns/op
google(G):  get-rand          1,000,000 ops in 858ms, 1,165,715/sec, 857 ns/op
tidwall:    get-rand          1,000,000 ops in 1110ms, 901,008/sec, 1109 ns/op
tidwall(G): get-rand          1,000,000 ops in 655ms, 1,526,501/sec, 655 ns/op
tidwall(M): get-rand          1,000,000 ops in 625ms, 1,599,495/sec, 625 ns/op
tidwall:    get-rand-hint     1,000,000 ops in 1159ms, 863,083/sec, 1158 ns/op
tidwall(G): get-rand-hint     1,000,000 ops in 716ms, 1,395,710/sec, 716 ns/op

** sequential pivot **
Test getting 10 consecutive items starting at a pivot.
google:     ascend-seq        1,000,000 ops in 341ms, 2,929,622/sec, 341 ns/op
google:     descend-seq       1,000,000 ops in 431ms, 2,322,505/sec, 430 ns/op
google(G):  ascend-seq        1,000,000 ops in 213ms, 4,702,900/sec, 212 ns/op
google(G):  descend-seq       1,000,000 ops in 275ms, 3,638,184/sec, 274 ns/op
tidwall:    ascend-seq        1,000,000 ops in 197ms, 5,075,743/sec, 197 ns/op
tidwall:    descend-seq       1,000,000 ops in 205ms, 4,883,863/sec, 204 ns/op
tidwall:    ascend-seq-hint   1,000,000 ops in 121ms, 8,252,368/sec, 121 ns/op
tidwall:    descend-seq-hint  1,000,000 ops in 126ms, 7,950,834/sec, 125 ns/op
tidwall(G): ascend-seq        1,000,000 ops in 151ms, 6,631,924/sec, 150 ns/op
tidwall(G): descend-seq       1,000,000 ops in 153ms, 6,521,033/sec, 153 ns/op
tidwall(G): ascend-seq-hint   1,000,000 ops in 80ms, 12,487,434/sec, 80 ns/op
tidwall(G): descend-seq-hint  1,000,000 ops in 81ms, 12,329,924/sec, 81 ns/op
tidwall(G): iter-seq          1,000,000 ops in 213ms, 4,691,868/sec, 213 ns/op
tidwall(G): iter-seq-hint     1,000,000 ops in 138ms, 7,248,690/sec, 137 ns/op

** random pivot **
Test getting 10 consecutive items starting at a pivot.
google:     ascend-rand       1,000,000 ops in 1916ms, 521,904/sec, 1916 ns/op
google:     descend-rand      1,000,000 ops in 2351ms, 425,348/sec, 2351 ns/op
google(G):  ascend-rand       1,000,000 ops in 1043ms, 959,150/sec, 1042 ns/op
google(G):  descend-rand      1,000,000 ops in 1173ms, 852,528/sec, 1172 ns/op
tidwall:    ascend-rand       1,000,000 ops in 1166ms, 857,545/sec, 1166 ns/op
tidwall:    descend-rand      1,000,000 ops in 1164ms, 858,910/sec, 1164 ns/op
tidwall:    ascend-rand-hint  1,000,000 ops in 1211ms, 825,455/sec, 1211 ns/op
tidwall:    descend-rand-hint 1,000,000 ops in 1224ms, 817,205/sec, 1223 ns/op
tidwall(G): ascend-rand       1,000,000 ops in 726ms, 1,376,815/sec, 726 ns/op
tidwall(G): descend-rand      1,000,000 ops in 712ms, 1,403,801/sec, 712 ns/op
tidwall(G): ascend-rand-hint  1,000,000 ops in 771ms, 1,296,249/sec, 771 ns/op
tidwall(G): descend-rand-hint 1,000,000 ops in 768ms, 1,302,896/sec, 767 ns/op
tidwall(G): iter-rand         1,000,000 ops in 811ms, 1,232,549/sec, 811 ns/op
tidwall(G): iter-rand-hint    1,000,000 ops in 848ms, 1,179,822/sec, 847 ns/op

** scan **
Test scanning over every item in the tree
google:     ascend            1,000,000 ops in 10ms, 97,801,108/sec, 10 ns/op
google(G):  ascend            1,000,000 ops in 11ms, 89,880,910/sec, 11 ns/op
tidwall:    ascend            1,000,000 ops in 8ms, 124,326,554/sec, 8 ns/op
tidwall(G): scan              1,000,000 ops in 9ms, 108,123,701/sec, 9 ns/op
tidwall(G): walk              1,000,000 ops in 3ms, 331,945,689/sec, 3 ns/op
tidwall(G): iter              1,000,000 ops in 11ms, 94,063,785/sec, 10 ns/op
```
