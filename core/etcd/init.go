package etcd

import (
	"github.com/spf13/viper"
)

var etcdSlice []string

func init() {
	//获取配置的etcd节点
	etcdSlice = viper.GetStringSlice("etcd")
}
