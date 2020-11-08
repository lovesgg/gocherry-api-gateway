package common_enum

const (
	ETCD_KEYS_APP_LIST                = "/keyAppListAll/"       //应用列表前缀key
	ETCD_KEYS_APP_CLUSTER_LIST        = "/keyClusterListAll/"   //app->服务 前缀key
	ETCD_KEYS_APP_CLUSTER_SERVER_LIST = "/keyClusterServerAll/" //app->服务->节点 前缀key 存储详情
	ETCD_KEYS_APP_API_LIST            = "/keyAppApiListAll/"    //app->api
	ETCD_KEYS_APP_SERVER_REGISTER     = "/keyServiceRegisterAll/"  //服务发现唯一新key 只存储ip
)
