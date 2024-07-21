package tests

import (
	"testing"

	"github.com/banderveloper/go-ast-describer/model"
)

func slicesEquals[T comparable](slice1, slice2 []T) bool {

	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func TestGetCommentsWithPrefix(t *testing.T) {

	// test cases
	tcs := []struct {
		model  model.StructModel // structure description struct
		prefix string            // comment prefix (function accepts)
		exp    []string          // expected slice of comments with given prefix
	}{
		{
			model: model.StructModel{
				Comments: []string{"codegen: hello world", "not_codegen: foobar", "codegen: fizbuzz"},
			},
			prefix: "codegen",
			exp:    []string{"codegen: hello world", "codegen: fizbuzz"},
		},
		{
			model: model.StructModel{
				Comments: []string{},
			},
			prefix: "",
			exp:    []string{},
		},
		{
			model: model.StructModel{
				Comments: []string{},
			},
			prefix: "codegen",
			exp:    []string{},
		},
		{
			model: model.StructModel{
				Comments: []string{"not_codegen: foobar", "not_codegen2: foobar"},
			},
			prefix: "codegen",
			exp:    []string{},
		},
		{
			model: model.StructModel{
				Comments: []string{"not_codegen: foobar", "not_codegen: foobar"},
			},
			prefix: "",
			exp:    []string{"not_codegen: foobar", "not_codegen: foobar"},
		},
	}

	for i, tc := range tcs {

		act := tc.model.GetCommentsWithPrefix(tc.prefix)
		if !slicesEquals(act, tc.exp) {
			t.Errorf(`[%d] TestGetCommentsWithPrefix("%s"): got=%v, exp=%v`, i, tc.prefix, act, tc.exp)
		}
	}
}

func TestHasCommentWithPrefix(t *testing.T) {

	tcs := []struct {
		model  model.StructModel
		prefix string
		exp    bool
	}{
		{
			model: model.StructModel{
				Comments: []string{"codegen: hello world", "not_codegen: foobar", "codegen: fizbuzz"},
			},
			prefix: "codegen",
			exp:    true,
		},
		{
			model: model.StructModel{
				Comments: []string{"codegen: hello world", "not_codegen: foobar", "codegen: fizbuzz"},
			},
			prefix: "not_codegen",
			exp:    true,
		},
		{
			model: model.StructModel{
				Comments: []string{"codegen: hello world", "not_codegen: foobar", "codegen: fizbuzz"},
			},
			prefix: "non_existing_prefix",
			exp:    false,
		},
		{
			model: model.StructModel{
				Comments: []string{},
			},
			prefix: "",
			exp:    false,
		},
		{
			model: model.StructModel{
				Comments: []string{"not_codegen: foobar"},
			},
			prefix: "",
			exp:    true,
		},
	}

	for i, tc := range tcs {

		act := tc.model.HasCommentWithPrefix(tc.prefix)
		if act != tc.exp {
			t.Errorf(`[%d] TestHasCommentWithPrefix("%s"): got=%v, exp=%v`, i, tc.prefix, act, tc.exp)
		}
	}
}

func TestHasComment(t *testing.T) {

	tcs := []struct {
		model   model.StructModel
		comment string
		exp     bool
	}{
		{
			model: model.StructModel{
				Comments: []string{"codegen: hello world", "hello world", "codegen: fizbuzz"},
			},
			comment: "hello world",
			exp:     true,
		},
		{
			model: model.StructModel{
				Comments: []string{"codegen: hello world", "not_codegen: foobar", "codegen: fizbuzz"},
			},
			comment: "not_codege",
			exp:     false,
		},
		{
			model: model.StructModel{
				Comments: []string{"codegen: hello world", "not_codegen: foobar", "codegen: fizbuzz"},
			},
			comment: "not_codegen",
			exp:     false,
		},
		{
			model: model.StructModel{
				Comments: []string{},
			},
			comment: "",
			exp:     false,
		},
		{
			model: model.StructModel{
				Comments: []string{},
			},
			comment: "test",
			exp:     false,
		},
	}

	for i, tc := range tcs {

		act := tc.model.HasComment(tc.comment)
		if act != tc.exp {
			t.Errorf(`[%d] TestHasComment("%s"): got=%v, exp=%v`, i, tc.comment, act, tc.exp)
		}
	}
}
