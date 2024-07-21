package main

import (
	"fmt"
	"log"
)

func main() {

	node, err := GetParsedFile("/home/nikita/Downloads/golang_web_services_2024-04-26/5/99_hw/codegen/api.go")
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

		fmt.Println("Methods:")
		for _, method := range str.Methods {

			fmt.Printf("\tName: %s | Return types: %v | Accepts: %v\n", method.Name, method.Returnings, method.Arguments)
			fmt.Printf("\tComments: %v\n", method.Comments)
		}

		fmt.Println()
	}
}
