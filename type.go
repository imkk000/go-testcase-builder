package tc

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func GetKindOf[T any](t T) reflect.Kind {
	typeOf := reflect.TypeOf(t)
	switch typeOf.Kind() {
	case reflect.Ptr:
		return typeOf.Elem().Kind()
	}
	return typeOf.Kind()
}

func GetElements(s string) []string {
	return GetAllValues(`\[([\w\d,."+-]*)\]`, s)
}

func GetAllString(s string) []string {
	elms := GetAllValues(`"(\w*)"`, s)
	for i := range elms {
		elms[i] = strings.Trim(elms[i], "\"")
	}
	return elms
}

func GetAllBool(s string) []bool {
	elms := GetAllValues(`(true|false)`, strings.ToLower(s))
	v := make([]bool, 0)
	for i := range elms {
		b, err := strconv.ParseBool(elms[i])
		if err != nil {
			continue
		}
		v = append(v, b)
	}
	return v
}

func GetAllInt(s string) []int {
	elms := GetAllValues(`([-+\d]*)`, s)
	v := make([]int, 0)
	for _, e := range elms {
		e = strings.TrimSpace(e)
		n, err := strconv.Atoi(e)
		if err != nil {
			continue
		}
		v = append(v, n)
	}
	return v
}

func GetAllFloat64(s string) []float64 {
	elms := GetAllValues(`([-+.\d]*)`, s)
	v := make([]float64, 0)
	for _, e := range elms {
		n, err := strconv.ParseFloat(e, 64)
		if err != nil {
			continue
		}
		v = append(v, n)
	}
	return v
}

func GetAllValues(pattern, s string) []string {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return []string{}
	}
	elms := regex.FindAllString(s, -1)
	if elms == nil {
		return []string{}
	}
	return elms
}
