package logger

import (
	"encoding/json"
)

type DefaultLogger struct{}

func (l DefaultLogger) DebugPrintError(err error) {}

func (l DefaultLogger) DebugPrintStringResponse(str string) {}

func (l DefaultLogger) DebugPrintGetRequest(url string) {}

func (l DefaultLogger) DebugPrintPostJSONRequest(url string, body []byte) {}

func (l DefaultLogger) DebugPrintPostMultipartRequest(url string, body []byte) {}

func (l DefaultLogger) DecodeJSON(body []byte, v interface{}) error {
	bs := removeJsonSpace(body)
	return json.Unmarshal(bs, v)
}
