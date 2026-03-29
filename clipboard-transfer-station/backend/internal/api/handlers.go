package api


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"clipboard-transfer-station/backend/internal/service"
	"clipboard-transfer-station/backend/internal/ws"
)

var (
	clipboardManager = service.NewClipboardManager()
	hub             = ws.NewHub()
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/clipboard", getClipboard)
	router.POST("/api/clipboard", setClipboard)
	router.GET("/ws", ws.ServeWs(hub, clipboardManager))
}


// getClipboard 处理 GET /api/clipboard 请求，返回当前剪贴板内容
func getClipboard(c *gin.Context) {
       data := clipboardManager.GetContent()
       if data.Content == "" {
	       c.JSON(http.StatusNotFound, gin.H{"error": "No clipboard content"})
	       return
       }
       c.JSON(http.StatusOK, data)
}


// setClipboard 处理 POST /api/clipboard 请求，更新剪贴板内容
func setClipboard(c *gin.Context) {
       var req struct {
	       Content string `json:"content" binding:"required"`
       }
       if err := c.ShouldBindJSON(&req); err != nil {
	       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	       return
       }
       clipboardManager.UpdateContent(req.Content)
       c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
