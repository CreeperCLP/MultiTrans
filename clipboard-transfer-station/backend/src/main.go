package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // 创建一个新的Gin路由引擎
    r := gin.Default()

    // 设置路由
    r.GET("/health", healthCheck)

    // 启动服务器
    if err := r.Run(":8080"); err != nil {
        panic("Failed to start server: " + err.Error())
    }
}

// healthCheck 返回服务器健康状态
func healthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}