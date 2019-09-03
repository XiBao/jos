package dsp

type Paginator struct {
	PageNum  int `json:"pageNum,omitempty" codec:"pageNum,omitempty"`
	Items    int `json:"items,omitempty" codec:"items,omitempty"`
	PageSize int `json:"pageSize,omitempty" codec:"pageSize,omitempty"`
}
