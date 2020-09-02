package etcd_client

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"gocherry-api-gateway/admin/services"
	"log"
	"time"
)

var etcdClient *clientv3.Client

func init() {
	etcdClient = GetClient()
}

func GetClient() *clientv3.Client {
	appConfig := services.GetAppConfig()
	ipPort := appConfig.Common.EtcdIp + ":" + appConfig.Common.EtcdPort

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{ipPort},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("etcd 连接失败")
	}

	return client
}

func PutKv(key string, value string) (*clientv3.PutResponse, error) {
	kv := clientv3.NewKV(etcdClient)
	ret, err := kv.Put(context.TODO(), key, string(value))
	return ret, err
}

func GetKv(key string) (*clientv3.GetResponse, error) {
	kv := clientv3.NewKV(etcdClient)
	ret, err := kv.Get(context.TODO(), key)
	return ret, err
}

func GetKvPrefix(key string) (*clientv3.GetResponse, error) {
	kv := clientv3.NewKV(etcdClient)
	ret, err := kv.Get(context.TODO(), key, clientv3.WithPrefix())
	return ret, err
}

func DelKvPrefix(prefix string) (*clientv3.DeleteResponse, error) {
	kv := clientv3.NewKV(etcdClient)
	ret, err := kv.Delete(context.TODO(), prefix, clientv3.WithPrefix())
	return ret, err
}

func DelKv(key string) (*clientv3.DeleteResponse, error) {
	kv := clientv3.NewKV(etcdClient)
	ret, err := kv.Delete(context.TODO(), key)
	return ret, err
}
