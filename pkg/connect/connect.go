package connect

import (
	"net/http"
	"time"
	"github.com/zeromicro/go-zero/core/logx"
)

// 全局http客户端
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// Get 判断url是否可达
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorw("connect client.Get failed", logx.Field("err", err.Error()))
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}