package model

// description of structure method
type StructMethodModel struct {
	Name        string // method name
	ReturnTypes []string
	Arguments   []StructMethodArgument
}

// method argument
type StructMethodArgument struct {
	Name string
	Type string
}
