package filestorage

import "strings"

func (d *S3Client) insertIntoStructure(root *map[string]interface{}, path string, content string) {
	parts := strings.Split(path, "/")
	current := root

	for i, part := range parts {
		if i == len(parts)-1 {
			(*current)[part] = content
		} else {
			if _, exists := (*current)[part]; !exists {
				(*current)[part] = make(map[string]interface{})
			}

			nextMap, ok := (*current)[part].(map[string]interface{})
			if !ok {
				nextMap = make(map[string]interface{})
				(*current)[part] = nextMap
			}
			current = &nextMap
		}
	}
}
