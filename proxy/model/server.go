package model

type Server struct {
	ClusterName string `json:"cluster_name"`
	AppName     string `json:"app_name"`
	ServerName  string `json:"server_name"`
	Ip          string `json:"ip"`
}
