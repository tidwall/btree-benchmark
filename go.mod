module github.com/tidwall/btree-benchmark

go 1.18

require (
	github.com/google/btree v1.0.1
	github.com/tidwall/btree v0.7.1
	github.com/tidwall/lotsa v1.0.2
)

require github.com/tidwall/btree/generics v0.0.0-00010101000000-000000000000

// replace github.com/tidwall/btree/generics => github.com/tidwall/btree generics
replace github.com/tidwall/btree/generics => github.com/tidwall/btree v0.7.2-0.20211210225549-ac22dbfec23d
