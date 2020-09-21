package models

import "time"

//api存储相关的数据 admin和proxy公用这组数据 谨慎修改哦
type Api struct {
	BaseApiName        string        `json:"base_api_name"`
	BaseApiUrl         string        `json:"base_api_url"`
	BaseRedirectUrl    string        `json:"base_redirect_url"`
	BaseApiRequestType []string      `json:"base_api_request_type"`
	BaseApiStatus      bool          `json:"base_api_status"`
	BaseClusterName    string        `json:"base_cluster_name"`
	CacheSave          time.Duration `json:"cache_save"`
	LimitRequest       int64           `json:"limit_request"`
	ReduceLevel        bool          `json:"reduce_level"`
	TimeOut            time.Duration `json:"time_out"`
	IpBlack            string        `json:"ip_black"`
	UserAuth           bool          `json:"user_auth"`
	JwtAuth            bool          `json:"jwt_auth"`
	WhiteList          string        `json:"white_list"`
	WhiteListCheck     bool          `json:"white_list_check"`
	UpdateTime         string        `json:"update_time"`
	RetryRequest       bool          `json:"retry_request"`
	HeaderForms        string        `json:"header_forms"`
}
