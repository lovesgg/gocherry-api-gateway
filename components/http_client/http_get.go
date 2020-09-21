package http_client

import (
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/proxy/enum"
	"io/ioutil"
	"net/http"
	"time"
)

/**
http get
*/
func Get(url string, timeRequest time.Duration, headers []string, ctx context.Context) (res string, statusCode int, errMsg string) {
	client := &http.Client{
		Timeout: timeRequest * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)

	for _, header := range headers {
		req.Header.Add(header, ctx.GetHeader(header))
	}

	if err != nil {
		return "", enum.STATUS_CODE_FAILED, "网络请求请求有误..."
	}
	response, resErr := client.Do(req)
	defer response.Body.Close()

	if resErr != nil {
		return "", enum.STATUS_CODE_FAILED, "网络请求状态码有误..."
	}

	result, _ := ioutil.ReadAll(response.Body)
	content := string(result)

	return content, enum.STATUS_CODE_OK, "ok"
}
