package lca

import "testing"

func TestMakeSet(t *testing.T) {
	var key Key
	key = "some key"
	td := prepareEmptyData()
	td.makeSet(key)

	if td.rank[key] != 0 {
		t.Fail()
	}

	if td.parent[key] != key {
		t.Fail()
	}
}

func TestFindSet(t *testing.T) {
	td := prepareEmptyData()

	var key1 Key
	key1 = "some key 1"
	td.makeSet(key1)

	var key2 Key
	key2 = "some key 2"
	td.makeSet(key2)

	// Make key1 as child and key2 as parent
	td.parent[key1] = key2
	td.rank[key2]++

	setOfKey1 := td.findSet(key1)
	if setOfKey1 != key2 {
		t.Fail()
	}

	setOfKey2 := td.findSet(key2)
	if setOfKey2 != key2 {
		t.Fail()
	}
}

func TestLink(t *testing.T) {
	td := prepareEmptyData()

	var key1 Key
	key1 = "some key 1"
	td.makeSet(key1)

	var key2 Key
	key2 = "some key 2"
	td.makeSet(key2)

	// Make key1 as child and key2 as parent
	td.link(key1, key2)

	setOfKey1 := td.findSet(key1)
	if setOfKey1 != key2 {
		t.Fail()
	}

	setOfKey2 := td.findSet(key2)
	if setOfKey2 != key2 {
		t.Fail()
	}
}

func TestUnion(t *testing.T) {
	td := prepareEmptyData()

	var key1 Key
	key1 = "some key 1"
	td.makeSet(key1)

	var key2 Key
	key2 = "some key 2"
	td.makeSet(key2)

	// Make key1 as child and key2 as parent
	td.union(key1, key2)

	setOfKey1 := td.findSet(key1)
	if setOfKey1 != key2 {
		t.Fail()
	}

	setOfKey2 := td.findSet(key2)
	if setOfKey2 != key2 {
		t.Fail()
	}
}

func prepareEmptyData() *Tarjan {
	return &Tarjan{
		parent:   make(map[Key]Key),
		rank:     make(map[Key]int),
		ancestor: make(map[Key]Key),
		colored:  make(map[Key]Key),
	}
}
