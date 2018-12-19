package mila

// Variables represents a group of variables
type Variables map[string]*Variable

// Variable represents a variable
type Variable struct {
	Default     string
	Type        string
	Required    bool
	Description string
	Value       string
}
