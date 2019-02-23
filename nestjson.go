package nestjson

import (
	"encoding/json"
	"strings"
)

func Str2Map(input string) (map[string]int, error) {
	m := make(map[string]int)
	err := json.Unmarshal([]byte(input), &m)
	return m, err
}

func Nest(m map[string]interface{}, ks []string, v int) {
	if len(ks) == 1 {
		m[ks[0]] = v
		return
	}
	_, ok := m[ks[0]]
	if !ok {
		m[ks[0]] = make(map[string]interface{})
	}
	m1 := m[ks[0]].(map[string]interface{})
	Nest(m1, ks[1:], v)
}

func NestFlattenMap(m map[string]int) map[string]interface{} {
	out := make(map[string]interface{})
	for k, v := range m {
		ks := strings.Split(k, ".")
		Nest(out, ks, v)
	}
	return out
}
