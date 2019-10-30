/**
 * @Author: DollarKiller
 * @Description: url 做反向解析
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:01 2019-10-30
 */
package erguotou

import (
	"fmt"
	"strings"
	"sync"
)

// 中心url
type mainUrl struct {
	Data sync.Map
}

// url做解析用
type urlp struct {
	url  string
	name string
}

var mainUrls *mainUrl

func init() {
	mainUrls = &mainUrl{}
}

func (u *urlp) Name(name string) {
	u.name = name
	mainUrls.Data.Store(u.name, u)
}

func (m *mainUrl) getPath(key string) string {
	value, ok := m.Data.Load(key)
	if ok {
		u := value.(*urlp)
		return u.url
	} else {
		return ""
	}
}

func (m *mainUrl) getPaths(key ...string) string {
	name := string(key[0])

	value, ok := m.Data.Load(name)
	if !ok {
		return ""
	}
	u := value.(*urlp)
	url := u.url

	count := strings.Count(url, "/:")
	if count != len(key)-1 {
		return ""
	}

	args := key[1:]
	split := strings.Split(url, "/")
	id := 0
	for k, i := range split {
		if strings.Index(i, ":") != -1 {
			split[k] = args[id]
			id += 1
		}
	}

	return m.listToString(split)
}

func (m *mainUrl) listToString(data []string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(data), "[]"), " ", "/", -1)
}
