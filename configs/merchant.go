package configs

type UmsConfig struct {
	AppId     string `json:"app_id"`
	AppKey    string `json:"app_key"`
	ApiUrl    string `json:"api_url"`
	Mid       string `json:"mid"` //商户号
	Tid       string `json:"tid"` //终端号
	ReturnUrl string `json:"return_url"`
	NotifyUrl string `json:"notify_url"`
	RefundUrl string `json:"refund_url"`
}
