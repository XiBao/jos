package sdk

import (
	"fmt"
)

type Error struct {
	Code    int
	SubCode string
	Msg     string
	SubMsg  string
}

func (e Error) Error() string {
	return fmt.Sprintf("CODE:%v, SUB_CODE:%v, MSG:%v, SUB_MSG:%v", e.Code, e.SubCode, e.Msg, e.SubMsg)
}
