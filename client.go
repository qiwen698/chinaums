package chinaums

import (
	"errors"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

func init() {
	InitLog()
}
func InitLog() {
	logName := "chinaums" + time.Now().Format("2006-01-02") + ".log"
	logDir := "./log/"
	if runtime.GOOS == "windows" {
		logDir = ".\\log\\"
	}
	_, err := os.Stat(logDir)
	if nil != err {
		os.MkdirAll(logDir, os.ModePerm)
	}
	file, e := os.OpenFile(logDir+logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if e != nil {
		log.Panic("日志文件打开异常")
	}
	log.SetOutput(file)
}

var GlobalConfig UmsConfig

type NetPay struct {
	Config    UmsConfig `json:"config"`
	Pay       UmsPay    `json:"pay"`
	OrderInfo OrderInfo `json:"orderInfo"`
}

func NewClient(payType string, isPro bool, conf UmsConfig, orders ...OrderOption) (pay NetPay, err error) {
	if !isPro {
		conf.ApiUrl = "https://test-api-open.chinaums.com/v1/netpay"
	} else {
		conf.ApiUrl = "https://api-mop.chinaums.com/v1/netpay"
	}
	if conf.Mid == "" || conf.Tid == "" || conf.AppId == "" || conf.AppSecret == "" {
		log.Print("配置信息不能为空")
		return pay, errors.New("配置信息不能为空")
	}
	if payType == WxAppPay && conf.SubAppId == "" {
		log.Print("微信子商户信息不能为空")
		return pay, errors.New("微信子商户信息不能为空")
	}
	pay.Config = conf
	GlobalConfig = conf

	for _, fn := range orders {
		fn(&pay.OrderInfo)
	}
	pay.Pay, err = GetNetPayInstance(payType)
	return pay, nil
}

func GetNetPayInstance(mType string) (UmsPay, error) {
	var pay UmsPay
	switch strings.ToLower(mType) {
	case WxAppPay:
		pay = &WxApp{}
		return pay, nil
	default:
		return nil, errors.New("该支付方式暂未实现")

	}
}
