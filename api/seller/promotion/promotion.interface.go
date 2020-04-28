package promotion

type PromotionList struct {
	AccountId             uint64 `json:"vender_id" codec:"vender_id"`
	PromoId               uint64 `json:"promo_id" codec:"promo_id"`
	PromoName             string `json:"promo_name" codec:"promo_name"`
	PromoType             uint8  `json:"promo_type" codec:"promo_type"`                     // 促销类型。1：单品促销 4：赠品促销 6 套装促销 10：总价促销
	FavorMode             uint8  `json:"favor_mode,omitempty" codec:"favor_mode,omitempty"` // 促销子类型。0:满赠,1:满减,2:每满减, 3:百分比满减, 4: 阶梯满减,5:满赠加价购,6:满M件减N件,7:阶梯买M件减N件,13:M元任选N件,15:M件N折,16:满减送（元）
	BeginTime             string `json:"begin_time" codec:"begin_time"`
	EndTime               string `json:"end_time" codec:"end_time"`
	Bound                 uint8  `json:"bound" codec:"bound"`                         // 促销范围。1：部分商品参加 2：全部商品参加 3：部分商品部参加 4：全场促销
	Member                uint   `json:"member" codec:"member"`                       // 参加促销的会员级别
	Slogan                string `json:"slogan,omitempty" codec:"slogan,omitempty"`   // 促销宣传语
	Comment               string `json:"comment,omitempty" codec:"comment,omitempty"` // 促销备注信息
	PromoStatus           uint8  `json:"promo_status" codec:"promo_status"`           // 促销状态。-1:促销被创建但还没commit,0:无效,1:驳回,2:未审核,3:人工审,4:已审核,5:已生效,6:已暂停,7:强制暂停
	Created               uint64 `json:"created" codec:"created"`
	Modified              uint64 `json:"modified" codec:"modified"`
	Platform              uint8  `json:"platform" codec:"platform"`                                       // 平台类型（1、全渠道，2、手机客户端，3、微信，5、手Q）
	Link                  string `json:"link,omitempty" codec:"link,omitempty"`                           // 活动链接
	ShopMember            uint   `json:"shop_member,omitempty" codec:"shop_member,omitempty"`             // 店铺会员级别
	QqMember              uint   `json:"qq_member,omitempty" codec:"qq_member,omitempty"`                 // QQ会员级别
	PlusMember            uint   `json:"plus_member,omitempty" codec:"plus_member,omitempty"`             // plus会员级别
	MemberLevelOnly       bool   `json:"member_level_only" codec:"member_level_only"`                     // 是否是会员级别专享
	AllowOthersOperate    bool   `json:"allow_others_operate" codec:"allow_others_operate"`               // 是否允许其他来源操作该促销
	AllowOthersCheck      bool   `json:"allow_others_check" codec:"allow_others_check"`                   // 是否允许其他来源审核该促销
	AllowOtherUserOperate bool   `json:"allow_other_user_operate" codec:"allow_other_user_operate"`       // 是否允许其他人操作该促销
	AllowOtherUserCheck   bool   `json:"allow_other_user_check" codec:"allow_other_user_check"`           // 是否允许其他人审核该促销
	NeedManualCheck       bool   `json:"need_manual_check" codec:"need_manual_check"`                     // 促销是否需要人工审核
	AllowCheck            bool   `json:"allow_check" codec:"allow_check"`                                 // 是否允许审核该促销
	AllowOperate          bool   `json:"allow_operate" codec:"allow_operate"`                             // 是否允许操作该促销
	IsJingdouRequired     bool   `json:"is_jingdou_required" codec:"is_jingdou_required"`                 // 是否是京豆优惠购促销销
	FreqBound             uint8  `json:"freq_bound,omitempty" codec:"freq_bound,omitempty"`               // 促销限购一次形式：（0，不限，1、限ip、账号 2、限ip 3、限账号）
	PerMaxNum             uint   `json:"per_max_num,omitempty" codec:"per_max_num,omitempty"`             // 单次最大购买数量：0、不限
	PerMinNum             uint   `json:"per_min_num,omitempty" codec:"per_min_num,omitempty"`             // 单次最小购买数量：0、不限
	PropType              uint8  `json:"prop_type,omitempty" codec:"prop_type,omitempty"`                 // 道具类型：2、京豆，10 、店铺京券
	PropNum               uint   `json:"prop_num,omitempty" codec:"prop_num,omitempty"`                   // 道具数值
	PropUsedWay           uint8  `json:"prop_used_way,omitempty" codec:"prop_used_way,omitempty"`         // 道具使用方式：1、消耗，2、奖励
	CouponId              uint   `json:"coupon_id,omitempty" codec:"coupon_id,omitempty"`                 // 赠送优惠券ID
	CouponBatchKey        string `json:"coupon_batch_key,omitempty" codec:"coupon_batch_key,omitempty"`   // 赠送优惠券批次key
	CouponValidDays       uint   `json:"coupon_valid_days,omitempty" codec:"coupon_valid_days,omitempty"` // 优惠券的有效天数
	Quota                 string `json:"quota,omitempty" codec:"quota,omitempty"`                         // 订单额度
	Rate                  string `json:"rate,omitempty" codec:"rate,omitempty"`                           // 优惠力度
	Plus                  string `json:"plus,omitempty" codec:"plus,omitempty"`                           // 加价金额
	OrderModeDesc         string `json:"order_mode_desc,omitempty" codec:"order_mode_desc,omitempty"`     // 订单规则描述
	TokenUseNum           uint   `json:"token_use_num,omitempty" codec:"token_use_num,omitempty"`         // 用户使用令牌的次数
	UserPins              string `json:"user_pins,omitempty" codec:"user_pins,omitempty"`                 // 令牌用户列表
	PromoAreaType         uint8  `json:"promo_area_type,omitempty" codec:"promo_area_type,omitempty"`     // 促销区域类型： 2 白名单 3 黑名单
	PromoAreas            string `json:"promo_areas,omitempty" codec:"promo_areas,omitempty"`             // 促销区域列表（英文分号分隔）
}

type PromotionUserList struct {
	Pin       string `json:"pin" codec:"pin"`
	BeginTime string `json:"begin_time" codec:"begin_time"`
	EndTime   string `json:"end_time" codec:"end_time"`
	Created   string `json:"created" codec:"created"`
}

type PromotionSkuList struct {
	PromoSkuId  uint64 `json:"promo_sku_id" codec:"promo_sku_id"`
	WareId      uint64 `json:"ware_id" codec:"ware_id"`
	SkuId       uint64 `json:"sku_id" codec:"sku_id"`
	SkuName     string `json:"sku_name" codec:"sku_name"`
	BindType    uint8  `json:"bind_type" codec:"bind_type"`            // sku绑定类型（1、主商品，2、赠品，4、附件）
	JdPrice     string `json:"jd_price" codec:"jd_price"`              // 京东价
	PromoPrice  string `json:"promo_price" codec:"promo_price"`        // 促销价
	ItemNum     string `json:"item_num" codec:"item_num"`              // 货号
	LimitNum    uint   `json:"limit_num" codec:"limit_num"`            // 限购数量
	SkuStatus   uint8  `json:"sku_status" codec:"sku_status"`          // sku生效状态（0、未生效，1、已生效）
	Seq         uint64 `json:"seq" codec:"seq"`                        // 序号
	Display     uint8  `json:"display" codec:"display"`                // 是否展示
	IsNeedToBuy uint8  `json:"is_need_to_buy" codec: "is_need_to_buy"` // 是否必买
	Created     uint64 `json:"created" codec:"created"`
	Modified    uint64 `json:"modified" codec:"modified"`
	Rfld        uint64 `json:"rfId" codec:"rfId"`
	Logo        string `json:"logo,omitempty" codec:"logo,omitempty"`
	StockNum    int    `json:"stock_num,omitempty" codec:"stock_num,omitempty"`
	WareTitle   string `json:"ware_title,omitempty" codec:"ware_title,omitempty"`
}
