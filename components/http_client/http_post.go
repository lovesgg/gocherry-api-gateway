package http_client

import (
	"bytes"
	"encoding/json"
	"gocherry-api-gateway/proxy/enum"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/kataras/iris/context"
)

/**
http post
*/
func Post(url string, data interface{}, timeRequest time.Duration, headers []string, ctx context.Context) (res string, statusCode int, errMsg string) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")

	for _, header := range headers {
		req.Header.Add(header, ctx.GetHeader(header))
	}

	if err != nil {
		return "", enum.STATUS_CODE_FAILED, "请求下游接口出错啦"
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: timeRequest * time.Second}
	resp, errorReq := client.Do(req)
	if errorReq != nil {
		return "", enum.STATUS_CODE_FAILED, "请求接口出错啦"
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)
	return content, enum.STATUS_CODE_OK, "ok"
}
