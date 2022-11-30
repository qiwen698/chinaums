package chinaums

import (
	"encoding/json"
	"log"
	"time"
)

type WxApp struct {
}

func (wx *WxApp) CreatePayment(conf UmsConfig, order OrderInfo) (res WxAppCreatePaymentResponse, err error) {
	nowTime := time.Now()
	app_param := make(map[string]interface{})
	app_param["notifyUrl"] = conf.NotifyUrl
	app_param["showUrl"] = conf.ShowUrl
	app_param["requestTimestamp"] = nowTime.Format("2006-01-02 15:04:05")
	app_param["mid"] = conf.Mid
	app_param["tid"] = conf.Tid
	app_param["instMid"] = "APPDEFAULT" //业务类型
	app_param["tradeType"] = "MINI"
	app_param["merOrderId"] = order.MerOrderId
	app_param["totalAmount"] = order.TotalAmount * 100 //
	app_param["subAppId"] = conf.SubAppId
	app_param["subOpenId"] = order.SubOpenId
	postData, _ := json.Marshal(app_param)
	authorzation := Authorization(conf, postData, nowTime)
	apiUrl := conf.ApiUrl + "/wx/unified-order" //请求的方法
	if conf.WriteLog {
		log.Printf("apiUrl:%v", apiUrl)
		log.Printf("postData:%v", string(postData))
		log.Printf("authorzation:%v", authorzation)
	}
	body, err := Request(apiUrl, postData, authorzation)
	if err != nil {
		log.Printf("err:%v", err)
		return res, err
	}
	if conf.WriteLog {
		log.Fatal(string(body))
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Printf("err:%v", err)
		return res, err
	}
	return res, err

}

func (wx WxApp) QueryPayment(conf UmsConfig, order OrderInfo) (res WxAppQueryPaymentResponse, err error) {
	nowTime := time.Now()
	app_param := make(map[string]interface{})
	app_param["requestTimestamp"] = nowTime.Format("2006-01-02 15:04:05")
	app_param["mid"] = conf.Mid
	app_param["tid"] = conf.Tid
	app_param["instMid"] = "APPDEFAULT" //业务类型
	app_param["merOrderId"] = order.MerOrderId
	postData, _ := json.Marshal(app_param)
	authorzation := Authorization(conf, postData, nowTime)
	apiUrl := conf.ApiUrl + "/query" //请求的方法
	if conf.WriteLog {
		log.Printf("apiUrl:%v", apiUrl)
		log.Printf("postData:%v", string(postData))
		log.Printf("authorzation:%v", authorzation)
	}
	body, err := Request(apiUrl, postData, authorzation)
	if err != nil {
		log.Printf("err:%v", err)
		return res, err
	}
	if conf.WriteLog {
		log.Fatal(string(body))
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Printf("err:%v", err)
		return res, err
	}
	return res, err
}
