package main

import "github.com/lupc/go-http_proxy/proxy"

func main() {

	// 转发请求
	//proxy.VerOption.ApiHost = "http://www.baidu.com"
	// VerOption.ApiHost = "http://112.35.1.155:1992"
	// VerOption.Port = "9009"

	proxy.LoadConfig("config.yaml")

	// 启动服务
	proxy.Proxy(proxy.VerOption)
}
