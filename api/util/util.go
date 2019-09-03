package util

import (
	"bytes"
	"encoding/json"
)

func Json(obj interface{}) string {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(obj)

	return buf.String()
}

func RemoveJsonSpace(data []byte) []byte {
	rs := bytes.Replace(data, []byte("\n"), []byte(""), -1)
	rs = bytes.Replace(rs, []byte("\t"), []byte(""), -1)
	return rs
}
