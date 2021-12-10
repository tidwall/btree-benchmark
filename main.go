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

func newBTreeG() *tbtreeG.BTree[intT] {
	return tbtreeG.NewOptions[intT](lessG, tbtreeG.Options{NoLocks: true})
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

	// go array
	print("go-arr:     append         ")
	var arr []interface{}
	lotsa.Ops(N, 1, func(i, _ int) {
		arr = append(arr, keys[i])
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
	print("tidwall:    set-again      ")
	lotsa.Ops(N, 1, func(i, _ int) {
		ttr.Set(keys[i])
	})
	print("tidwall(G): set-again      ")
	lotsa.Ops(N, 1, func(i, _ int) {
		ttrG.Set(keys[i])
	})
	print("tidwall(:   set-after-copy ")
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
	println("** random get **")
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
