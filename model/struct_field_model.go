package model

import "strings"

// description of field of structure
type StructFieldModel struct {
	Name      string // field name
	Type      string // field type
	StructTag string // field struct tag without ``
}

// get field's struct tags as KV map
func (sfm *StructFieldModel) GetTags() map[string]string {

	result := make(map[string]string, 0)

	// if field has no tag - result empty map
	if sfm.StructTag == "" {
		return result
	}

	// split tag by parts if case if there are a few keys
	parts := strings.Split(sfm.StructTag, " ")

	for _, part := range parts {
		keyValue := strings.SplitN(part, ":", 2)
		if len(keyValue) == 2 {
			key := keyValue[0]
			value := strings.Trim(keyValue[1], `"`)
			result[key] = value
		} else if len(keyValue) == 1 {
			// Handle tags without values (e.g., `key`)
			result[keyValue[0]] = ""
		}
	}

	return result
}
