package areas

type Result struct {
	ParentId uint64 `json:"parentId" codec:"parentId"`
	Level    uint8  `json:"level" codec:"level"`
	AreaName string `json:"areaName" codec:"areaName"`
	Status   int8   `json:"status" codec:"status"`
	AreaId   uint64 `json:"areaId" codec:"areaId"`
}

type BaseAreasServiceResponse struct {
	Data []Result `json:"data,omitempty" codec:"data,omitempty"`
}
