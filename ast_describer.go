package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/banderveloper/go-ast-describer/model"
)

func GetParsedFile(filePath string) (*ast.File, error) {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func GetStructsModels(fileNode *ast.File) ([]model.StructModel, error) {

	if fileNode == nil {
		return nil, errors.New("file node is nil")
	}

	result := make([]model.StructModel, 0)

	// iterate over file declarations
	for _, decl := range fileNode.Decls {

		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		// Iterate over file specs (types)
		for _, spec := range genDecl.Specs {
			// if type
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// if type is struct
			currStruct, ok := currType.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// Create current structure model
			curStructModel := model.StructModel{}
			// Fill structure name
			curStructModel.Name = currType.Name.Name

			if genDecl.Doc != nil {
				// Fill structure comments
				for _, comment := range genDecl.Doc.List {
					trimmedComment := strings.TrimSpace(strings.TrimLeft(comment.Text, "//"))
					curStructModel.Comments = append(curStructModel.Comments, trimmedComment)
				}
			}

			// Iterave over structure fields
			for i, field := range currStruct.Fields.List {

				curStructModel.Fields = append(curStructModel.Fields, model.StructFieldModel{
					Name: field.Names[0].Name,
					Type: getTypeName(field.Type),
				})

				if field.Tag != nil {
					curStructModel.Fields[i].StructTag = strings.Trim(field.Tag.Value, "`")
				}
			}

			// Add filled structure to result slice
			result = append(result, curStructModel)
		}
	}

	return result, nil
}

// getTypeName returns the name of the given AST expression as a string
func getTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return getTypeName(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + getTypeName(t.Elt)
	case *ast.StarExpr:
		return "*" + getTypeName(t.X)
	case *ast.MapType:
		return "map[" + getTypeName(t.Key) + "]" + getTypeName(t.Value)
	case *ast.ChanType:
		return "chan " + getTypeName(t.Value)
	default:
		return "unknown"
	}
}
