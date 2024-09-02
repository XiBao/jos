package center

import (
	"github.com/XiBao/jos/sdk"
)

type CreateEvaluateActivityRequest struct {
	Request *sdk.Request
}

type CreateEvaluateActivityClientSource struct {
	AppName string `json:"appName" codec:"appName"` // 服务商名称
	Channel uint8  `json:"channel" codec:"channel"` // 请求渠道 (PC为1, APP为2, 任务中心为3,发现频道为4, 上海运营模板为5 , 微信为 6, QQ为 7, ISV为8)
}

type CreateEvaluateActivityRule struct {
	ValidateDay uint    `json:"validateDay,omitempty" codec:"validateDay,omitempty"` // 有效期天数
	AssetItemId uint8   `json:"assetItemId,omitempty" codec:"assetItemId,omitempty"` // 权益-资产项id 1 1M流量 2 1元E卡 3 5元E卡 4 20元E卡 5 50元E卡 6 100元E卡 7 京东PLUS会员 8爱奇艺会员月卡 12爱奇艺会员季卡 13 爱奇艺会员半年卡 14爱奇艺会员年卡 15红包
	Type        uint8   `json:"type" codec:"type"`                                   // 奖项类型 1京豆 2京劵 4东券 6流量包 8积分 10e卡 11plus会员 12爱奇艺会员 14红包
	Discount    float64 `json:"discount,omitempty" codec:"discount,omitempty"`       // 优惠金额
	AwardType   uint8   `json:"awardType,omitempty" codec:"awardType,omitempty"`     // 红包发奖方式类型 1随机 2固定
	Quota       uint    `json:"quota,omitempty" codec:"quota,omitempty"`             // 满足优惠的最低消费额
	RulePrice   float64 `json:"rulePrice" codec:"rulePrice"`                         // 单位发放金额
	FloatRatio  float64 `json:"floatRatio,omitempty" codec:"floatRatio,omitempty"`   // 奖项浮动数（目前仅支持红包）小于 100, 发奖方式为 1的时候必填
	Nums        uint    `json:"nums" codec:"nums"`                                   // 总发放人数
	BatchKey    string  `json:"batchKey,omitempty" codec:"batchKey,omitempty"`       // 自营优惠券key
	ExpireType  uint8   `json:"expireType,omitempty" codec:"expireType,omitempty"`   // pop优惠券日期类型 1相对时间 5 绝对时间
}

type CreateEvaluateActivitySku struct {
	SkuName     string  `json:"skuName,omitempty" codec:"skuName,omitempty"`         // 商品名称
	Cate3rdName string  `json:"cate3rdName,omitempty" codec:"cate3rdName,omitempty"` // 三级类目名称
	Cate1stName string  `json:"cate1stName,omitempty" codec:"cate1stName,omitempty"` // 一级类目名称
	Cate3rdCode uint    `json:"cate3rdCode,omitempty" codec:"cate3rdCode,omitempty"` // 三级类目code
	Cate2ndName string  `json:"cate2ndName,omitempty" codec:"cate2ndName,omitempty"` // 二级类目名称
	Cate1stCode uint    `json:"cate1stCode,omitempty" codec:"cate1stCode,omitempty"` // 一级类目code
	Cate2ndCode uint    `json:"cate2ndCode,omitempty" codec:"cate2ndCode,omitempty"` // 二级类目code
	Price       float64 `json:"price" codec:"price"`                                 // 商品价格
	SkuId       uint64  `json:"skuId" codec:"skuId"`                                 // SKU编号
	WareId      uint64  `json:"wareId,omitempty" codec:"wareId,omitempty"`           // 商品编号
}

type CreateEvaluateActivityBody struct {
	SupplierCode       string                       `json:"supplierCode,omitempty" codec:"supplierCode,omitempty"`             // 供应商简码
	EndTime            string                       `json:"endTime" codec:"endTime"`                                           // 活动结束时间
	StartTime          string                       `json:"startTime" codec:"startTime"`                                       // 活动开始时间
	PictureRequirement uint                         `json:"pictureRequirement,omitempty" codec:"pictureRequirement,omitempty"` // 晒单图片数量最低要求
	ShopName           string                       `json:"shopName,omitempty" codec:"shopName,omitempty"`                     // 店铺名称
	Name               string                       `json:"name" codec:"name"`                                                 // 活动名称
	VedioRequirement   uint                         `json:"vedioRequirement,omitempty" codec:"vedioRequirement,omitempty"`     // 视频需要的数量
	WordRequirement    uint                         `json:"wordRequirement,omitempty" codec:"wordRequirement,omitempty"`       // 字数需要的数量 页面上大于10
	Modifier           string                       `json:"modifier,omitempty" codec:"modifier,omitempty"`                     // 活动修改者
	EvaluateRuleList   []CreateEvaluateActivityRule `json:"evaluateRuleList" codec:"evaluateRuleList"`                         // 活动规则集合
	EvaluateSkuList    []CreateEvaluateActivitySku  `json:"evaluateSkuList" codec:"evaluateSkuList"`                           // 活动商品集合
}

// create new request
func NewCreateEvaluateActivityRequest() (req *CreateEvaluateActivityRequest) {
	request := sdk.Request{MethodName: "jingdong.com.jd.interact.center.api.service.write.EvaluateActivityWriteService.createActivity", Params: make(map[string]interface{})}
	req = &CreateEvaluateActivityRequest{
		Request: &request,
	}
	return
}

func (req *CreateEvaluateActivityRequest) SetClientSource(clientSource *CreateEvaluateActivityClientSource) {
	req.Request.Params["ClientSource"] = clientSource
}

func (req *CreateEvaluateActivityRequest) GetClientSource() *CreateEvaluateActivityClientSource {
	clientSource, found := req.Request.Params["ClientSource"]
	if found {
		return clientSource.(*CreateEvaluateActivityClientSource)
	}
	return nil
}

func (req *CreateEvaluateActivityRequest) SetEvaluateActivity(evaluateActivity *CreateEvaluateActivityBody) {
	req.Request.Params["EvaluateActivity"] = evaluateActivity
}

func (req *CreateEvaluateActivityRequest) GetEvaluateActivity() *CreateEvaluateActivityBody {
	evaluateActivity, found := req.Request.Params["EvaluateActivity"]
	if found {
		return evaluateActivity.(*CreateEvaluateActivityBody)
	}
	return nil
}
