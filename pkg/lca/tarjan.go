package lca

// TarjanData represents necessary structures to store data.
type TarjanData struct {
	Parent   map[Key]Key
	Rank     map[Key]int
	Ancestor map[Key]Key
	Colored  map[Key]bool

	LCA Key
	A   Key
	B   Key
}

// Tarjan release Tarjan's LCA algorithm.
func Tarjan(d *Directory, a, b Key) Key {
	t := &TarjanData{
		Parent:   make(map[Key]Key),
		Rank:     make(map[Key]int),
		Ancestor: make(map[Key]Key),
		Colored:  make(map[Key]bool),
		A:        a,
		B:        b,
	}

	t.calculateLCA(d)
	return t.LCA
}

// calculateLCA finds LCA using path compression and union by rank heuristics.
func (td *TarjanData) calculateLCA(u *Directory) bool {
	td.makeSet(u.Name)
	td.Ancestor[td.findSet(u.Name)] = u.Name

	for _, v := range u.Employees {
		if td.calculateLCA(&v) {
			return true
		}

		td.union(u.Name, v.Name)
		td.Ancestor[td.findSet(u.Name)] = u.Name
	}

	td.Colored[u.Name] = true

	// If u is one of two people (A and B) we are looking for,
	// let's see if we already "colored" another person (k) from this couple.
	// If so, LCA between A and B is the ancestor of k's set
	var k Key
	if u.Name == td.A {
		k = td.B
	} else if u.Name == td.B {
		k = td.A
	}

	if len(k) > 0 && td.Colored[k] {
		td.LCA = td.Ancestor[td.findSet(k)]
		return true
	}

	return false
}

// makeSet creates a tree with one node.
// It makes a new set from an element
// (by default an element represent its own set).
func (td *TarjanData) makeSet(x Key) {
	td.Parent[x] = x
	td.Rank[x] = 0
}

// findSet follows parent links in the tree until the root is reached.
// Find set is a "path compression" - it follows the chain of element's parents
// until is reached an elements whose parent is itself.
func (td *TarjanData) findSet(x Key) Key {
	if x != td.Parent[x] {
		td.Parent[x] = td.findSet(td.Parent[x])
	}
	return td.Parent[x]
}

// link makes a link from a root from one element to another.
// An element with a bigger rank becomes a parent.
func (td *TarjanData) link(x Key, y Key) {
	if td.Rank[x] > td.Rank[y] {
		td.Parent[y] = x
	} else {
		td.Parent[x] = y

		if td.Rank[x] == td.Rank[y] {
			td.Rank[y]++
		}
	}
}

// union create a union (by rank) of trees contain x and y elements.
func (td *TarjanData) union(x, y Key) {
	td.link(td.findSet(x), td.findSet(y))
}
