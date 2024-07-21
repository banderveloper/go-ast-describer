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

			curStructModel.Name = currType.Name.Name
			curStructModel.Comments = getStructComments(genDecl.Doc)
			curStructModel.Fields = getStructFields(currStruct)
			curStructModel.Methods = getStructMethods(curStructModel.Name, fileNode)

			// Add filled structure to result slice
			result = append(result, curStructModel)
		}
	}

	return result, nil
}

func getStructComments(typeDoc *ast.CommentGroup) []string {

	result := make([]string, 0)

	if typeDoc != nil {
		for _, comment := range typeDoc.List {
			trimmedComment := strings.TrimSpace(strings.TrimLeft(comment.Text, "//"))
			result = append(result, trimmedComment)
		}
	}

	return result
}

func getStructFields(str *ast.StructType) []model.StructFieldModel {

	result := make([]model.StructFieldModel, 0)

	// Iterave over structure fields
	for i, field := range str.Fields.List {

		result = append(result, model.StructFieldModel{
			Name: field.Names[0].Name,
			Type: getTypeName(field.Type),
		})

		if field.Tag != nil {
			result[i].StructTag = strings.Trim(field.Tag.Value, "`")
		}
	}

	return result
}

// getStructMethods returns the methods of the given struct in the desired format.
func getStructMethods(structName string, node ast.Node) []model.StructMethodModel {
	var methods []model.StructMethodModel

	// Inspect the AST to find the struct and its methods.
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for function declarations.
		if fn, ok := n.(*ast.FuncDecl); ok {
			// Check if the function has a receiver and the receiver is of the specified struct type.
			if fn.Recv != nil && len(fn.Recv.List) > 0 {
				if starExpr, ok := fn.Recv.List[0].Type.(*ast.StarExpr); ok {
					if ident, ok := starExpr.X.(*ast.Ident); ok && ident.Name == structName {
						method := model.StructMethodModel{
							Name:        fn.Name.Name,
							ReturnTypes: getReturnTypes(fn.Type.Results),
							Arguments:   getArguments(fn.Type.Params),
						}
						methods = append(methods, method)
					}
				}
			}
		}
		return true
	})

	return methods
}

// getReturnTypes extracts the return types from the function result fields.
func getReturnTypes(results *ast.FieldList) []string {
	var returnTypes []string
	if results != nil {
		for _, result := range results.List {
			returnTypes = append(returnTypes, getTypeName(result.Type))
		}
	}
	return returnTypes
}

// getArguments extracts the arguments from the function parameter fields.
func getArguments(params *ast.FieldList) []model.StructMethodArgument {
	var arguments []model.StructMethodArgument
	if params != nil {
		for _, param := range params.List {
			for _, name := range param.Names {
				argument := model.StructMethodArgument{
					Name: name.Name,
					Type: getTypeName(param.Type),
				}
				arguments = append(arguments, argument)
			}
		}
	}
	return arguments
}

// getTypeName returns the name of the given AST expression as a string
func getTypeName(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.StarExpr:
		return "*" + getTypeName(v.X)
	case *ast.SelectorExpr:
		return getTypeName(v.X) + "." + v.Sel.Name
	case *ast.ArrayType:
		return "[]" + getTypeName(v.Elt)
	case *ast.MapType:
		return "map[" + getTypeName(v.Key) + "]" + getTypeName(v.Value)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return "unknown"
	}
}
