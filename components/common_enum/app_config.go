package common_enum

type Config struct {
	Admin  *Admin        `toml:"admin"`
	Proxy  *Proxy        `toml:"proxy"`
	Common *CommonConfig `toml:"common"`
}

type Admin struct {
	AppName string `toml:"app_name"`
	AppPort string `toml:"app_port"`
	LogDir  string `toml:"log_dir"`
}

type Proxy struct {
	AppName string `toml:"app_name"`
	AppPort string `toml:"app_port"`
	LogDir  string `toml:"log_dir"`
}

type CommonConfig struct {
	Test           string `toml:"test"`
	EtcdIp         string `toml:"etcd_ip"`
	EtcdPort       string `toml:"etcd_port"`
	RedisProxyIp   string `toml:"redis_proxy_ip"`
	RedisProxyPort string `toml:"redis_proxy_port"`
	RedisProxyPass string `toml:"redis_proxy_pass"`
	LogDir         string `toml:"log_dir"`
	MysqlHost      string `toml:"mysql_host"`
	MysqlPort      string `toml:"mysql_port"`
	MysqlUser      string `toml:"mysql_user"`
	MysqlPass      string `toml:"mysql_pass"`
	MysqlDb        string `toml:"mysql_db"`
	MysqlDebug     bool   `toml:"mysql_debug"`
}
