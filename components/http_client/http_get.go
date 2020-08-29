package http_client

import (
	"gocherry-api-gateway/proxy/enum"
	"io/ioutil"
	"net/http"
	"time"
)

/**
http get
*/
func Get(url string, timeRequest time.Duration) (res string, statusCode int, errMsg string) {
	timeout := time.Duration(timeRequest * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", enum.STATUS_CODE_FAILED, "网络请求请求有误..."
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return "", enum.STATUS_CODE_FAILED, "网络状态码有误"
	}

	resBody := string(body)

	return resBody, enum.STATUS_CODE_OK, "ok"
}
