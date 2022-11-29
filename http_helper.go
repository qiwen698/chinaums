package chinaums

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Get
func Get(apiUrl string, parm map[string]string, header map[string]string, isHttps bool) ([]byte, error) {

	if len(parm) > 0 {
		apiUrl = fmt.Sprintf("%s%s", apiUrl, "?")
		p := ""
		for k, v := range parm {
			if p == "" {
				p = fmt.Sprintf("%s=%s", k, v)
			} else {
				p = fmt.Sprintf("%s&%s=%s", p, k, v)
			}
		}
		apiUrl = fmt.Sprintf("%s%s", apiUrl, p)
	}
	client := &http.Client{}

	if isHttps {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	reqest, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		reqest.Header.Set(k, v)
	}

	response, err := client.Do(reqest)
	if nil != err {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}

// post
func Post(apiUrl string, data []byte, header map[string]string, isHttps bool) ([]byte, error) {

	client := &http.Client{Timeout: 5 * time.Second}

	if isHttps {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	reqest, err := http.NewRequest("POST", apiUrl, bytes.NewReader(data))
	if err != nil {
		log.Printf("post error:%s", err)
		return []byte(""), err
	}
	for k, v := range header {
		reqest.Header.Set(k, v)
	}
	response, err := client.Do(reqest)
	if nil != err {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}

// post
func PostMap(apiUrl string, parm map[string]string, header map[string]string, isHttps bool) ([]byte, error) {

	data := url.Values{}
	for k, v := range parm {
		data.Set(k, v)
	}

	reqParams := ioutil.NopCloser(strings.NewReader(data.Encode()))
	client := &http.Client{}

	if isHttps {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	reqest, _ := http.NewRequest("POST", apiUrl, reqParams)

	for k, v := range header {
		reqest.Header.Set(k, v)
	}

	response, err := client.Do(reqest)
	if nil != err {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}

// 获取远程ip
func GetRemoteIP(r *http.Request) string {
	addr := r.Header.Get("Remote_addr")
	if addr == "" {
		addr = r.RemoteAddr
	}

	return strings.Split(addr, ":")[0]
}

// ParseQuery 参数解析，兼容部分字段有encode部分没encode的情况
func ParseQuery(query string) (m url.Values, err error) {
	m = make(url.Values)
	for query != "" {
		key := query
		if i := strings.IndexAny(key, "&"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value1, err1 := url.QueryUnescape(value)
		if err1 != nil {
			if !strings.HasPrefix(err1.Error(), "invalid URL escape") {
				if err == nil {
					err = err1
				}
				continue
			}
			m[key] = append(m[key], value)
		} else {
			m[key] = append(m[key], value1)
		}
	}
	return m, err
}

// 获取客户端地址
func GetIPAdress(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		if r.Header.Get(h) != "" {
			addresses := strings.Split(r.Header.Get(h), ",")
			for i := len(addresses) - 1; i >= 0; i-- {
				ip := strings.TrimSpace(addresses[i])
				realIP := net.ParseIP(ip)
				if !realIP.IsGlobalUnicast() {
					continue
				}
				return ip
			}
		}
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
