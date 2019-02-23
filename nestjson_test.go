package nestjson_test

import (
	"encoding/json"
	"testing"

	"github.com/hmgle/nestjson"
)

func TestStr2Map(t *testing.T) {
	input := `{"A": 1, "B.A": 2, "B.B": 3, "C.D.E": 4, "C.D.F": 5}`
	m, err := nestjson.Str2Map(input)
	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}
	t.Logf("m: %+v\n", m)
	out := nestjson.NestFlattenMap(m)
	t.Logf("out: %v\n", out)
	jsonStr, _ := json.MarshalIndent(out, "", "\t")
	t.Logf("%s\n", string(jsonStr))
}
