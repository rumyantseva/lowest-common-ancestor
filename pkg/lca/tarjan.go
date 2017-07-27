package lca

// TarjanData represents necessary structures to store data.
type TarjanData struct {
	Parent   map[Key]Key
	Rank     map[Key]int
	Ancestor map[Key]Key
	Colored  map[Key]bool

	Total map[Key]map[Key]Key

	LCA Key
}

// Tarjan release Tarjan's LCA algorithm.
func Tarjan(d *Directory) map[Key]map[Key]Key {
	t := &TarjanData{
		Parent:   make(map[Key]Key),
		Rank:     make(map[Key]int),
		Ancestor: make(map[Key]Key),
		Colored:  make(map[Key]bool),
		Total:    make(map[Key]map[Key]Key),
	}

	t.calculateLCA(d)
	return t.Total
}

// calculateLCA finds LCA using path compression and union by rank heuristics.
func (td *TarjanData) calculateLCA(u *Directory) {
	td.makeSet(u.Name)
	td.Ancestor[td.findSet(u.Name)] = u.Name

	for _, v := range u.Employees {
		td.calculateLCA(&v)

		td.union(u.Name, v.Name)
		td.Ancestor[td.findSet(u.Name)] = u.Name
	}

	td.Colored[u.Name] = true

	// If u is one of two people (A and B) we are looking for,
	// let's see if we already "colored" another person (key) from this couple.
	// If so, LCA between A and B is the ancestor of key's set
	for key := range td.Colored {
		lca := td.Ancestor[td.findSet(key)]
		var emp1, emp2 Key
		if u.Name <= key {
			emp1 = u.Name
			emp2 = key
		} else {
			emp1 = key
			emp2 = u.Name
		}

		if len(td.Total[emp1]) == 0 {
			td.Total[emp1] = make(map[Key]Key)
		}

		// In our case LCA mustn't be a person itself,
		// so if LCA between two employees is one of them,
		// we need choose a manager of this employee
		/*if lca == emp1 {
			lca = TODO
		} else if lca == emp2 {
			lca = TODO
		}*/

		td.Total[emp1][emp2] = lca
	}
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
func (td *TarjanData) link(x, y Key) {
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
