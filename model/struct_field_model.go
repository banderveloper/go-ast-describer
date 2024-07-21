package model

// description of field of structure
type StructFieldModel struct {
	Name      string // field name
	Type      string // field type
	StructTag string // field struct tag without ``
}

// description of field's struct tag KV
type StructTagKV struct {
	Key   string
	Value string
}
