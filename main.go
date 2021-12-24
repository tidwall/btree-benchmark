package main

import (
	"math/rand"
	"os"
	"sort"

	gbtree "github.com/google/btree"
	tbtree "github.com/tidwall/btree"
	tbtreeG "github.com/tidwall/btree/generics"
	"github.com/tidwall/lotsa"
)

type intT struct {
	val int
}

func (i intT) Less(other gbtree.Item) bool {
	return i.val < other.(intT).val
}

func lessG(a, b intT) bool {
	return a.val < b.val
}

func less(a, b interface{}) bool {
	return a.(intT).val < b.(intT).val
}

func newBTree() *tbtree.BTree {
	return tbtree.NewNonConcurrent(less)
}

func newBTreeG() *tbtreeG.Generic[intT] {
	return tbtreeG.NewGenericOptions[intT](lessG, tbtreeG.Options{NoLocks: true})
}
func newBTreeM() *tbtreeG.Map[int,struct{}] {
	return new(tbtreeG.Map[int,struct{}])
}

func main() {
	N := 1_000_000
	keys := make([]intT, N)
	for i := 0; i < N; i++ {
		keys[i] = intT{i}
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

	degree := 128

	println()
	println("** sequential set **")
	sortInts()

	// google
	print("google:     set-seq        ")
	gtr := gbtree.New(degree)
	lotsa.Ops(N, 1, func(i, _ int) {
		gtr.ReplaceOrInsert(keys[i])
	})

	// non-generics tidwall
	print("tidwall:    set-seq        ")
	ttr := newBTree()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttr.Set(keys[i])
	})
	print("tidwall(G): set-seq        ")
	ttrG := newBTreeG()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrG.Set(keys[i])
	})
	print("tidwall(M): set-seq        ")
	ttrM := newBTreeM()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrM.Set(keys[i].val, struct{}{})
	})
		
	print("tidwall:    set-seq-hint   ")
	ttr = newBTree()
	var hint tbtree.PathHint
	lotsa.Ops(N, 1, func(i, _ int) {
		ttr.SetHint(keys[i], &hint)
	})
	print("tidwall(G): set-seq-hint   ")
	ttrG = newBTreeG()
	var hintG tbtreeG.PathHint
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrG.SetHint(keys[i], &hintG)
	})
	print("tidwall:    load-seq       ")
	ttr = newBTree()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttr.Load(keys[i])
	})
	print("tidwall(G): load-seq       ")
	ttrG = newBTreeG()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrG.Load(keys[i])
	})
	print("tidwall(M): load-seq       ")
	ttrM = newBTreeM()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrM.Load(keys[i].val, struct{}{})
	})




	// go array
	print("go-arr:     append         ")
	var arr []intT
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
		re, ok := ttrM.Get(keys[i].val)
		if !ok {
			panic(re)
		}
	})
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

	println()
	println("** random set **")
	shuffleInts()

	print("google:     set-rand       ")
	gtr = gbtree.New(degree)
	lotsa.Ops(N, 1, func(i, _ int) {
		gtr.ReplaceOrInsert(keys[i])
	})
	print("tidwall:    set-rand       ")
	ttr = newBTree()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttr.Set(keys[i])
	})
	print("tidwall(G): set-rand       ")
	ttrG = newBTreeG()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrG.Set(keys[i])
	})
	print("tidwall(M): set-rand       ")
	ttrM = newBTreeM()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrM.Set(keys[i].val, struct{}{})
	})
	print("tidwall:    set-rand-hint  ")
	ttr = newBTree()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttr.SetHint(keys[i], &hint)
	})
	print("tidwall(G): set-rand-hint  ")
	ttrG = newBTreeG()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrG.SetHint(keys[i], &hintG)
	})
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
	ttr = newBTree()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttr.Load(keys[i])
	})
	print("tidwall(G): load-rand      ")
	ttrG = newBTreeG()
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrG.Load(keys[i])
	})
	ttrM = newBTreeM()
	print("tidwall(M): load-rand      ")
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrM.Load(keys[i].val, struct{}{})
	})

	println()
	println("** random get **")
	
	shuffleInts()
	gtr = gbtree.New(degree)
	ttr = newBTree()
	ttrM = newBTreeM()
	ttrG = newBTreeG()
	for _,key:=range keys{
		gtr.ReplaceOrInsert(key)
		ttrG.Set(key)
		ttr.Set(key)
		ttrM.Set(key.val, struct{}{})
	}
	shuffleInts()

	print("google:     get-rand       ")
	lotsa.Ops(N, 1, func(i, _ int) {
		re := gtr.Get(keys[i])
		if re == nil {
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
		re, ok := ttrM.Get(keys[i].val)
		if !ok {
			panic(re)
		}
	})
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

	println()
	println("** range **")
	print("google:     ascend        ")
	lotsa.Ops(N, 1, func(i, _ int) {
		if i == 0 {
			var x int
			gtr.Ascend(func(item gbtree.Item) bool {
				x += item.(intT).val
				return true
			})
		}
	})
	print("tidwall:    ascend        ")
	lotsa.Ops(N, 1, func(i, _ int) {
		if i == 0 {
			var x int
			ttr.Ascend(nil, func(item interface{}) bool {
				x += item.(intT).val
				return true
			})
		}
	})
	print("tidwall(G): iter          ")
	lotsa.Ops(N, 1, func(i, _ int) {
		if i == 0 {
			var x int
			iter := ttrG.Iter()
			for ok := iter.First(); ok; ok = iter.Next() {
				x += iter.Item().val
			}
			iter.Release()
		}
	})
	print("tidwall(G): scan          ")
	lotsa.Ops(N, 1, func(i, _ int) {
		if i == 0 {
			var x int
			ttrG.Scan(func(item intT) bool {
				x += item.val
				return true
			})
		}
	})
	print("tidwall(G): walk          ")
	lotsa.Ops(N, 1, func(i, _ int) {
		if i == 0 {
			var x int
			ttrG.Walk(func(items []intT) bool {
				for j := 0; j < len(items); j++ {
					x += items[j].val
				}
				return true
			})
		}
	})

	print("go-arr:     for-loop      ")
	lotsa.Ops(N, 1, func(i, _ int) {
		if i == 0 {
			var x int
			for j := 0; j < len(arr); j++ {
				x += arr[j].val
			}
		}
	})

}
