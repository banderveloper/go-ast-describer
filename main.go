package main

import (
	"fmt"

	"github.com/banderveloper/go-ast-describer/model"
)

func main() {

	// fmt.Println("Hello world")

	m := model.StructFieldModel{
		Name:      "Name",
		Type:      "string",
		StructTag: `json:"hello" database:"idk" emma:"what"`,
	}

	tags := m.GetTags()
	fmt.Println(tags["database"] == "idk")

}
