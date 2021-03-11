# shinedone业务层框架

## 安装protoc和grpc
```
1.到下面网站，下载protoc可执行文件，文件名 protoc-x.xx.x-win64.zip
https://github.com/protocolbuffers/protobuf/releases
2.执行以下指令
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
```

## 更换module名称
```
修改go.mod文件里面的module，修改为 github.com/shinedone/srv-微服务名称,比如github.com/shinedone/srv-org
全局替换module名
```

## 如何开发grpc
1. 编写proto文件，类似以下enterprise.proto
``` protobuf
syntax = "proto3";
package enterprise;

import "google/protobuf/timestamp.proto";

option go_package = "github.com/shinedone/srv-framework/pkg/adm/grpc/enterprise";


service EnterpriseSrv{
  rpc GetByEnterpriseId (GetByEnterpriseIdReq) returns (GetByEnterpriseIdResp);
  rpc GetByEnterpriseId2 (GetByEnterpriseIdReq) returns (GetByEnterpriseIdResp);
}

message GetByEnterpriseIdReq{
  string enterprise_id = 1;
}

message GetByEnterpriseIdResp{
  string enterprise_id = 1;
  string name = 2;
  google.protobuf.Timestamp created_at = 3;

}

```
2. 创建gen.go
```go
package enterprise

//go:generate protoc --go_opt=paths=source_relative --go_out=plugins=grpc:.  *.proto

```
3. 执行gen.go文件里面的 go:generate

