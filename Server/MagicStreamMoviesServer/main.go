package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxfeizi04-cloud/MagicStreamMovies/Server/MagicStreamMoviesServer/routes"
)

// main 是程序入口函数。
// 这里负责完成三件事：
// 1. 创建 Gin 路由引擎。
// 2. 注册当前服务对外暴露的 HTTP 接口。
// 3. 启动 Web 服务并监听 8080 端口。
//
// 当前项目对外提供的接口如下：
// - GET /hello：用于快速测试服务是否启动成功。
// - GET /movies：查询全部电影数据。
// - GET /movie/:imdb_id：根据 imdb_id 查询单个电影详情。
func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	routes.SetupUnProtectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server")
	}
}
