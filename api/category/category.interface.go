package category

import (
	"github.com/XiBao/jos/api/ware"
)

type Category struct {
	Fid      uint64         `json:"fid,omitempty" codec:"fid,omitempty"`           // 	类目父ID
	Id       uint64         `json:"id,omitempty" codec:"id,omitempty"`             // 类目id
	Lev      int            `json:"lev,omitempty" codec:"lev,omitempty"`           // 	类目级别
	Name     string         `json:"name,omitempty" codec:"name,omitempty"`         // 类目名称
	Order    int            `json:"order,omitempty" codec:"order,omitempty"`       // 排序
	Features []ware.Feature `json:"features,omitempty" codec:"features,omitempty"` // 类目特殊属性列表
}
