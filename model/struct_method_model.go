package model

// description of structure method
type StructMethodModel struct {
	Name       string // method name
	Returnings []StructMethodArg
	Arguments  []StructMethodArg
}

// method acceping/return argument
type StructMethodArg struct {
	Index int
	Name  string
	Type  string
}
