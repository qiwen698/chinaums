package chinaums

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ValidSign(conf UmsConfig, params Params) bool {
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
	b := strings.Builder{}
	for _, v := range keys {
		if v == "sign" || v == "" {
			continue
		}
		b.WriteString(v)
		b.WriteString("=")
		b.WriteString(params[v])
		b.WriteString("&")
		//bufSignSrc = bufSignSrc + v + "=" + params[v] + "&"
	}
	//去除字符串首尾处的空白字符（或者其他字符）
	//bufSignSrc = bufSignSrc[0 : len(bufSignSrc)-1]
	//bufSignSrc = bufSignSrc + conf.AppSecret
	bufSignSrc = strings.TrimRight(bufSignSrc, "&")
	b.WriteString(conf.AppSecret)
	bufSignSrc = b.String()

	hash := sha256.New()
	hash.Write([]byte(bufSignSrc))
	sh := fmt.Sprintf("%x", hash.Sum(nil))
	sh = strings.ToUpper(sh)
	return sh == params.GetString("sign")
}

func genHMAC256(ciphertext, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(ciphertext)
	hmac := mac.Sum(nil)
	return hmac
}
func Authorization(conf UmsConfig, postData []byte, t time.Time) string {
	timeYmd := t.Format("20060102150405")
	m := md5.New()
	m.Write([]byte(strconv.FormatInt(time.Now().UnixMicro(), 10)))
	random := hex.EncodeToString(m.Sum(nil))
	hash := sha256.New()
	hash.Write(postData)
	signBody := fmt.Sprintf("%x", hash.Sum(nil))
	appKey := []byte(conf.AppKey)
	appId := conf.AppId
	b := bytes.Buffer{}
	b.WriteString(appId)
	b.WriteString(timeYmd)
	b.WriteString(random)
	b.WriteString(signBody)
	hmac := genHMAC256(b.Bytes(), appKey)
	sign := base64.StdEncoding.EncodeToString(hmac)
	s := strings.Builder{}
	s.WriteString("OPEN-BODY-SIG AppId=\"")
	s.WriteString(appId)
	s.WriteString("\",Timestamp=\"")
	s.WriteString(timeYmd)
	s.WriteString("\",Nonce=\"")
	s.WriteString(random)
	s.WriteString("\",Signature=\"")
	s.WriteString(sign)
	s.WriteString("\"")
	return s.String()
}
