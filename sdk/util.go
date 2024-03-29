package sdk

import (
	"bytes"
	"encoding/json"
)

func Json(obj interface{}) []byte {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(obj)

	return buf.Bytes()
}

func StringsJoin(strs ...string) string {
	var n int
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}
	if n <= 0 {
		return ""
	}
	builder := GetStringsBuilder()
	builder.Grow(n)
	for _, s := range strs {
		builder.WriteString(s)
	}
	ret := builder.String()
	PutStringsBuilder(builder)
	return ret
}
