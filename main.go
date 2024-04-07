package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	var port string = "8080"
	for _, value := range os.Args {
		port = value
	}

	// 创建引擎对象
	engine := gin.Default()
	// 创建控制器
	engine.Handle("GET", "/", func(context *gin.Context) {
		// 获取请求者的IP地址
		var ip = context.ClientIP()
		// 返回响应数据
		context.Writer.Write([]byte(ip))
	})
	// 开启服务器
	engine.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
