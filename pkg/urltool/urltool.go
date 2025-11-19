package urltool

import (
	"net/url"
	"path"
	"errors"
)

func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}
	if len(myUrl.Host) == 0 {
		return "", errors.New("no host in targetUrl")
	}
	return path.Base(myUrl.Path), nil
}