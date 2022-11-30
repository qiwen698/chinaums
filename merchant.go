package chinaums

type UmsConfig struct {
	AppId     string `json:"app_id"`
	SubAppId  string `json:"sub_app_id"` //微信子商户 appId
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	ApiUrl    string `json:"api_url"`
	Mid       string `json:"mid"` //商户号
	Tid       string `json:"tid"` //终端号
	NotifyUrl string `json:"notify_url"`
	RefundUrl string `json:"refund_url"`
	ShowUrl   string `json:"show_url"`  //订单展示页面
	WriteLog  bool   `json:"write_log"` //是否写入日志
}

type OrderInfo struct {
	TotalAmount int64  `json:"total_amount"`
	MerOrderId  string `json:"mer_order_id"`
	OrderDesc   string `json:"order_desc"`
	SubOpenId   string `json:"sub_open_id"`
	//跳转地址，订单有传值，则会覆盖配置信息
	NotifyUrl string `json:"notify_url"`
	RefundUrl string `json:"refund_url"`
	ShowUrl   string `json:"show_url"` //订单展示页面
}

type OrderOption func(order *OrderInfo) (o *OrderInfo)

func SetTotalAmount(amount int64) OrderOption {
	return func(order *OrderInfo) (o *OrderInfo) {
		order.TotalAmount = amount
		return order
	}
}

func SetMerOrderId(orderSn string) OrderOption {
	return func(order *OrderInfo) (o *OrderInfo) {
		order.MerOrderId = orderSn
		return order
	}
}
func SetOrderDesc(desc string) OrderOption {
	return func(order *OrderInfo) (o *OrderInfo) {
		order.OrderDesc = desc
		return order
	}
}

func SetSubOpenId(openId string) OrderOption {
	return func(order *OrderInfo) (o *OrderInfo) {
		order.SubOpenId = openId
		return order
	}
}
func SetNotifyUrl(notifyUrl string) OrderOption {
	return func(order *OrderInfo) (o *OrderInfo) {
		order.NotifyUrl = notifyUrl
		return order
	}
}

func SetShowUrl(showUrl string) OrderOption {
	return func(order *OrderInfo) (o *OrderInfo) {
		order.ShowUrl = showUrl
		return order
	}
}

func SetRefundUrl(refundUrl string) OrderOption {
	return func(order *OrderInfo) (o *OrderInfo) {
		order.RefundUrl = refundUrl
		return order
	}
}
