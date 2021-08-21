package main

import (
	"math/rand"
	"os"
	"sort"

	gbtree "github.com/google/btree"
	tbtree "github.com/tidwall/btree"
	"github.com/tidwall/lotsa"
)

type intT struct {
	val int
}

func (i *intT) Less(other gbtree.Item) bool {
	return i.val < other.(*intT).val
}

func main() {
	less := func(a, b interface{}) bool {
		return a.(*intT).val < b.(*intT).val
	}
	N := 1_000_000
	keys := make([]intT, N)
	for i := 0; i < N; i++ {
		keys[i] = intT{i}
	}
	lotsa.Output = os.Stdout
	lotsa.MemUsage = true

	sortInts := func() {
		sort.Slice(keys, func(i, j int) bool {
			return less(&keys[i], &keys[j])
		})
	}

	shuffleInts := func() {
		for i := range keys {
			j := rand.Intn(i + 1)
			keys[i], keys[j] = keys[j], keys[i]
		}
	}

	gMaxItems := 256

	println()
	println("** sequential set **")

	print("google:  set-seq        ")
	tr2 := gbtree.New(gMaxItems)
	sortInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr2.ReplaceOrInsert(&keys[i])
	})
	print("tidwall: set-seq        ")
	tr := tbtree.NewNonConcurrent(less)
	sortInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Set(&keys[i])
	})
	print("tidwall: set-seq-hint   ")
	tr = tbtree.NewNonConcurrent(less)
	sortInts()
	var hint tbtree.PathHint
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.SetHint(&keys[i], &hint)
	})
	print("tidwall: load-seq       ")
	tr = tbtree.NewNonConcurrent(less)
	sortInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Load(&keys[i])
	})
	print("go-arr:  append         ")
	var arr []interface{}
	sortInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		arr = append(arr, &keys[i])
	})

	println()
	println("** random set **")

	print("google:  set-rand       ")
	tr2 = gbtree.New(gMaxItems)
	shuffleInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr2.ReplaceOrInsert(&keys[i])
	})
	print("tidwall: set-rand       ")
	tr = tbtree.NewNonConcurrent(less)
	shuffleInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Set(&keys[i])
	})
	print("tidwall: set-rand-hint  ")
	tr = tbtree.NewNonConcurrent(less)
	shuffleInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.SetHint(&keys[i], &hint)
	})
	print("tidwall: set-again      ")
	shuffleInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Set(&keys[i])
	})
	print("tidwall: set-after-copy ")
	tr = tr.Copy()
	shuffleInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Set(&keys[i])
	})
	print("tidwall: load-rand      ")
	tr = tbtree.NewNonConcurrent(less)
	shuffleInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Load(&keys[i])
	})

	println()
	println("** sequential get **")

	print("google:  get-seq        ")
	sortInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr2.Get(&keys[i])
	})
	print("tidwall: get-seq        ")
	sortInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Get(&keys[i])
	})
	print("tidwall: get-seq-hint   ")
	sortInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.GetHint(&keys[i], &hint)
	})

	println()
	println("** random get **")

	print("google:  get-rand       ")
	shuffleInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr2.Get(&keys[i])
	})
	print("tidwall: get-rand       ")
	shuffleInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Get(&keys[i])
	})
	print("tidwall: get-rand-hint  ")
	shuffleInts()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.GetHint(&keys[i], &hint)
	})
}
