grpc

```
 go get -u google.golang.org/grpc
 go get -u github.com/golang/protobuf/protoc-gen-go
```

安装protoc：

windows安装：
```
1. 直接下载安装包解压: https://github.com/protocolbuffers/protobuf/releases/
2. 将解压后的bin目录配置到环境变量 
3. 配置完成后 执行 `protoc --version` 正确显示版本即可
```
mac安装：
```
brew install protobuf
```

踩坑记录：
```
1. protoc 和 protoc-gen-go都要添加到环境变量中
```




本测试使用：
```
1. 在 \dailyPractice\grpc\server 目录下执行 protoc -I. --go_out=plugins=grpc:. proto/test.proto 将写好的proto文件编译成go代码
2. protoc --go_out=. --go_opt=paths=source_relative   --go-grpc_out=. --go-grpc_opt=paths=source_relative     ./pb/hello.proto
```

参考文档：
```
https://blog.51cto.com/u_16099242/9767201
https://blog.csdn.net/sinat_41672927/article/details/127206112
https://blog.csdn.net/qq_36268452/article/details/123359308
```