package lca

// Tarjan represents Tarjan algorithm.
type Tarjan struct {
	parent   map[Key]Key
	rank     map[Key]int
	ancestor map[Key]Key
	colored  map[Key]Key

	lcaMatrix map[Key]map[Key]Key
}

// NewTarjan release Tarjan's LCA algorithm.
func NewTarjan(d *Node) *Tarjan {
	t := &Tarjan{
		parent:    make(map[Key]Key),
		rank:      make(map[Key]int),
		ancestor:  make(map[Key]Key),
		colored:   make(map[Key]Key),
		lcaMatrix: make(map[Key]map[Key]Key),
	}

	t.precalculateLCA(d.Key, d)
	return t
}

// Find realises Finder interface using Tarjan algorithm
// with path compression and union by rank heuristics.
// The algorithm uses preprocessing during NewTarjan method,
// so at this moment its complexity is ~O(1).
func (tar *Tarjan) Find(a, b Key) *Key {
	if a > b {
		a, b = b, a
	}

	if val, ok := tar.lcaMatrix[a][b]; ok {
		return &val
	}

	return nil
}

// precalculateLCA finds LCA using path compression and union by rank heuristics.
func (tar *Tarjan) precalculateLCA(manager Key, u *Node) {
	tar.makeSet(u.Key)
	tar.ancestor[tar.findSet(u.Key)] = u.Key

	for _, v := range u.Subnodes {
		tar.precalculateLCA(u.Key, &v)

		tar.union(u.Key, v.Key)
		tar.ancestor[tar.findSet(u.Key)] = u.Key
	}

	tar.colored[u.Key] = manager

	// If u is one of two people (A and B) we are looking for,
	// let's see if we already "colored" another person (key) from this couple.
	// If so, LCA between A and B is the ancestor of key's set
	for key, keyManager := range tar.colored {
		lca := tar.ancestor[tar.findSet(key)]

		emp1 := u.Key
		emp2 := key
		if u.Key > key {
			emp1, emp2 = emp2, emp1
		}

		if len(tar.lcaMatrix[emp1]) == 0 {
			tar.lcaMatrix[emp1] = make(map[Key]Key)
		}

		// In our case LCA mustn't be a person itself,
		// so if LCA between two employees is one of them,
		// we need choose a manager of this employee
		if lca == u.Key {
			lca = manager
		} else if lca == key {
			lca = keyManager
		}

		tar.lcaMatrix[emp1][emp2] = lca
	}
}

// makeSet creates a tree with one node.
// It makes a new set from an element
// (by default an element represent its own set).
func (tar *Tarjan) makeSet(x Key) {
	tar.parent[x] = x
	tar.rank[x] = 0
}

// findSet follows parent links in the tree until the root is reached.
// Find set is a "path compression" - it follows the chain of element's parents
// until is reached an elements whose parent is itself.
func (tar *Tarjan) findSet(x Key) Key {
	if x != tar.parent[x] {
		tar.parent[x] = tar.findSet(tar.parent[x])
	}
	return tar.parent[x]
}

// link makes a link from a root from one element to another.
// An element with a bigger rank becomes a parent.
func (tar *Tarjan) link(x, y Key) {
	if tar.rank[x] > tar.rank[y] {
		tar.parent[y] = x
	} else {
		tar.parent[x] = y

		if tar.rank[x] == tar.rank[y] {
			tar.rank[y]++
		}
	}
}

// union create a union (by rank) of trees contain x and y elements.
func (tar *Tarjan) union(x, y Key) {
	tar.link(tar.findSet(x), tar.findSet(y))
}
