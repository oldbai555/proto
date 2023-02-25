package impl

import (
	"fmt"
	"github.com/oldbai555/bgg/client/{{.ServerName}}"
	"github.com/oldbai555/bgg/webtool"
	"github.com/oldbai555/lbtool/log"
)

var lb *Tool

type Tool struct {
	*webtool.WebTool

	// 可以向下扩展其他的rpc服务
}

func StartServer() {
	var err error
	lb = &Tool{}
	lb.WebTool, err = webtool.NewWebTool(webtool.OptionWithOrm(

	), webtool.OptionWithRdb(), webtool.OptionWithStorage())

	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	// 初始化数据库工具类
	InitDbOrm()

	// 地址
	addr := fmt.Sprintf(":%d", lb.Sc.Port)
	// 1.监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", addr)
	// 2.实例化gRPC
	s := grpc.NewServer()
	// 3.在gRPC上注册微服务
	{{.ServerName}}.Register{{.UpperServerName}}Server(s, &{{.ServerName}}Server)
	// 4.启动服务端
	if err = s.Serve(listener); err != nil {
		log.Errorf("err is %v", err)
	}
}
