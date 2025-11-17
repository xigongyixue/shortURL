package urltool

import (
	"net/url"
	"path"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(req.LongUrl)
	if err == nil {
		return nil, err
	}
	return path.Base(myUrl.Path), nil
}