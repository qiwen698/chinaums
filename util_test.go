package chinaums

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAuthorization(t *testing.T) {
	conf := UmsConfig{}
	app_parm := make(map[string]string)
	app_parm["requestTimestamp"] = "2022-11-25 09:42:48"
	app_parm["merOrderId"] = "11WE20221125094248"
	app_parm["instMid"] = "APPDEFAULT"
	app_parm["mid"] = "TESTAAAP"
	app_parm["tid"] = "TEST11111"
	data, _ := json.Marshal(app_parm)
	authorization := Authorization(conf, data, time.Now())
	t.Logf("%v", string(data))
	t.Logf("%v", authorization)
}
