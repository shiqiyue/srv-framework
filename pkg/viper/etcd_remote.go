package viper

import (
	"errors"
	"github.com/shinedone/srv-framework/internal/global/instance"
	"github.com/shinedone/srv-framework/pkg/assert"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
)

var (
	viperEtcdEndpointConfig string = "viper.etcd.endpoints"
	viperEtcdPathConfig     string = "viper.etcd.path"
)

// 使用etcd进行配置
func ConfigWithEtcd(onValueChange func(v *viper.Viper)) {
	valueChangeCh := make(chan bool)
	viper.RemoteConfig = &EtcdRemoteConfig{ValueChangeCh: valueChangeCh}
	v := viper.New()
	// 设置v变量到全局变量中
	instance.SetViper(v)

	// 从命令行中获取配置中心默认值
	pflag.String(viperEtcdEndpointConfig, "192.168.3.132:2379", "etcd配置中心地址")
	pflag.String(viperEtcdPathConfig, "/dev/app/app.yaml", "配置文件在配置中心的路径")
	pflag.Parse()
	err := v.BindPFlags(pflag.CommandLine)
	assert.Nil(err, errors.New("bind pflag fail"))

	// 从环境变量取值
	err = v.BindEnv(viperEtcdEndpointConfig)
	assert.Nil(err, errors.New("bind env fail"))
	err = v.BindEnv(viperEtcdPathConfig)
	assert.Nil(err, errors.New("bind env fail"))
	v.AutomaticEnv()

	// 从配置中心取值
	err = v.AddRemoteProvider("etcd", v.GetString(viperEtcdEndpointConfig), v.GetString(viperEtcdPathConfig))
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
