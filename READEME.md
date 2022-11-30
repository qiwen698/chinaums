```go
package main

import (
	"fmt"
	"github.com/qiwen698/chinaums"
)

func main() {
	QueryPayment()  //查询
	CreatePayment() //发起支付
}

func QueryPayment() {
	conf := chinaums.UmsConfig{}
	conf.Mid = "8xxxxxxxAP"
	conf.Tid = "KxxxxxxT"
	conf.AppId = "8a8xxxxxxxxb64037a"
	conf.SubAppId = "wxee35a7xxxx"
	conf.AppKey = "c8xxxxx59c0bb"
	conf.AppSecret = "xxxxx"
	conf.NotifyUrl = ""
	conf.ShowUrl = ""
	conf.WriteLog = true
	client, err := chinaums.NewClient(chinaums.WxAppPay, true, conf,
		chinaums.SetMerOrderId("b3dac2294c3ba1504864d65ff311"),
	)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	params, err := client.Pay.QueryPayment(client.Config, client.OrderInfo)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Printf("params:%#v", params)
}

func CreatePayment() {
	conf := chinaums.UmsConfig{}
	conf.Mid = "8xxxxxxxAP"
	conf.Tid = "KxxxxxxT"
	conf.AppId = "8a8xxxxxxxxb64037a"
	conf.SubAppId = "wxee35a7xxxx"
	conf.AppKey = "c8xxxxx59c0bb"
	conf.AppSecret = "xxxxx"
	conf.NotifyUrl = ""
	conf.ShowUrl = ""
	conf.WriteLog = true
	client, err := NewClient(WxAppPay, true, conf,
		chinaums.SetTotalAmount(10),
		chinaums.SetMerOrderId("11WE32234234"),
		chinaums.SetOrderDesc("测试"),
		chinaums.SetSubOpenId("OPEN-ID234234234"),
		chinaums.SetNotifyUrl(""),
		chinaums.SetShowUrl(""),
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

```