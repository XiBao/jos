package sdk

import (
	"strconv"
)

type Error struct {
	Code    int
	SubCode string
	Msg     string
	SubMsg  string
}

func (e Error) Error() string {
	return StringsJoin("CODE:", strconv.Itoa(e.Code), ", SUB_CODE:", e.SubCode, ", MSG:", e.Msg, ", SUB_MSG:", e.SubMsg)
}

func ErrorString[T int | int64 | string](code T, msg string) string {
	switch t := any(code).(type) {
	case int:
		return StringsJoin("code:", strconv.Itoa(t), ", msg:", msg)
	case int64:
		return StringsJoin("code:", strconv.FormatInt(t, 10), ", msg:", msg)
	case string:
		return StringsJoin("code:", t, ", msg:", msg)
	}
	return msg
}
