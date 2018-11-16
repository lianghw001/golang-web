package main

import (
	"os"

	"github.com/lianghw001/golang-web/service"
	flag "github.com/spf13/pflag"
)

//PORT 是备用端口
const PORT string = "9090"

func main() {

	//返回当前进程的环境变量"PORT"的值,若变量没有定义时返回nil
	//使用当前"PORT",没有则使用备用端口
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	//若启动程序时，给了port的flag
	//则使用
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	//建立服务
	server := service.NewServer()

	//开始服务
	server.Run(":" + port)
}
