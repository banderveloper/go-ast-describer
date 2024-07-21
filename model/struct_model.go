package model

import "strings"

// description of struct
type StructModel struct {
	Name     string              // typename of structure
	Comments []string            // list of comments above structure, without // and trimmed
	Fields   []StructFieldModel  // list of fields
	Methods  []StructMethodModel // list of methods
}

// get all structure comments with given prefix, including it
func (sm *StructModel) GetCommentsWithPrefix(prefix string) []string {

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
func (sm *StructModel) HasCommentWithPrefix(prefix string) bool {

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
func (sm *StructModel) HasComment(comment string) bool {

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
