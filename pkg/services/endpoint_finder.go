package services

import (
	"fmt"
	"github.com/spf13/viper"
)

// 端点查找
type EndpointFinder struct {
	v *viper.Viper
}

// 获取gql端点
func (f *EndpointFinder) GetGqlServerEndpoint(serverName, moduleName string) string {
	return f.getEndpoint(serverName, moduleName, "gql")
}

// 获取rest端点
func (f EndpointFinder) GetRestServerEndpoint(serverName, moduleName string) string {
	return f.getEndpoint(serverName, moduleName, "restful")
}

// 获取grpc端点
func (f EndpointFinder) GetGrpcServerEndpoint(serverName, moduleName string) string {
	return f.getEndpoint(serverName, moduleName, "grpc")
}

// 获取端点
func (f *EndpointFinder) getEndpoint(serverName, moduleName, protocol string) string {
	endpoint := f.v.GetString(fmt.Sprintf("endpoint.%s.%s.%s.url", serverName, moduleName, protocol))
	return endpoint
}
