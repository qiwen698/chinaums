package netpay

type UmsPay interface {
	BuildChargeParams() (map[string]interface{}, error)
}
