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

func print_label(label, action string) {
	fmt.Printf("%-11s %-17s ", label+":", action)
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
	withPivot := true
	withScan := true
	withHints := true

	fmt.Printf("\ndegree=%d, key=string (16 bytes), val=int64, count=%d\n",
		degree, N)

	var hint tbtree.PathHint
	var hintG tbtree.PathHint

	if withSeq {
		println()
		println("** sequential set **")
		sortInts()

		// google
		print_label("google", "set-seq")
		gtr = newGBTree(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			gtr.ReplaceOrInsert(items[i])
		})

		print_label("google(G)", "set-seq")
		gtrG = newGBTreeG(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			gtrG.ReplaceOrInsert(items[i])
		})

		// non-generics tidwall
		print_label("tidwall", "set-seq")
		ttr = newTBTree(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttr.Set(items[i])
		})
		print_label("tidwall(G)", "set-seq")
		ttrG = newTBTreeG(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrG.Set(items[i])
		})
		print_label("tidwall(M)", "set-seq")
		ttrM = newTBTreeM(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrM.Set(items[i].key, items[i].val)
		})

		if withHints {
			print_label("tidwall", "set-seq-hint")
			ttr = newTBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.SetHint(items[i], &hint)
			})
			print_label("tidwall(G)", "set-seq-hint")
			ttrG = newTBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.SetHint(items[i], &hintG)
			})
		}
		print_label("tidwall", "load-seq")
		ttr = newTBTree(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttr.Load(items[i])
		})
		print_label("tidwall(G)", "load-seq")
		ttrG = newTBTreeG(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrG.Load(items[i])
		})
		print_label("tidwall(M)", "load-seq")
		ttrM = newTBTreeM(degree)
		lotsa.Ops(N, 1, func(i, _ int) {
			ttrM.Load(items[i].key, items[i].val)
		})

		println()
		println("** sequential get **")
		sortInts()

		print_label("google", "get-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := gtr.Get(items[i])
			if re == nil {
				panic(re)
			}
		})
		print_label("google(G)", "get-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := gtrG.Get(items[i])
			if !ok {
				panic(re)
			}
		})
		print_label("tidwall", "get-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := ttr.Get(items[i])
			if re == nil {
				panic(re)
			}
		})
		print_label("tidwall(G)", "get-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrG.Get(items[i])
			if !ok {
				panic(re)
			}
		})
		print_label("tidwall(M)", "get-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrM.Get(items[i].key)
			if !ok {
				panic(re)
			}
		})
		if withHints {
			print_label("tidwall", "get-seq-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				re := ttr.GetHint(items[i], &hint)
				if re == nil {
					panic(re)
				}
			})
			print_label("tidwall(G)", "get-seq-hint")
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
			print_label("google", "set-rand")
			gtr = newGBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				gtr.ReplaceOrInsert(items[i])
			})
			print_label("google(G)", "set-rand")
			gtrG = newGBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				gtrG.ReplaceOrInsert(items[i])
			})
			print_label("tidwall", "set-rand")
			ttr = newTBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.Set(items[i])
			})
			print_label("tidwall(G)", "set-rand")
			ttrG = newTBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.Set(items[i])
			})
			print_label("tidwall(M)", "set-rand")
			ttrM = newTBTreeM(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrM.Set(items[i].key, items[i].val)
			})
			if withHints {
				print_label("tidwall", "set-rand-hint")
				ttr = newTBTree(degree)
				lotsa.Ops(N, 1, func(i, _ int) {
					ttr.SetHint(items[i], &hint)
				})
				print_label("tidwall(G)", "set-rand-hint")
				ttrG = newTBTreeG(degree)
				lotsa.Ops(N, 1, func(i, _ int) {
					ttrG.SetHint(items[i], &hintG)
				})
			}
			print_label("tidwall", "set-after-copy")
			ttr = ttr.Copy()
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.Set(items[i])
			})
			print_label("tidwall(G)", "set-after-copy")
			ttrG = ttrG.Copy()
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.Set(items[i])
			})
			print_label("tidwall", "load-rand")
			ttr = newTBTree(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttr.Load(items[i])
			})
			print_label("tidwall(G)", "load-rand")
			ttrG = newTBTreeG(degree)
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrG.Load(items[i])
			})
			ttrM = newTBTreeM(degree)
			print_label("tidwall(M)", "load-rand")
			lotsa.Ops(N, 1, func(i, _ int) {
				ttrM.Load(items[i].key, items[i].val)
			})

			map_ := make(map[keyT]valT)
			print("go-map:     set            ")
			lotsa.Ops(N, 1, func(i, _ int) {
				map_[items[i].key] = items[i].val
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
		map_ := make(map[keyT]valT)
		for _, item := range items {
			gtr.ReplaceOrInsert(item)
			gtrG.ReplaceOrInsert(item)
			ttrG.Set(item)
			ttr.Set(item)
			ttrM.Set(item.key, item.val)
			map_[item.key] = item.val
		}
		shuffleInts()

		print_label("google", "get-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := gtr.Get(items[i])
			if re == nil {
				panic(re)
			}
		})
		print_label("google(G)", "get-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := gtrG.Get(items[i])
			if !ok {
				panic(re)
			}
		})
		print_label("tidwall", "get-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			re := ttr.Get(items[i])
			if re == nil {
				panic(re)
			}
		})
		print_label("tidwall(G)", "get-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrG.Get(items[i])
			if !ok {
				panic(re)
			}
		})
		print_label("tidwall(M)", "get-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := ttrM.Get(items[i].key)
			if !ok {
				panic(re)
			}
		})
		if withHints {
			print_label("tidwall", "get-rand-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				re := ttr.GetHint(items[i], &hint)
				if re == nil {
					panic(re)
				}
			})
			print_label("tidwall(G)", "get-rand-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				re, ok := ttrG.GetHint(items[i], &hintG)
				if !ok {
					panic(re)
				}
			})
		}
		print("go-map:     get            ")
		lotsa.Ops(N, 1, func(i, _ int) {
			re, ok := map_[items[i].key]
			if !ok {
				panic(re)
			}
		})
	}

	if !withRand {
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
	}

	if withPivot {
		sortInts()
		const M = 10
		var hint tbtree.PathHint
		println()
		fmt.Printf("** sequential pivot **\n")
		fmt.Printf("Test getting %d consecutive items starting at a pivot.\n", M)
		print_label("google", "ascend-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			gtr.AscendGreaterOrEqual(items[i], func(item gbtree.Item) bool {
				count++
				return count < M
			})
		})
		print_label("google", "descend-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			gtr.DescendLessOrEqual(items[i], func(item gbtree.Item) bool {
				count++
				return count < M
			})
		})
		print_label("google(G)", "ascend-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			gtrG.AscendGreaterOrEqual(items[i], func(item itemT) bool {
				count++
				return count < M
			})
		})
		print_label("google(G)", "descend-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			gtrG.DescendLessOrEqual(items[i], func(item itemT) bool {
				count++
				return count < M
			})
		})
		print_label("tidwall", "ascend-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			ttr.Ascend(items[i], func(item any) bool {
				count++
				return count < M
			})
		})
		print_label("tidwall", "descend-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			ttr.Ascend(items[i], func(item any) bool {
				count++
				return count < M
			})
		})
		if withHints {
			print_label("tidwall", "ascend-seq-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				var count int
				ttr.AscendHint(items[i], func(item any) bool {
					count++
					return count < M
				}, &hint)
			})
			print_label("tidwall", "descend-seq-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				var count int
				ttr.DescendHint(items[i], func(item any) bool {
					count++
					return count < M
				}, &hint)
			})
		}
		print_label("tidwall(G)", "ascend-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			ttrG.Ascend(items[i], func(item itemT) bool {
				count++
				return count < M
			})
		})
		print_label("tidwall(G)", "descend-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			ttrG.Descend(items[i], func(item itemT) bool {
				count++
				return count < M
			})
		})
		if withHints {
			print_label("tidwall(G)", "ascend-seq-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				var count int
				ttrG.AscendHint(items[i], func(item itemT) bool {
					count++
					return count < M
				}, &hint)
			})
			print_label("tidwall(G)", "descend-seq-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				var count int
				ttrG.DescendHint(items[i], func(item itemT) bool {
					count++
					return count < M
				}, &hint)
			})
		}
		print_label("tidwall(G)", "iter-seq")
		lotsa.Ops(N, 1, func(i, _ int) {
			iter := ttrG.Iter()
			var count int
			for ok := iter.Seek(items[i]); ok; ok = iter.Next() {
				count++
				if count == M {
					break
				}
			}
			iter.Release()
		})
		print_label("tidwall(G)", "iter-seq-hint")
		lotsa.Ops(N, 1, func(i, _ int) {
			iter := ttrG.Iter()
			var count int
			for ok := iter.SeekHint(items[i], &hint); ok; ok = iter.Next() {
				count++
				if count == M {
					break
				}
			}
			iter.Release()
		})
	}

	if withPivot {
		shuffleInts()
		const M = 10
		var hint tbtree.PathHint
		println()
		fmt.Printf("** random pivot **\n")
		fmt.Printf("Test getting %d consecutive items starting at a pivot.\n", M)
		print_label("google", "ascend-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			gtr.AscendGreaterOrEqual(items[i], func(item gbtree.Item) bool {
				count++
				return count < M
			})
		})
		print_label("google", "descend-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			gtr.DescendLessOrEqual(items[i], func(item gbtree.Item) bool {
				count++
				return count < M
			})
		})
		print_label("google(G)", "ascend-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			gtrG.AscendGreaterOrEqual(items[i], func(item itemT) bool {
				count++
				return count < M
			})
		})
		print_label("google(G)", "descend-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			gtrG.DescendLessOrEqual(items[i], func(item itemT) bool {
				count++
				return count < M
			})
		})
		print_label("tidwall", "ascend-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			ttr.Ascend(items[i], func(item any) bool {
				count++
				return count < M
			})
		})
		print_label("tidwall", "descend-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			ttr.Ascend(items[i], func(item any) bool {
				count++
				return count < M
			})
		})
		if withHints {
			print_label("tidwall", "ascend-rand-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				var count int
				ttr.AscendHint(items[i], func(item any) bool {
					count++
					return count < M
				}, &hint)
			})
			print_label("tidwall", "descend-rand-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				var count int
				ttr.DescendHint(items[i], func(item any) bool {
					count++
					return count < M
				}, &hint)
			})
		}
		print_label("tidwall(G)", "ascend-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			ttrG.Ascend(items[i], func(item itemT) bool {
				count++
				return count < M
			})
		})
		print_label("tidwall(G)", "descend-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			var count int
			ttrG.Descend(items[i], func(item itemT) bool {
				count++
				return count < M
			})
		})
		if withHints {
			print_label("tidwall(G)", "ascend-rand-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				var count int
				ttrG.AscendHint(items[i], func(item itemT) bool {
					count++
					return count < M
				}, &hint)
			})
			print_label("tidwall(G)", "descend-rand-hint")
			lotsa.Ops(N, 1, func(i, _ int) {
				var count int
				ttrG.DescendHint(items[i], func(item itemT) bool {
					count++
					return count < M
				}, &hint)
			})
		}
		print_label("tidwall(G)", "iter-rand")
		lotsa.Ops(N, 1, func(i, _ int) {
			iter := ttrG.Iter()
			var count int
			for ok := iter.Seek(items[i]); ok; ok = iter.Next() {
				count++
				if count == M {
					break
				}
			}
			iter.Release()
		})
		print_label("tidwall(G)", "iter-rand-hint")
		lotsa.Ops(N, 1, func(i, _ int) {
			iter := ttrG.Iter()
			var count int
			for ok := iter.SeekHint(items[i], &hint); ok; ok = iter.Next() {
				count++
				if count == M {
					break
				}
			}
			iter.Release()
		})
	}

	if withScan {
		println()
		println("** scan **")
		println("Test scanning over every item in the tree")
		print_label("google", "ascend")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				gtr.Ascend(func(item gbtree.Item) bool {
					return true
				})
			}
		})
		print_label("google(G)", "ascend")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				gtrG.Ascend(func(item itemT) bool {
					return true
				})
			}
		})
		print_label("tidwall", "ascend")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				ttr.Ascend(nil, func(item interface{}) bool {
					return true
				})
			}
		})
		print_label("tidwall(G)", "scan")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				ttrG.Scan(func(item itemT) bool {
					return true
				})
			}
		})
		print_label("tidwall(G)", "walk")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				ttrG.Walk(func(items []itemT) bool {
					for j := 0; j < len(items); j++ {

					}
					return true
				})
			}
		})
		print_label("tidwall(G)", "iter")
		lotsa.Ops(N, 1, func(i, _ int) {
			if i == 0 {
				iter := ttrG.Iter()
				for ok := iter.First(); ok; ok = iter.Next() {
				}
				iter.Release()
			}
		})
	}
}
