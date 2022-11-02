package order

import (
	"strconv"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/order"
	"github.com/daviddengcn/ljson"
)

type UnionOrderQueryRowRequest struct {
	api.BaseRequest
	PageIndex    uint   `json:"pageIndex`
	PageSize     uint   `json:"pageSize"`               // 每页包含条数，上限为500
	Type         uint   `json:"type"`                   // 订单时间查询类型(1：下单时间，2：完成时间（购买用户确认收货时间），3：更新时间
	StartTime    string `json:"startTime"`              // 开始时间 格式yyyy-MM-dd HH:mm:ss，与endTime间隔不超过1小时
	EndTime      string `json:"endTime"`                // 结束时间 格式yyyy-MM-dd HH:mm:ss，与startTime间隔不超过1小时
	ChildUnionId string `json:"childUnionId,omitempty"` // 子推客unionID，传入该值可查询子推客的订单，注意不可和key同时传入。（需联系运营开通PID权限才能拿到数据）
	Key          string `json:"key,omitempty"`          // 工具商传入推客的授权key，可帮助该推客查询订单，注意不可和childUnionid同时传入。（需联系运营开通工具商权限才能拿到数据）
	Fields       string `json:"fields,omitempty"`       // 支持出参数据筛选，逗号','分隔，目前可用：goodsInfo（商品信息）,categoryInfo(类目信息）
}

type UnionOrderQueryRowResponse struct {
	ErrorResp *api.ErrorResponnse             `json:"error_response,omitempty"`
	Data      *UnionOrderQueryRowResponseData `json:"jd_union_open_order_row_query_responce,omitempty"`
}

type UnionOrderQueryRowResponseData struct {
	Result string `json:"queryResult,omitempty"`
}

type UnionOrderQueryRowResult struct {
	Code    int            `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Data    []OrderRowResp `json:"data,omitempty"`
	HasMore bool           `json:"hasMore,omitempty"`
}

type OrderRowResp struct {
	Id                  string    `json:"id"`                       // 标记唯一订单行：订单+sku维度的唯一标识
	OrderId             uint64    `json:"orderId"`                  // 订单号
	ParentId            uint64    `json:"parentId"`                 // 父单的订单号：如一个订单拆成多个子订单时，原订单号会作为父单号，拆分的订单号为子单号存储在orderid中。若未发生拆单，该字段为0
	OrderTime           string    `json:"orderTime"`                // 下单时间,格式yyyy-MM-dd HH:mm:ss
	FinishTime          string    `json:"finishTime"`               // 完成时间（购买用户确认收货时间）,格式yyyy-MM-dd HH:mm:ss
	ModifyTime          string    `json:"modifyTime`                // 更新时间,格式yyyy-MM-dd HH:mm:ss
	OrderEmt            uint      `json:"orderEmt"`                 // 下单设备 1.pc 2.无线
	Plus                uint      `json:"plus"`                     // 下单用户是否为PLUS会员 0：否，1：是
	UnionId             uint64    `json:"unionId"`                  // 推客ID
	SkuId               uint64    `json:"skuId"`                    // 商品ID
	SkuName             string    `json:"skuName"`                  // 商品名称
	SkuNum              uint      `json:"skuNum"`                   // 商品数量
	SkuReturnNum        uint      `json:"skuReturnNum"`             // 商品已退货数量
	SkuFrozenNum        uint      `json:"skuFrozenNum"`             // 商品售后中数量
	Price               float64   `json:"price"`                    // 商品单价
	CommissionRate      float64   `json:"commissionRate"`           // 佣金比例(投放的广告主计划比例)
	SubSideRate         float64   `json:"subSideRate"`              // 分成比例（单位：%）
	SubsidyRate         float64   `json:"subsidyRate"`              // 补贴比例（单位：%）
	FinalRate           float64   `json:"finalRate"`                // 最终分佣比例（单位：%）=分成比例+补贴比例
	EstimateCosPrice    float64   `json:"estimateCosPrice"`         // 预估计佣金额：由订单的实付金额拆分至每个商品的预估计佣金额，不包括运费，以及京券、东券、E卡、余额等虚拟资产支付的金额。该字段仅为预估值，实际佣金以actualCosPrice为准进行计算
	EstimateFee         float64   `json:"estimateFee"`              // 推客的预估佣金（预估计佣金额*佣金比例*最终比例），如订单完成前发生退款，此金额也会更新。
	ActualCosPrice      float64   `json:"actualCosPrice"`           // 实际计算佣金的金额。订单完成后，会将误扣除的运费券金额更正。如订单完成后发生退款，此金额会更新。
	ActualFee           float64   `json:"actualFee"`                // 推客分得的实际佣金（实际计佣金额*佣金比例*最终比例）。如订单完成后发生退款，此金额会更新。
	ValidCode           int       `json:"validCode"`                // sku维度的有效码（-1：未知,2.无效-拆单,3.无效-取消,4.无效-京东帮帮主订单,5.无效-账号异常,6.无效-赠品类目不返佣,7.无效-校园订单,8.无效-企业订单,9.无效-团购订单,11.无效-乡村推广员下单,13.无效-违规订单,14.无效-来源与备案网址不符,15.待付款,16.已付款,17.已完成（购买用户确认收货）,20.无效-此复购订单对应的首购订单无效,21.无效-云店订单，22.无效-PLUS会员佣金比例为0，23.无效-支付有礼，24.已付定金
	TraceType           uint      `json:"traceType"`                // 同跨店：2同店 3跨店
	PositionId          uint64    `json:"positionId"`               // 推广位ID
	SiteId              uint64    `json:"siteId"`                   // 应用id（网站id、appid、社交媒体id）
	UnionAlias          string    `json:"unionAlias"`               // PID所属母账号平台名称（原第三方服务商来源），两方分佣会有该值
	Pid                 string    `json:"pid"`                      // 格式:子推客ID_子站长应用ID_子推客推广位ID
	Cid1                uint64    `json:"cid1"`                     // 一级类目id
	Cid2                uint64    `json:"cid2"`                     // 二级类目id
	Cid3                uint64    `json:"cid3"`                     // 三级类目id
	SubUnionId          string    `json:"subUnionId"`               // 子渠道标识，在转链时可自定义传入，格式要求：字母、数字或下划线，最多支持80个字符(需要联系运营开放白名单才能拿到数据)
	UnionTag            string    `json:"unionTag"`                 // 联盟标签数据（32位整型二进制字符串：00000000000000000000000000000001。数据从右向左进行，每一位为1表示符合特征，第1位：红包，第2位：组合推广，第3位：拼购，第5位：有效首次购（0000000000011XXX表示有效首购，最终奖励活动结算金额会结合订单状态判断，以联盟后台对应活动效果数据报表https://union.jd.com/active为准）,第8位：复购订单，第9位：礼金，第10位：联盟礼金，第11位：推客礼金，第12位：京喜APP首购，第13位：京喜首购，第14位：京喜复购，第15位：京喜订单，第16位：京东极速版APP首购，第17位白条首购，第18位校园订单，第19位是0或1时，均代表普通订单，第20位：预售订单，第21位：学生订单，第22位：全球购订单 例如：00000000000000000000000000000001:红包订单，00000000000000000000000000000010:组合推广订单，00000000000000000000000000000100:拼购订单，00000000000000000000000000011000:有效首购，00000000000000000000000000000111：红包+组合推广+拼购等） 注：一个订单同时使用礼金和红包，仅礼金位数为1，红包位数为0
	PopId               uint64    `json:"popId"`                    // 商家ID
	Ext1                string    `json:"ext1"`                     // 推客生成推广链接时传入的扩展字段（需要联系运营开放白名单才能拿到数据）。
	PayMonth            uint      `json:"payMonth"`                 // 预估结算时间，订单完成后才会返回，格式：yyyyMMdd，默认：0。表示最新的预估结算日期。当payMonth为当前的未来时间时，表示该订单可结算；当payMonth为当前的过去时间时，表示该订单已结算
	CpActId             uint64    `json:"cpActId"`                  // 招商团活动id：当商品参加了招商团会有该值，为0时表示无活动cpActId
	UnionRole           uint      `json:"unionRole"`                // 站长角色：1 推客 2 团长 3内容服务商
	GiftCouponOcsAmount float64   `json:"giftCouponOcsAmount"`      // 礼金分摊金额：使用礼金的订单会有该值
	GiftCouponKey       string    `json:"giftCouponKey"`            // 礼金批次ID：使用礼金的订单会有该值
	BalanceExt          string    `json:"balanceExt"`               // 计佣扩展信息，表示结算月:每月实际佣金变化情况，格式：{20191020:10,20191120:-2}，订单完成后会有该值
	Sign                string    `json:"sign,omitempty"`           // 数据签名，用来核对出参数据是否被修改，入参fields中写入sign时返回
	ProPriceAmount      float64   `json:"proPriceAmount,omitempty"` // 价保赔付金额：订单申请价保或赔付的金额，实际计佣金额已经减去此金额，您无需处理
	Rid                 uint64    `json:"rid,omitempty"`            // 团长渠道ID，仅限招商团长管理渠道使用，团长开通权限后才可使用。
	ExpressStatus       uint      `json:"expressStatus,omitempty"`  // 发货状态（10：待发货，20：已发货）
	ChannelId           uint      `json:"channelId,omitempty"`      // 渠道关系ID
	GoodsInfo           *GoodInfo `json:"goodsInfo,omitempty"`      // 商品信息，入参传入fields，goodsInfo获取
	CategoryInfo        *CatInfo  `json:"categoryInfo,omitempty"`   // 类目信息,入参传入fields，categoryInfo获取
}

type CatInfo struct {
	Cid1     uint64 `json:"cid1"`     // 一级类目id
	Cid2     uint64 `json:"cid2"`     // 二级类目id
	Cid3     uint64 `json:"cid3"`     // 三级类目id
	Cid1Name string `json:"cid1Name"` // 一级类目名称
	Cid2Name string `json:"cid2Name"` // 二级类目名称
	Cid3Name string `json:"cid3Name"` // 三级类目名称
}

type GoodInfo struct {
	ImageUrl  string `json:"imageUrl,omitempty"`  // sku主图链接
	Owner     string `json:"owner,omitempty"`     // g=自营，p=pop
	MainSkuId uint64 `json:"mainSkuId,omitempty"` // 自营商品主Id（owner=g取此值）
	ProductId uint64 `json:"productId,omitempty"` // 非自营商品主Id（owner=p取此值）
	ShopName  string `json:"shopName,omitempty"`  // 店铺名称（或供应商名称）
	ShopId    int64  `json:"shopId,omitempty"`    // 店铺Id
}

// 订单查询接口
func UnionOrderQueryRow(req *UnionOrderQueryRowRequest) (bool, []OrderRowResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewUnionOrderQueryRowRequest()
	orderReq := &order.OrderRowReq{
		PageIndex:    req.PageIndex,
		PageSize:     req.PageSize,
		Type:         req.Type,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		ChildUnionId: req.ChildUnionId,
		Key:          req.Key,
		Fields:       req.Fields,
	}
	r.SetOrderRowReq(orderReq)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, nil, err
	}
	var response UnionOrderQueryRowResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, nil, err
	}
	if response.Data == nil {
		return false, nil, nil
	}
	var ret UnionOrderQueryRowResult
	err = ljson.Unmarshal([]byte(response.Data.Result), &ret)
	if err != nil {
		return false, nil, err
	}

	if ret.Code != 200 {
		return false, nil, &api.ErrorResponnse{Code: strconv.FormatInt(int64(ret.Code), 10), ZhDesc: ret.Message}
	}

	return ret.HasMore, ret.Data, nil
}
