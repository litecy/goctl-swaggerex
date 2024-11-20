module github.com/litecy/goctl-swaggerex

go 1.21

toolchain go1.22.4

replace github.com/zeromicro/goctl-swagger => github.com/litecy/goctl-swaggerex latest

require (
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/urfave/cli/v2 v2.11.0
	github.com/zeromicro/go-zero v1.6.5
	github.com/zeromicro/go-zero/tools/goctl v1.6.5
	golang.org/x/oauth2 v0.17.0
)
