package chinaums

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	conf := UmsConfig{}
	client, err := NewClient(WxAppPay, true, conf,
		SetTotalAmount(10),
		SetMerOrderId("11WE32234234"),
		SetOrderDesc("测试"),
		SetSubOpenId("OPEN-ID234234234"),
		SetNotifyUrl(""),
		SetShowUrl(""),
	)
	if err != nil {
		t.Logf("err:%v", err)
		return
	}
	params, err := client.Pay.CreatePayment(client.Config, client.OrderInfo)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	t.Logf("params:%#v", params)
}
