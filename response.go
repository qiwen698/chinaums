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
	PayTime           string `json:"payTime"`
	BuyerCashPayAmt   int    `json:"buyerCashPayAmt"`
	ConnectSys        string `json:"connectSys"`
	ErrMsg            string `json:"errMsg"`
	MerName           string `json:"merName"`
	Mid               string `json:"mid"`
	InvoiceAmount     int    `json:"invoiceAmount"`
	SettleDate        string `json:"settleDate"`
	BillFunds         string `json:"billFunds"`
	BuyerID           string `json:"buyerId"`
	Tid               string `json:"tid"`
	ReceiptAmount     int    `json:"receiptAmount"`
	CouponAmount      int    `json:"couponAmount"`
	TargetMid         string `json:"targetMid"`
	CardAttr          string `json:"cardAttr"`
	TargetOrderID     string `json:"targetOrderId"`
	BillFundsDesc     string `json:"billFundsDesc"`
	SubBuyerID        string `json:"subBuyerId"`
	TargetStatus      string `json:"targetStatus"`
	SeqID             string `json:"seqId"`
	MerOrderID        string `json:"merOrderId"`
	RefundAmount      int    `json:"refundAmount"`
	TargetSys         string `json:"targetSys"`
	BankInfo          string `json:"bankInfo"`
	DelegatedFlag     string `json:"delegatedFlag"`
	SettleRefID       string `json:"settleRefId"`
	TotalAmount       int    `json:"totalAmount"`
	ChnlCost          string `json:"chnlCost"`
	ResponseTimestamp string `json:"responseTimestamp"`
	ErrCode           string `json:"errCode"`
	BuyerPayAmount    int    `json:"buyerPayAmount"`
	Status            string `json:"status"`
}
