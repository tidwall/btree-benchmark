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

func (i intT) Less(other gbtree.Item) bool {
	return i.val < other.(intT).val
}

func main() {
	less := func(a, b interface{}) bool {
		return a.(intT).val < b.(intT).val
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

	print("google:  set-seq        ")
	tr2 := gbtree.New(degree)
	lotsa.Ops(N, 1, func(i, _ int) {
		tr2.ReplaceOrInsert(keys[i])
	})
	print("tidwall: set-seq        ")
	tr := tbtree.NewNonConcurrent(less)
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Set(keys[i])
	})
	print("tidwall: set-seq-hint   ")
	tr = tbtree.NewNonConcurrent(less)
	var hint tbtree.PathHint
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.SetHint(keys[i], &hint)
	})
	print("tidwall: load-seq       ")
	tr = tbtree.NewNonConcurrent(less)
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Load(keys[i])
	})
	print("go-arr:  append         ")
	var arr []interface{}
	lotsa.Ops(N, 1, func(i, _ int) {
		arr = append(arr, keys[i])
	})

	println()
	println("** random set **")
	shuffleInts()

	print("google:  set-rand       ")
	tr2 = gbtree.New(degree)
	lotsa.Ops(N, 1, func(i, _ int) {
		tr2.ReplaceOrInsert(keys[i])
	})
	print("tidwall: set-rand       ")
	tr = tbtree.NewNonConcurrent(less)
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Set(keys[i])
	})
	print("tidwall: set-rand-hint  ")
	tr = tbtree.NewNonConcurrent(less)
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.SetHint(keys[i], &hint)
	})
	print("tidwall: set-again      ")
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Set(keys[i])
	})
	print("tidwall: set-after-copy ")
	tr = tr.Copy()
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Set(keys[i])
	})
	print("tidwall: load-rand      ")
	tr = tbtree.NewNonConcurrent(less)
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Load(keys[i])
	})

	println()
	println("** sequential get **")
	sortInts()

	print("google:  get-seq        ")
	lotsa.Ops(N, 1, func(i, _ int) {
		re := tr2.Get(keys[i])
		if re == nil {
			panic(re)
		}
	})
	print("tidwall: get-seq        ")
	lotsa.Ops(N, 1, func(i, _ int) {
		re := tr.Get(keys[i])
		if re == nil {
			panic(re)
		}
	})
	print("tidwall: get-seq-hint   ")
	lotsa.Ops(N, 1, func(i, _ int) {
		re := tr.GetHint(keys[i], &hint)
		if re == nil {
			panic(re)
		}
	})

	println()
	println("** random get **")
	shuffleInts()

	print("google:  get-rand       ")
	lotsa.Ops(N, 1, func(i, _ int) {
		re := tr2.Get(keys[i])
		if re == nil {
			panic(re)
		}
	})
	print("tidwall: get-rand       ")
	lotsa.Ops(N, 1, func(i, _ int) {
		re := tr.Get(keys[i])
		if re == nil {
			panic(re)
		}
	})
	print("tidwall: get-rand-hint  ")
	lotsa.Ops(N, 1, func(i, _ int) {
		re := tr.GetHint(keys[i], &hint)
		if re == nil {
			panic(re)
		}
	})

	println()
	println("** sequential delete **")
	sortInts()

	print("google:  del-seq        ")
	lotsa.Ops(N, 1, func(i, _ int) {
		tr2.Delete(keys[i])
	})
	print("tidwall: del-seq        ")
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Delete(keys[i])
	})
	print("tidwall: del-seq-hint   ")
	tr = tbtree.NewNonConcurrent(less)
	for i := range keys {
		tr.Set(keys[i])
	}
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.DeleteHint(keys[i], &hint)
	})

	tr2 = gbtree.New(degree)
	for i := range keys {
		tr2.ReplaceOrInsert(keys[i])
	}
	tr = tbtree.NewNonConcurrent(less)
	for i := range keys {
		tr.Set(keys[i])
	}

	println()
	println("** random delete **")
	shuffleInts()

	print("google:  del-rand        ")
	lotsa.Ops(N, 1, func(i, _ int) {
		tr2.Delete(keys[i])
	})
	print("tidwall: del-rand        ")
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.Delete(keys[i])
	})
	print("tidwall: del-rand-hint   ")
	tr = tbtree.NewNonConcurrent(less)
	for i := range keys {
		tr.Set(keys[i])
	}
	lotsa.Ops(N, 1, func(i, _ int) {
		tr.DeleteHint(keys[i], &hint)
	})
	tr2 = gbtree.New(degree)
	for i := range keys {
		tr2.ReplaceOrInsert(keys[i])
	}
	tr = tbtree.NewNonConcurrent(less)
	for i := range keys {
		tr.Set(keys[i])
	}

}
