package service

import (
	"sync"
	"time"
)


// ClipboardData 表示剪贴板的内容和最后更新时间
type ClipboardData struct {
	Content     string `json:"clipboardContent"` // 剪贴板内容
	LastUpdated int64  `json:"lastUpdated"`      // 最后更新时间戳（毫秒）
}


// ClipboardManager 管理剪贴板状态的并发安全服务
type ClipboardManager struct {
	mu        sync.RWMutex           // 读写锁，保证并发安全
	state     ClipboardData          // 当前剪贴板状态
	Broadcast chan ClipboardData     // 广播通道，推送内容变更
}


// NewClipboardManager 创建并初始化剪贴板管理器
func NewClipboardManager() *ClipboardManager {
       return &ClipboardManager{
	       state:     ClipboardData{},
	       Broadcast: make(chan ClipboardData, 100),
       }
}


// GetContent 获取当前剪贴板内容（并发安全）
func (cm *ClipboardManager) GetContent() ClipboardData {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.state
}


// UpdateContent 更新剪贴板内容并广播变更（并发安全）
func (cm *ClipboardManager) UpdateContent(newContent string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.state.Content = newContent
	cm.state.LastUpdated = time.Now().UnixMilli()
	// 非阻塞推送到广播通道
	select {
	case cm.Broadcast <- cm.state:
	default:
	}
}
