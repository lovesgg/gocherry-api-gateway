package http_client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

/**
http post
*/
func Post(url string, data interface{}, timeRequest time.Duration) (res string, statusCode int, errMsg string) {
	jsonStr, _ := json.Marshal(data)
	//fmt.Println(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	if err != nil {

	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, errorReq := client.Do(req)
	if errorReq != nil {
		panic(errorReq)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content := string(result)
	return content, 1, ""
}
