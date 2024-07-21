package main

import (
	"fmt"
	"log"
)

func main() {

	node, err := GetParsedFile("model/struct_model.go")
	if err != nil {
		log.Fatal(err)
	}

	structs, err := GetStructsModels(node)
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range structs {
		fmt.Printf("Struct %s\n", str.Name)

		fmt.Println("Comments:")
		for _, comment := range str.Comments {
			fmt.Printf("\tComment: %s\n", comment)
		}

		fmt.Println("Fields:")
		for _, field := range str.Fields {

			fmt.Printf("\tName: %s | Type: %s | Tag: %s\n", field.Name, field.Type, field.StructTag)
		}
	}
}
