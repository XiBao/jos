package center

import (
	"errors"

	"github.com/XiBao/jos/api/util"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	center "github.com/XiBao/jos/sdk/request/interact/center/evaluate"
	"github.com/daviddengcn/ljson"
)

type CreateEvaluateActivityWithResultRequest struct {
	api.BaseRequest
	AppName            string `json:"appName" codec:"appName"`                                           // 服务商名称
	Channel            uint8  `json:"channel" codec:"channel"`                                           // 请求渠道 (PC为1, APP为2, 任务中心为3,发现频道为4, 上海运营模板为5 , 微信为 6, QQ为 7, ISV为8)
	SupplierCode       string `json:"supplierCode,omitempty" codec:"supplierCode,omitempty"`             // 供应商简码
	EndTime            string `json:"endTime" codec:"endTime"`                                           // 活动结束时间
	StartTime          string `json:"startTime" codec:"startTime"`                                       // 活动开始时间
	PictureRequirement uint   `json:"pictureRequirement,omitempty" codec:"pictureRequirement,omitempty"` // 晒单图片数量最低要求
	ShopName           string `json:"shopName,omitempty" codec:"shopName,omitempty"`                     // 店铺名称
	Name               string `json:"name" codec:"name"`                                                 // 活动名称
	VedioRequirement   uint   `json:"vedioRequirement,omitempty" codec:"vedioRequirement,omitempty"`     // 视频需要的数量
	WordRequirement    uint   `json:"wordRequirement,omitempty" codec:"wordRequirement,omitempty"`       // 字数需要的数量 页面上大于10
	Modifier           string `json:"modifier,omitempty" codec:"modifier,omitempty"`                     // 活动修改者
	SkuName            string `json:"skuName,omitempty" codec:"skuName,omitempty"`                       // 商品名称
	Cate3rdName        string `json:"cate3rdName,omitempty" codec:"cate3rdName,omitempty"`               // 三级类目名称
	Cate1stName        string `json:"cate1stName,omitempty" codec:"cate1stName,omitempty"`               // 一级类目名称
	Cate3rdCode        string `json:"cate3rdCode,omitempty" codec:"cate3rdCode,omitempty"`               // 三级类目code
	Cate2ndName        string `json:"cate2ndName,omitempty" codec:"cate2ndName,omitempty"`               // 二级类目名称
	Cate1stCode        string `json:"cate1stCode,omitempty" codec:"cate1stCode,omitempty"`               // 一级类目code
	Cate2ndCode        string `json:"cate2ndCode,omitempty" codec:"cate2ndCode,omitempty"`               // 二级类目code
	Price              string `json:"price" codec:"price"`                                               // 商品价格
	SkuId              string `json:"skuId" codec:"skuId"`                                               // SKU编号
	WareId             string `json:"wareId,omitempty" codec:"wareId,omitempty"`                         // 商品编号
	ValidateDay        string `json:"validateDay,omitempty" codec:"validateDay,omitempty"`               // 有效期天数
	AssetItemId        string `json:"assetItemId,omitempty" codec:"assetItemId,omitempty"`               // 权益-资产项id 1 1M流量 2 1元E卡 3 5元E卡 4 20元E卡 5 50元E卡 6 100元E卡 7 京东PLUS会员 8爱奇艺会员月卡 12爱奇艺会员季卡 13 爱奇艺会员半年卡 14爱奇艺会员年卡 15红包
	Type               string `json:"type" codec:"type"`                                                 // 奖项类型 1京豆 2京劵 4东券 6流量包 8积分 10e卡 11plus会员 12爱奇艺会员 14红包
	Discount           string `json:"discount,omitempty" codec:"discount,omitempty"`                     // 优惠金额
	AwardType          string `json:"awardType,omitempty" codec:"awardType,omitempty"`                   // 红包发奖方式类型 1随机 2固定
	Quota              string `json:"quota,omitempty" codec:"quota,omitempty"`                           // 满足优惠的最低消费额
	RulePrice          string `json:"rulePrice" codec:"rulePrice"`                                       // 单位发放金额
	FloatRatio         string `json:"floatRatio,omitempty" codec:"floatRatio,omitempty"`                 // 奖项浮动数（目前仅支持红包）小于 100, 发奖方式为 1的时候必填
	Nums               string `json:"nums" codec:"nums"`                                                 // 总发放人数
	BatchKey           string `json:"batchKey,omitempty" codec:"batchKey,omitempty"`                     // 自营优惠券key
	ExpireType         string `json:"expireType,omitempty" codec:"expireType,omitempty"`                 // pop优惠券日期类型 1相对时间 5 绝对时间
}

type CreateEvaluateActivityWithResultResponse struct {
	ErrorResp *api.ErrorResponnse                   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CreateEvaluateActivityWithResultData `json:"jingdong_com_jd_interact_center_api_write_EvaluateActivityWriteService_createActivityWithResult_responce,omitempty" codec:"jingdong_com_jd_interact_center_api_write_EvaluateActivityWriteService_createActivityWithResult_responce,omitempty"`
}

type CreateEvaluateActivityWithResultData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    uint64 `json:"result,omitempty" codec:"result,omitempty"`
}

func CreateEvaluateActivityWithResult(req *CreateEvaluateActivityWithResultRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewCreateEvaluateActivityWithResultRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetPrice(req.Price)
	r.SetSkuId(req.SkuId)
	r.SetEndTime(req.EndTime)
	r.SetStartTime(req.StartTime)
	r.SetType(req.Type)
	r.SetRulePrice(req.RulePrice)
	r.SetNums(req.Nums)
	r.SetName(req.Name)

	if req.SkuName != "" {
		r.SetSkuName(req.SkuName)
	}

	if req.Cate3rdName != "" {
		r.SetCate3rdName(req.Cate3rdName)
	}

	if req.Cate1stName != "" {
		r.SetCate1stName(req.Cate1stName)
	}

	if req.Cate3rdCode != "" {
		r.SetCate3rdCode(req.Cate3rdCode)
	}

	if req.Cate2ndName != "" {
		r.SetCate2ndName(req.Cate2ndName)
	}

	if req.WareId != "" {
		r.SetWareId(req.WareId)
	}

	if req.Cate1stCode != "" {
		r.SetCate1stCode(req.Cate1stCode)
	}

	if req.Cate2ndCode != "" {
		r.SetCate2ndCode(req.Cate2ndCode)
	}

	if req.SupplierCode != "" {
		r.SetSupplierCode(req.SupplierCode)
	}

	if req.Modifier != "" {
		r.SetModifier(req.Modifier)
	}

	if req.PictureRequirement > 0 {
		r.SetPictureRequirement(req.PictureRequirement)
	}

	if req.ShopName != "" {
		r.SetShopName(req.ShopName)
	}

	if req.ValidateDay != "" {
		r.SetValidateDay(req.ValidateDay)
	}

	if req.AssetItemId != "" {
		r.SetAssetItemId(req.AssetItemId)
	}

	if req.Discount != "" {
		r.SetDiscount(req.Discount)
	}

	if req.AwardType != "" {
		r.SetAwardType(req.AwardType)
	}

	if req.Quota != "" {
		r.SetQuota(req.Quota)
	}

	if req.FloatRatio != "" {
		r.SetFloatRatio(req.FloatRatio)
	}

	if req.BatchKey != "" {
		r.SetBatchKey(req.BatchKey)
	}

	if req.ExpireType != "" {
		r.SetExpireType(req.ExpireType)
	}

	if req.VedioRequirement > 0 {
		r.SetVedioRequirement(req.VedioRequirement)
	}

	if req.WordRequirement > 0 {
		r.SetWordRequirement(req.WordRequirement)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	result = util.RemoveJsonSpace(result)

	var response CreateEvaluateActivityWithResultResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.Result, nil
}
