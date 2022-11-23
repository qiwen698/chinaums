package common

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/qiwen698/chinaums/configs"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

var appSecret = ""

func ValidSign(params Params) bool {
	if params.GetString("sign") == "" {
		log.Printf("sign is null:%v\n", params.GetString("sign"))
		return false
	}
	// 得到各个key
	var keys []string
	var bufSignSrc string
	for key := range params {
		keys = append(keys, key)
	}
	// 给key排序，从小到大
	sort.Strings(keys)
	for _, v := range keys {
		if v == "sign" {
			continue
		}
		bufSignSrc = bufSignSrc + v + "=" + params[v] + "&"
	}
	//去除字符串首尾处的空白字符（或者其他字符）
	bufSignSrc = bufSignSrc[0 : len(bufSignSrc)-1]
	bufSignSrc = bufSignSrc + appSecret
	hash := sha256.New()
	hash.Write([]byte(bufSignSrc))
	sh := fmt.Sprintf("%x", hash.Sum(nil))
	sh = strings.ToUpper(sh)
	return sh == params.GetString("sign")
}

func bin2hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

func genHMAC256(ciphertext, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(ciphertext)
	hmac := mac.Sum(nil)
	return hmac
}
func Authorization(postJsonStr string, t time.Time) string {
	var conf = configs.UmsConfig{}
	fmt.Printf("conf:%#v", conf)
	timeYmd := t.Format("20060102150405")

	m := md5.New()
	m.Write([]byte(time.Microsecond.String()))
	random := hex.EncodeToString(m.Sum(nil))

	hash := sha256.New()
	hash.Write([]byte(postJsonStr))
	sh := fmt.Sprintf("%x", hash.Sum(nil))
	signBody, _ := bin2hex(sh)

	appKey := []byte(conf.AppKey)
	appId := conf.AppId
	b := strings.Builder{}
	b.WriteString(appId)
	b.WriteString(timeYmd)
	b.WriteString(random)
	b.WriteString(signBody)
	data := b.String()
	hmac := genHMAC256([]byte(data), appKey)
	sign := base64.StdEncoding.EncodeToString(hmac)
	a := strings.Builder{}
	a.WriteString("OPEN-BODY-SIG AppId=\"")
	a.WriteString(appId)
	a.WriteString("\",Timestamp=\"")
	a.WriteString(timeYmd)
	a.WriteString("\",Nonce=\"")
	a.WriteString(random)
	a.WriteString("\",Signature=\"")
	a.WriteString(sign)
	a.WriteString("\"")
	return a.String()
}
