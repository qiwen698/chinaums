package common

import (
	"crypto/sha256"
	"fmt"
	"log"
	"sort"
	"strings"
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
	//fmt.Println(bufSignSrc)
	//hash->sha256
	hash := sha256.New()
	hash.Write([]byte(bufSignSrc))
	sh := fmt.Sprintf("%x", hash.Sum(nil))
	//fmt.Println(sh)
	sh = strings.ToUpper(sh)
	//fmt.Printf("sh:%v,sign:%v\n", sh, params.GetString("sign"))
	return sh == params.GetString("sign")
}
