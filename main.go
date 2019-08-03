package main

import (
	"fmt"
	"web/app/router"
	"web/library"
)

func main() {
	defer library.DB.Close()

	config := library.NewConfig("")
	httpPort := config.Get("http.port")

	port := httpPort.(string)

	//初始化router
	r := router.Router()
	r.Run(port)
	fmt.Println("启动HTTP服务器:", port)
}
