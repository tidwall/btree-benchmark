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

// type itemT int

// func intToItemT(i int) itemT { return itemT(i) }

type itemT string

func intToItemT(i int) itemT { return itemT(fmt.Sprintf("%016d", i)) }

func (i itemT) Less(other gbtree.Item) bool {
	return i < other.(itemT)
}

func lessG(a, b itemT) bool {
	return a < b
}

func less(a, b interface{}) bool {
	return a.(itemT) < b.(itemT)
}

func newTBTree(degree int) *tbtree.BTree {
	return tbtree.NewOptions(less, tbtree.Options{
		NoLocks: true,
		Degree:  degree,
	})
}
func newTBTreeG(degree int) *tbtree.Generic[itemT] {
	return tbtree.NewGenericOptions[itemT](lessG, tbtree.Options{
		NoLocks: true,
		Degree:  degree,
	})
}
func newTBTreeM(degree int) *tbtree.Map[itemT, struct{}] {
	return tbtree.NewMap[itemT, struct{}](degree)
}
func newGBTree(degree int) *gbtree.BTree {
	return gbtree.New(degree)
}
func newGBTreeG(degree int) *gbtree.BTreeG[itemT] {
	return gbtree.NewG[itemT](degree, lessG)
}

func main() {
	N := 1_000_000
	degree := 32
	flag.IntVar(&N, "count", N, "number of items")
	flag.IntVar(&degree, "degree", degree, "B-tree degree")
	flag.Parse()

	keys := make([]itemT, N)
	keysM := make(map[int]bool)
	for i := 0; i < N; i++ {
		for {
			key := rand.Intn(10000000000000000)
			if !keysM[key] {
				keysM[key] = true
				keys[i] = intToItemT(key)
				if len(keys[i]) != 16 {
					panic("!")
				}
				break
			}
		}
	}

	lotsa.Output = os.Stdout
	lotsa.MemUsage = true

	sortInts := func() {
		sort.Slice(keys, func(i, j int) bool {
			return less(keys[i], keys[j])
		})
	}

	shuffleInts := func() {
		for i := range keys {
			j := rand.Intn(i + 1)
			keys[i], keys[j] = keys[j], keys[i]
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

	fmt.Printf("\ndegree=%d, key=string (16 bytes), count=%d\n",
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
			gtr.ReplaceOrInsert(keys[i])
		})

		print("google(G):  set-seq        ")
		gtrG = newGBTreeG(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			gtrG.ReplaceOrInsert(keys[i])
		})

		// non-generics tidwall
		print("tidwall:    set-seq        ")
		ttr = newTBTree(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttr.Set(keys[i])
		})
		print("tidwall(G): set-seq        ")
		ttrG = newTBTreeG(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrG.Set(keys[i])
		})
		print("tidwall(M): set-seq        ")
		ttrM = newTBTreeM(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrM.Set(keys[i], struct{}{})
		})

		if withHints {
			print("tidwall:    set-seq-hint   ")
			ttr = newTBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.SetHint(keys[i], &hint)
			})
			print("tidwall(G): set-seq-hint   ")
			ttrG = newTBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.SetHint(keys[i], &hintG)
			})
		}
		print("tidwall:    load-seq       ")
		ttr = newTBTree(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttr.Load(keys[i])
		})
		print("tidwall(G): load-seq       ")
		ttrG = newTBTreeG(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrG.Load(keys[i])
		})
		print("tidwall(M): load-seq       ")
		ttrM = newTBTreeM(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrM.Load(keys[i], struct{}{})
		})

		// go array
		print("go-arr:     append         ")
		lotsa.Ops(N, 1, func(i, _ int) {
			arr = append(arr, keys[i])
		})

		println()
		println("** sequential get **")
		sortInts()

		print("google:     get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := gtr.Get(keys[i])
			if re == nil {
				panic(re)
			}
		})
		print("google(G):  get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := gtrG.Get(keys[i])
			if !ok {
				panic(re)
			}
		})
		print("tidwall:    get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := ttr.Get(keys[i])
			if re == nil {
				panic(re)
			}
		})
		print("tidwall(G): get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrG.Get(keys[i])
			if !ok {
				panic(re)
			}
		})
		print("tidwall(M): get-seq        ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrM.Get(keys[i])
			if !ok {
				panic(re)
			}
		})
		if withHints {
			print("tidwall:    get-seq-hint   ")
			lotsa.Ops(N, 1, func(i, _ int) {
				re := ttr.GetHint(keys[i], &hint)
				if re == nil {
					panic(re)
				}
			})
			print("tidwall(G): get-seq-hint   ")
			lotsa.Ops(N, 1, func(i, _ int) {
				re, ok := ttrG.GetHint(keys[i], &hintG)
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
				gtr.ReplaceOrInsert(keys[i])
			})
			print("google(G):  set-rand       ")
			gtrG = newGBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				gtrG.ReplaceOrInsert(keys[i])
			})
			print("tidwall:    set-rand       ")
			ttr = newTBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.Set(keys[i])
			})
			print("tidwall(G): set-rand       ")
			ttrG = newTBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.Set(keys[i])
			})
			print("tidwall(M): set-rand       ")
			ttrM = newTBTreeM(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrM.Set(keys[i], struct{}{})
			})
			if withHints {
				print("tidwall:    set-rand-hint  ")
				ttr = newTBTree(degree)
				lotsa.Ops(N, 1, func(i, _ int) {
					ttr.SetHint(keys[i], &hint)
				})
				print("tidwall(G): set-rand-hint  ")
				ttrG = newTBTreeG(degree)
				lotsa.Ops(N, 1, func(i, _ int) {
					ttrG.SetHint(keys[i], &hintG)
				})
			}
			print("tidwall:    set-after-copy ")
			ttr = ttr.Copy()
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.Set(keys[i])
			})
			print("tidwall(G): set-after-copy ")
			ttrG = ttrG.Copy()
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.Set(keys[i])
			})
			print("tidwall:    load-rand      ")
			ttr = newTBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.Load(keys[i])
			})
			print("tidwall(G): load-rand      ")
			ttrG = newTBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.Load(keys[i])
			})
			ttrM = newTBTreeM(degree)
			print("tidwall(M): load-rand      ")
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrM.Load(keys[i], struct{}{})
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
		for _, key := range keys {
			gtr.ReplaceOrInsert(key)
			gtrG.ReplaceOrInsert(key)
			ttrG.Set(key)
			ttr.Set(key)
			ttrM.Set(key, struct{}{})
		}
		shuffleInts()

		print("google:     get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := gtr.Get(keys[i])
			if re == nil {
				panic(re)
			}
		})
		print("google(G):  get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := gtrG.Get(keys[i])
			if !ok {
				panic(re)
			}
		})
		print("tidwall:    get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := ttr.Get(keys[i])
			if re == nil {
				panic(re)
			}
		})
		print("tidwall(G): get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrG.Get(keys[i])
			if !ok {
				panic(re)
			}
		})
		print("tidwall(M): get-rand       ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrM.Get(keys[i])
			if !ok {
				panic(re)
			}
		})
		if withHints {
			print("tidwall:    get-rand-hint  ")
			lotsa.Ops(N, 1, func(i, _ int) {
				re := ttr.GetHint(keys[i], &hint)
				if re == nil {
					panic(re)
				}
			})
			print("tidwall(G): get-rand-hint  ")
			lotsa.Ops(N, 1, func(i, _ int) {
				re, ok := ttrG.GetHint(keys[i], &hintG)
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
