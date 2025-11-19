package connect

import (
	// "net/http"
	// "net/http/httptest"
	"testing"
	// "time"

	// "github.com/stretchr/testify/assert"
	// "github.com/zeromicro/go-zero/core/logx"
	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {

	convey.Convey("Test Get", t, func() {
		// url := "baidu.com"
		url := "ppp/sadasd/sdada"
		got := Get(url)
		// 断言
		convey.ShouldBeFalse(got)
	})

	// // 设置测试日志
	// logx.Disable()

	// // 测试用例1: 正常响应200
	// t.Run("successful request", func(t *testing.T) {
	// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(http.StatusOK)
	// 	}))
	// 	defer server.Close()

	// 	result := Get(server.URL)
	// 	assert.True(t, result)
	// })

	// // 测试用例2: 响应非200状态码
	// t.Run("non-200 response", func(t *testing.T) {
	// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(http.StatusNotFound)
	// 	}))
	// 	defer server.Close()

	// 	result := Get(server.URL)
	// 	assert.False(t, result)
	// })

	// // 测试用例3: 请求超时
	// t.Run("timeout", func(t *testing.T) {
	// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		time.Sleep(3 * time.Second) // 超过2秒超时设置
	// 		w.WriteHeader(http.StatusOK)
	// 	}))
	// 	defer server.Close()

	// 	result := Get(server.URL)
	// 	assert.False(t, result)
	// })

	// // 测试用例4: 无效URL
	// t.Run("invalid URL", func(t *testing.T) {
	// 	result := Get("http://invalid.url.that.does.not.exist")
	// 	assert.False(t, result)
	// })

	// // 测试用例5: 服务器错误
	// t.Run("server error", func(t *testing.T) {
	// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		panic("server error")
	// 	}))
	// 	defer server.Close()

	// 	result := Get(server.URL)
	// 	assert.False(t, result)
	// })

	// // 测试用例6: 空字符串URL
	// t.Run("empty URL", func(t *testing.T) {
	// 	result := Get("")
	// 	assert.False(t, result)
	// })
}