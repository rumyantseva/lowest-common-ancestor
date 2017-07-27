package lca

// Key is a key to identify nodes.
type Key string

// Node represents a node of a tree.
// Key of node must be unique for identification.
type Node struct {
	Key      Key    `json:"name"`
	Subnodes []Node `json:"employees"`
}

// Finder calculates the lowest common ancestor.
type Finder interface {

	// Find returns the lowest common ancestor between a and b.
	// If a or b is root, the lowest common ancestor is root.
	// If one of a or b is an ancestor of another, the lowest common ancestor is the next their ancestor.
	// So, only root might be an ancestor of itself.
	Find(a, b Key) *Key
}
