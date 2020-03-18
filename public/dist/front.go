package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Static("/", "./")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))


	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":3000")

}
