package tc

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func MakeIntSlice(s string) []int {
	s = strings.TrimFunc(s, func(r rune) bool {
		return r == '[' || r == ']'
	})
	if len(s) == 0 {
		return []int{}
	}

	getInt := func(s string) int {
		v, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return v
	}
	elms := strings.Split(s, ",")
	m := make([]int, 0)
	for _, e := range elms {
		e = strings.TrimSpace(e)
		m = append(m, getInt(e))
	}
	return m
}

func MakeStringSlice(s string) []string {
	s = strings.TrimFunc(s, func(r rune) bool {
		return r == '[' || r == ']'
	})
	if len(s) == 0 {
		return []string{}
	}

	elms := strings.Split(s, ",")
	m := make([]string, 0)
	for _, e := range elms {
		e = strings.TrimSpace(e)
		m = append(m, e[1:len(e)-1])
	}
	return m
}

func GetSlice[E any](v []E, b ...int) []E {
	abs := func(n int) int {
		if n < 0 {
			return -1 * n
		}
		return n
	}

	switch len(b) {
	case 2:
		return v[b[0]:b[1]]
	case 1:
		n := b[0]
		if n < 0 {
			return v[len(v)-abs(n):]
		}
		return v[:n]
	}
	return v
}

func MakeSliceStr[S ~[]E, E any](s S, sep ...string) string {
	o := make([]string, len(s))
	for i := range s {
		switch reflect.TypeOf(s[i]).Kind() {
		case reflect.String:
			o[i] = fmt.Sprintf(`"%v"`, s[i])
			continue
		}
		o[i] = fmt.Sprintf("%v", s[i])
	}

	sp := ","
	if sep != nil {
		sp = sep[0]
	}
	return "[" + strings.Join(o, sp) + "]"
}
