package tc

import "strconv"

func GetStrToInt(s string) int {
	return GetStrToIntDefault(s, 0)
}

func GetStrToIntDefault(s string, def int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}
