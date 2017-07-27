package lca

type Key string

type Directory struct {
	Name      Key         `json:"name"`
	Employees []Directory `json:"employees"`
}
