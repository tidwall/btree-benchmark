package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"

	gbtree "github.com/google/btree"
	tbtree "github.com/tidwall/btree"
	"github.com/tidwall/lotsa"
)

type keyT string
type valT int64

type itemT struct {
	key keyT
	val valT
}

func int64ToItemT(i int64) itemT {
	return itemT{
		key: keyT(fmt.Sprintf("%016d", i)),
		val: valT(i),
	}
}

func (item itemT) Less(other gbtree.Item) bool {
	return item.key < other.(itemT).key
}

func lessG(a, b itemT) bool {
	return a.key < b.key
}

func less(a, b interface{}) bool {
	return a.(itemT).key < b.(itemT).key
}

func newTBTree(degree int) *tbtree.BTree {
	return tbtree.NewOptions(less, tbtree.Options{
		NoLocks: true,
		Degree:  degree,
	})
}
func newTBTreeG(degree int) *tbtree.BTreeG[itemT] {
	return tbtree.NewBTreeGOptions(lessG, tbtree.Options{
		NoLocks: true,
		Degree:  degree,
	})
}
func newTBTreeM(degree int) *tbtree.Map[keyT, valT] {
	return tbtree.NewMap[keyT, valT](degree)
}
func newGBTree(degree int) *gbtree.BTree {
	return gbtree.New(degree)
}
func newGBTreeG(degree int) *gbtree.BTreeG[itemT] {
	return gbtree.NewG(degree, lessG)
}

func main() {
	N := 1_000_000
	degree := 32
	flag.IntVar(&N, "count", N, "number of items")
	flag.IntVar(&degree, "degree", degree, "B-tree degree")
	flag.Parse()

	items := make([]itemT, N)
	itemsM := make(map[int64]bool)
	for i := 0; i < N; i++ {
		for {
			key := rand.Int63n(10000000000000000)
			if !itemsM[key] {
				itemsM[key] = true
				items[i] = int64ToItemT(key)
				if len(items[i].key) != 16 {
					panic("!")
				}
				break
			}
		}
	}

	lotsa.Output = os.Stdout
	lotsa.MemUsage = true

	sortInts := func() {
		sort.Slice(items, func(i, j int) bool {
			return less(items[i], items[j])
		})
	}

	shuffleInts := func() {
		for i := range items {
			j := rand.Intn(i + 1)
			items[i], items[j] = items[j], items[i]
		}
	}

	gtr := newGBTree(degree)
	gtrG := newGBTreeG(degree)
	ttr := newTBTree(degree)
	ttrG := newTBTreeG(degree)
	ttrM := newTBTreeM(degree)

	withSeq := true
	withRand := true
	withRandSet := true
	withRange := true
	withHints := true

	fmt.Printf("\ndegree=%d, key=string (16 bytes), val=int64, count=%d\n",
		degree, N)

	var hint tbtree.PathHint
	var hintG tbtree.PathHint
	var arr []itemT

	if withSeq {
		println()
		println("** sequential set **")
		sortInts()

		// google
		print("google:     set-seq        ")
		gtr = newGBTree(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			gtr.ReplaceOrInsert(items[i])
		})

		print("google(G):  set-seq        ")
		gtrG = newGBTreeG(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			gtrG.ReplaceOrInsert(items[i])
		})

		// non-generics tidwall
		print("tidwall:    set-seq        ")
		ttr = newTBTree(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttr.Set(items[i])
		})
		print("tidwall(G): set-seq        ")
		ttrG = newTBTreeG(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrG.Set(items[i])
		})
		print("tidwall(M): set-seq        ")
		ttrM = newTBTreeM(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrM.Set(items[i].key, items[i].val)
		})

		if withHints {
			print("tidwall:    set-seq-hint   ")
			ttr = newTBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.SetHint(items[i], &hint)
			})
			print("tidwall(G): set-seq-hint   ")
			ttrG = newTBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.SetHint(items[i], &hintG)
			})
		}
		print("tidwall:    load-seq       ")
		ttr = newTBTree(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttr.Load(items[i])
		})
		print("tidwall(G): load-seq       ")
		ttrG = newTBTreeG(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrG.Load(items[i])
		})
		print("tidwall(M): load-seq       ")
		ttrM = newTBTreeM(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrM.Load(items[i].key, items[i].val)
		})

		// go array
		print("go-arr:     append         ")
		lotsa.Ops(N, 1, func(i, _ int) {
			arr = append(arr, items[i])
		})

		println()
		println("** sequential get **")
		sortInts()

		print("google:     get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := gtr.Get(items[i])
			if re == nil {
				panic(re)
			}
		})
		print("google(G):  get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := gtrG.Get(items[i])
			if !ok {
				panic(re)
			}
		})
		print("tidwall:    get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := ttr.Get(items[i])
			if re == nil {
				panic(re)
			}
		})
		print("tidwall(G): get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrG.Get(items[i])
			if !ok {
				panic(re)
			}
		})
		print("tidwall(M): get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrM.Get(items[i].key)
			if !ok {
				panic(re)
			}
		})
		if withHints {
			print("tidwall:    get-seq-hint   ")
			lotsa.Ops(N, 1, func(i, _ int) {
				re := ttr.GetHint(items[i], &hint)
				if re == nil {
					panic(re)
				}
			})
			print("tidwall(G): get-seq-hint   ")
			lotsa.Ops(N, 1, func(i, _ int) {
				re, ok := ttrG.GetHint(items[i], &hintG)
				if !ok {
					panic(re)
				}
			})
		}
	}

	if withRand {
		if withRandSet {
			println()
			println("** random set **")
			shuffleInts()
			print("google:     set-rand       ")
			gtr = newGBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				gtr.ReplaceOrInsert(items[i])
			})
			print("google(G):  set-rand       ")
			gtrG = newGBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				gtrG.ReplaceOrInsert(items[i])
			})
			print("tidwall:    set-rand       ")
			ttr = newTBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.Set(items[i])
			})
			print("tidwall(G): set-rand       ")
			ttrG = newTBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.Set(items[i])
			})
			print("tidwall(M): set-rand       ")
			ttrM = newTBTreeM(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrM.Set(items[i].key, items[i].val)
			})
			if withHints {
				print("tidwall:    set-rand-hint  ")
				ttr = newTBTree(degree)
				lotsa.Ops(N, 1, func(i, _ int) {
					ttr.SetHint(items[i], &hint)
				})
				print("tidwall(G): set-rand-hint  ")
				ttrG = newTBTreeG(degree)
				lotsa.Ops(N, 1, func(i, _ int) {
					ttrG.SetHint(items[i], &hintG)
				})
			}
			print("tidwall:    set-after-copy ")
			ttr = ttr.Copy()
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.Set(items[i])
			})
			print("tidwall(G): set-after-copy ")
			ttrG = ttrG.Copy()
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.Set(items[i])
			})
			print("tidwall:    load-rand      ")
			ttr = newTBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.Load(items[i])
			})
			print("tidwall(G): load-rand      ")
			ttrG = newTBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.Load(items[i])
			})
			ttrM = newTBTreeM(degree)
			print("tidwall(M): load-rand      ")
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrM.Load(items[i].key, items[i].val)
			})
		}
		println()
		println("** random get **")

		shuffleInts()
		gtr = newGBTree(degree)
		gtrG = newGBTreeG(degree)
		ttr = newTBTree(degree)
		ttrM = newTBTreeM(degree)
		ttrG = newTBTreeG(degree)
		for _, item := range items {
			gtr.ReplaceOrInsert(item)
			gtrG.ReplaceOrInsert(item)
			ttrG.Set(item)
			ttr.Set(item)
			ttrM.Set(item.key, item.val)
		}
		shuffleInts()

		print("google:     get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := gtr.Get(items[i])
			if re == nil {
				panic(re)
			}
		})
		print("google(G):  get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := gtrG.Get(items[i])
			if !ok {
				panic(re)
			}
		})
		print("tidwall:    get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := ttr.Get(items[i])
			if re == nil {
				panic(re)
			}
		})
		print("tidwall(G): get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrG.Get(items[i])
			if !ok {
				panic(re)
			}
		})
		print("tidwall(M): get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrM.Get(items[i].key)
			if !ok {
				panic(re)
			}
		})
		if withHints {
			print("tidwall:    get-rand-hint  ")
			lotsa.Ops(N, 1, func(i, _ int) {
				re := ttr.GetHint(items[i], &hint)
				if re == nil {
					panic(re)
				}
			})
			print("tidwall(G): get-rand-hint  ")
			lotsa.Ops(N, 1, func(i, _ int) {
				re, ok := ttrG.GetHint(items[i], &hintG)
				if !ok {
					panic(re)
				}
			})
		}
	}

	if withRange {
		println()
		println("** range **")
		print("google:     ascend        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				gtr.Ascend(func(item gbtree.Item) bool {
					return true
				})
			}
		})
		print("google(G):  ascend        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				gtrG.Ascend(func(item itemT) bool {
					return true
				})
			}
		})
		print("tidwall:    ascend        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				ttr.Ascend(nil, func(item interface{}) bool {
					return true
				})
			}
		})
		print("tidwall(G): iter          ")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				iter := ttrG.Iter()
				for ok := iter.First(); ok; ok = iter.Next() {
				}
				iter.Release()
			}
		})
		print("tidwall(G): scan          ")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				ttrG.Scan(func(item itemT) bool {
					return true
				})
			}
		})
		print("tidwall(G): walk          ")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				ttrG.Walk(func(items []itemT) bool {
					for j := 0; j < len(items); j++ {
					}
					return true
				})
			}
		})

		print("go-arr:     for-loop      ")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				for j := 0; j < len(arr); j++ {
				}
			}
		})
	}
}
