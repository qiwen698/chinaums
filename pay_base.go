package chinaums

import (
	"strconv"
)

type UmsPay interface {
	CreatePayment(conf UmsConfig, order OrderInfo) (WxAppCreatePaymentResponse, error)
	QueryPayment(conf UmsConfig, order OrderInfo) (WxAppQueryPaymentResponse, error)
}

func Request(url string, postData []byte, authorization string) ([]byte, error) {
	heads := make(map[string]string)
	heads["AUTHORIZATION"] = authorization
	heads["Accept"] = "application/json"
	heads["Content-Type"] = "application/json;charset=utf-8"
	heads["Content-Length"] = strconv.Itoa(len(postData))
	body, err := Post(url, postData, heads, false)
	return body, err
}
