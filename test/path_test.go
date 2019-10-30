/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:30 2019-10-30
 */
package test

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	path := "/data/:page/:limit"
	data := []string{"1", "2"}
	spl := getSpl(path, data)
	log.Println(spl)
}

func getSpl(dst string, ppc []string) string {
	split := strings.Split(dst, "/")
	id := 0
	for k, i := range split {
		if strings.Index(i, ":") != -1 {
			split[k] = ppc[id]
			id += 1
		}
	}
	return c(split)
}

func c(cs []string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(cs), "[]"), " ", "/", -1)
}

type mainUrl struct {
}

func (m *mainUrl) getPaths(key ...string) string {
	url := string(key[0])

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

func TestMinUrl(t *testing.T) {
	url := mainUrl{}
	paths := url.getPaths("/one/:key/:index", "css", "asd")
	log.Println(paths)
}
