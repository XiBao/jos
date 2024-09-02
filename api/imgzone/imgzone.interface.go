package imgzone

type Categroy struct {
	CateId       int64  `json:"cate_id,omitempty" codec:"cate_id,omitempty"`               // 分类ID
	CateName     string `json:"cate_name,omitempty" codec:"cate_name,omitempty"`           // 分类名称
	CateLevel    int    `json:"cate_level,omitempty" codec:"cate_level,omitempty"`         // 分类层级，默认分类为0，父分类为1，子分类为2
	ParentCateId int64  `json:"parent_cate_id,omitempty" codec:"parent_cate_id,omitempty"` // 父分类ID
	CateOrder    int    `json:"cate_order,omitempty" codec:"cate_order,omitempty"`         // 同级分类排序值，正整数，唯一但不一定连续
	Created      uint64 `json:"created,omitempty" codec:"created,omitempty"`               // 创建时间
	Modified     uint64 `json:"modified,omitempty" codec:"modified,omitempty"`             // 修改时间
}

type Picture struct {
	PictureId     string `json:"picture_id,omitempty" codec:"picture_id,omitempty"`           // 图片ID
	PictureCateId int64  `json:"picture_cate_id,omitempty" codec:"picture_cate_id,omitempty"` // 图片所属分类ID
	PictureUrl    string `json:"picture_url,omitempty" codec:"picture_url,omitempty"`         // 图片url
	PictureName   string `json:"picture_name,omitempty" codec:"picture_name,omitempty"`       // 图片名称
	PictureType   string `json:"picture_type,omitempty" codec:"picture_type,omitempty"`       // 图片后缀 png
	Referenced    int    `json:"referenced,omitempty" codec:"referenced,omitempty"`           // 是否被引用：1，是；0，否
	PictureSize   int    `json:"picture_size,omitempty" codec:"picture_size,omitempty"`       // 图片大小，单位b
	PictureWidth  int    `json:"picture_width,omitempty" codec:"picture_width,omitempty"`     // 图片宽度，单位px
	PictureHeight int    `json:"picture_height,omitempty" codec:"picture_height,omitempty"`   // 图片高度，单位px
	Created       uint64 `json:"created,omitempty" codec:"created,omitempty"`                 // 图片创建时间
	Modified      uint64 `json:"modified,omitempty" codec:"modified,omitempty"`               // 图片修改时间
}
