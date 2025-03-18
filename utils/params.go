package utils

import "strconv"

func ParseIntQueryParam(queryParams map[string][]string, key string) *int {
	if val, ok := queryParams[key]; ok {
		if parsed, err := strconv.Atoi(val[0]); err == nil {
			return &parsed
		}
	}
	return nil
}