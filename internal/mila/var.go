package mila

// Vars represents a group of vars
type Vars map[string]*Var

// Var represents a var
type Var struct {
	Default     string
	Type        string
	Required    bool
	Description string
	Value       string
}
