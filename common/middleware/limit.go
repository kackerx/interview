package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Client 定义每个客户端的限流器
type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type RateLimiter struct {
	clients map[string]*Client
	mutex   sync.Mutex
	r       rate.Limit
	b       int
}

// NewRateLimiter 创建一个新的限流器
func NewRateLimiter() *RateLimiter {
	rl := &RateLimiter{
		clients: make(map[string]*Client),
		r:       1,
		b:       2,
	}

	// 启动清理协程
	go rl.cleanupClients()

	return rl
}

func (rl *RateLimiter) GetLimiter(clientID string) *rate.Limiter {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	client, exists := rl.clients[clientID]
	if !exists {
		limiter := rate.NewLimiter(rl.r, rl.b)
		rl.clients[clientID] = &Client{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}

	client.lastSeen = time.Now()
	return client.limiter
}

func (rl *RateLimiter) cleanupClients() {
	for {
		time.Sleep(time.Minute)
		rl.mutex.Lock()
		for clientID, client := range rl.clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(rl.clients, clientID)
			}
		}
		rl.mutex.Unlock()
	}
}

// RateLimitMiddleware 返回一个 Gin 中间件，使用 golang.org/x/time/rate 进行限流
func RateLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		limiter := rl.GetLimiter(clientIP)

		if limiter.Allow() {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too Many Requests",
			})
			return
		}
	}
}
