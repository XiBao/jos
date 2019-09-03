package sdk

import (
	"bytes"
	"encoding/json"
)

func Json(obj interface{}) []byte {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(obj)

	return buf.Byte()
}
