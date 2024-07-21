package model

import "strings"

// description of structure method
type StructMethodModel struct {
	Name       string // method name
	Returnings []StructMethodArg
	Arguments  []StructMethodArg
	Comments   []string
}

// method acceping/return argument
type StructMethodArg struct {
	Index int
	Name  string
	Type  string
}

// get all structure comments with given prefix, including it
func (sm *StructMethodModel) GetCommentsWithPrefix(prefix string) []string {

	if prefix == "" {
		return sm.Comments
	}

	resultComments := make([]string, 0)

	for _, comment := range sm.Comments {
		if strings.HasPrefix(comment, prefix) {
			resultComments = append(resultComments, comment)
		}
	}

	return resultComments
}

// whether structure has comment with prefix
func (sm *StructMethodModel) HasCommentWithPrefix(prefix string) bool {

	if prefix == "" && len(sm.Comments) > 0 {
		return true
	}

	for _, comment := range sm.Comments {
		if strings.HasPrefix(comment, prefix) {
			return true
		}
	}

	return false
}

// whether structure contains given comment
func (sm *StructMethodModel) HasComment(comment string) bool {

	if len(sm.Comments) == 0 {
		return false
	}

	for _, structComment := range sm.Comments {
		if structComment == comment {
			return true
		}
	}

	return false
}
