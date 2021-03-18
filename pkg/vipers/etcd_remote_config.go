package vipers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
)

type EtcdRemoteConfig struct {
	viper.RemoteProvider
	Username      string
	Password      string
	lock          *sync.Mutex
	client        *clientv3.Client
	Mutex         sync.Mutex
	ValueChangeCh chan bool
}

// 获取配置信息
func (c *EtcdRemoteConfig) Get(rp viper.RemoteProvider) (io.Reader, error) {
	c.RemoteProvider = rp

	return c.get()
}

// 监听配置信息
func (c *EtcdRemoteConfig) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	c.RemoteProvider = rp

	return c.get()
}

// 监听配置信息
func (c *EtcdRemoteConfig) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {
	c.RemoteProvider = rp

	rr := make(chan *viper.RemoteResponse)
	stop := make(chan bool)

	go func() {
		client, err := c.newClient()
		defer fmt.Println("watch return")
		if err != nil {
			panic(err)
		}
		ch := client.Watch(context.Background(), c.RemoteProvider.Path())
		for {
			select {
			case <-stop:
				return
			case res := <-ch:
				fmt.Println("watch value change")
				rr <- &viper.RemoteResponse{
					Value: res.Events[0].Kv.Value,
				}

			}
		}
	}()

	return rr, stop
}

// 新建etcd客户端
func (c *EtcdRemoteConfig) newClient() (*clientv3.Client, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	if c.client != nil {
		return c.client, nil
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(c.Endpoint(), ","),
		Username:    c.Username,
		Password:    c.Password,
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	c.client = client

	return client, nil
}

// 获取配置信息
func (c *EtcdRemoteConfig) get() (io.Reader, error) {
	client, err := c.newClient()

	if err != nil {
		return nil, err
	}
	// 1秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := client.Get(ctx, c.RemoteProvider.Path())
	cancel()

	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Kvs)

	return bytes.NewReader(resp.Kvs[0].Value), nil
}
