/**
 * @Author: DollarKiller
 * @Description: 引擎
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:01 2019-09-29
 */
package erguotou

import (
	"github.com/dollarkillerx/erguotou/fasthttp"
	"github.com/dollarkillerx/erguotou/fasthttprouter"
	"html/template"
	"log"
	"sync"
)

var (
	HtmlPool *sync.Pool
)

type Engine struct {
	RouterGroup
	fsroot *fasthttprouter.Router
}

func New() *Engine {
	eng := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			root:     true,
			basePath: "/",
		},
		fsroot: fasthttprouter.New(),
	}
	eng.RouterGroup.engine = eng
	return eng
}

func (e *Engine) Run(options ...Option) error {
	// 设置
	option := Options{}
	for _, k := range options {
		k(&option)
	}
	url := "http://"

	if string(option.Host[0]) == ":" {
		url += "0.0.0.0" + option.Host
	} else {
		url += option.Host
	}
	log.Println("Server Run " + url)

	err := fasthttp.ListenAndServe(option.Host, e.fsroot.Handler)

	return err
}

// 文件服务器
func (e *Engine) Status(path, dir string) {
	u := string(path[len(path)-1])
	if u == "/" {
		path = path + "*filepath"
	} else {
		path = path + "/*filepath"
	}
	e.engine.fsroot.ServeFiles(path, dir)
}

// 注册模板  ("templates/**/*")
func (e *Engine) LoadHTMLPath(path string) {
	HtmlPool = &sync.Pool{
		New: func() interface{} {
			HtmlGlob,err := template.ParseGlob(path)
			if err != nil {
				log.Fatal(err)
			}
			return HtmlGlob
		},
	}

}