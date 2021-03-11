package viper

import (
	"errors"
	"github.com/shinedone/srv-framework/pkg/assert"
	"github.com/spf13/viper"
	"time"
)

// 使用etcd进行配置
func ConfigWithEtcd(onValueChange func(v *viper.Viper), endpoint, path string) {
	valueChangeCh := make(chan bool)
	viper.RemoteConfig = &EtcdRemoteConfig{ValueChangeCh: valueChangeCh}
	v := viper.New()

	// 从配置中心取值
	err := v.AddRemoteProvider("etcd", endpoint, path)
	assert.Nil(err, errors.New("add remote provider fail"))
	v.SetConfigType("yaml")
	err = v.ReadRemoteConfig()
	assert.Nil(err, errors.New("read remote config fail"))
	onValueChange(v)

	assert.Nil(err, errors.New("unmarshal config fail"))

	// 定时从配置中心获取值
	go func() {
		for true {
			// 等待远程配置变化
			<-valueChangeCh
			// 暂停1秒，等待viper处理完毕
			time.Sleep(time.Second)
			onValueChange(v)
		}
	}()
}
