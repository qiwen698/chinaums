package chinaums

type WxAppCreatePaymentResponse struct {
	MerName        string `json:"merName"`
	Mid            string `json:"mid"`
	MiniPayRequest struct {
		Package   string `json:"package"`
		Appid     string `json:"appid"`
		Sign      string `json:"sign"`
		Partnerid string `json:"partnerid"`
		Prepayid  string `json:"prepayid"`
		Noncestr  string `json:"noncestr"`
		Timestamp string `json:"timestamp"`
	} `json:"miniPayRequest"`
	SettleRefID       string `json:"settleRefId"`
	Tid               string `json:"tid"`
	TotalAmount       int    `json:"totalAmount"`
	QrCode            string `json:"qrCode"`
	TargetMid         string `json:"targetMid"`
	ResponseTimestamp string `json:"responseTimestamp"`
	ErrCode           string `json:"errCode"`
	ErrMsg            string `json:"errMsg"`
	PrepayID          string `json:"prepayId"`
	TargetStatus      string `json:"targetStatus"`
	SeqID             string `json:"seqId"`
	MerOrderID        string `json:"merOrderId"`
	Status            string `json:"status"`
	TargetSys         string `json:"targetSys"`
}

type WxAppQueryPaymentResponse struct {
	BuyerUsername     string `json:"buyerUsername"`
	PayTime           string `json:"payTime"`
	SeqID             string `json:"seqId"`
	InvoiceAmount     int    `json:"invoiceAmount"`
	SettleDate        string `json:"settleDate"`
	BuyerID           string `json:"buyerId"`
	TotalAmount       int    `json:"totalAmount"`
	CouponAmount      int    `json:"couponAmount"`
	BuyerPayAmount    int    `json:"buyerPayAmount"`
	TargetOrderID     string `json:"targetOrderId"`
	MerOrderID        string `json:"merOrderId"`
	Status            string `json:"status"`
	TargetSys         string `json:"targetSys"`
	MerName           string `json:"merName"`
	Mid               string `json:"mid"`
	Tid               string `json:"tid"`
	InstMid           string `json:"instMid"`
	ResponseTimestamp string `json:"responseTimestamp"`
	ErrCode           string `json:"errCode"`
	ErrMsg            string `json:"errMsg"`
}
