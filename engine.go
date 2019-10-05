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
	HtmlPool      *ObjPoll
	HtmlTemporary *sync.Pool
	Path          string
	FuncMap       template.FuncMap
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
	log.Printf("Debug: %v", erguotou_debug)

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

// 注册模板  ("templates/**/*"),funcMap
func (e *Engine) LoadHTMLPath(path string,funcMap template.FuncMap) {
	Path = path
	FuncMap = funcMap
	HtmlPool = NewObjPoll(func() interface{} {
		return e.LoadHTMLDebug()
	}, 5)

	// 创建备用临时对象池
	HtmlTemporary = &sync.Pool{
		New: func() interface{} {
			return e.LoadHTMLDebug()
		},
	}

	log.Println("Html模板加载完毕")
}

// 开发默认html热加载
func (e *Engine) LoadHTMLDebug() *template.Template {
	funcMap := FuncMap
	HtmlGlob, err := template.ParseGlob(Path)
	if funcMap != nil {
		HtmlGlob.Funcs(funcMap)
	}
	if err != nil {
		log.Fatal(err)
	}
	return HtmlGlob
}
