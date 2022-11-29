package chinaums

import (
	"time"
)

import "testing"

func TestCreatePayment(t *testing.T) {
	conf := UmsConfig{}
	order := OrderInfo{TotalAmount: 1, MerOrderId: time.Now().Format("20060102150405"), OrderDesc: "测试测试", SubOpenId: "OP-" + time.Now().Format("20060102150405")}
	params, err := new(WxApp).CreatePayment(conf, order)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("params:%#v", params)
}
