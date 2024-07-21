package tests

import (
	"reflect"
	"testing"

	"github.com/banderveloper/go-ast-describer/model"
)

// mapsEquals compares two maps for equality
func mapsEquals(map1, map2 map[string]string) bool {
	// Check if the length of both maps is the same
	if len(map1) != len(map2) {
		return false
	}

	// Iterate over the first map
	for key, value1 := range map1 {
		// Check if the key exists in the second map
		value2, exists := map2[key]
		if !exists {
			return false
		}

		// Check if the values are the same
		if !reflect.DeepEqual(value1, value2) {
			return false
		}
	}

	return true
}

func TestGetTags(t *testing.T) {

	tcs := []struct {
		model model.StructFieldModel
		exp   map[string]string
	}{
		{
			model: model.StructFieldModel{
				StructTag: `json:"foo" database:"bar"`,
			},
			exp: map[string]string{
				"json":     "foo",
				"database": "bar",
			},
		},
		{
			model: model.StructFieldModel{
				StructTag: `json:"foo,hello" database:"bar"`,
			},
			exp: map[string]string{
				"json":     "foo,hello",
				"database": "bar",
			},
		},
		{
			model: model.StructFieldModel{},
			exp:   map[string]string{},
		},
		{
			model: model.StructFieldModel{
				StructTag: `json:"foo" database:"bar" yield`,
			},
			exp: map[string]string{
				"json":     "foo",
				"database": "bar",
				"yield":    "",
			},
		},
		{
			model: model.StructFieldModel{
				StructTag: `yield`,
			},
			exp: map[string]string{
				"yield": "",
			},
		},
	}

	for i, tc := range tcs {

		act := tc.model.GetTags()
		if !mapsEquals(act, tc.exp) {
			t.Errorf(`[%d] TestGetTags(): got=%v, exp=%v`, i, act, tc.exp)
		}
	}
}
