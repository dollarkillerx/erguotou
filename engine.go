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
)

type Engine struct {
	Option Options
	RouterGroup
	router *fasthttprouter.Router
	Html   *HtmlTemplate
}

func New() *Engine {
	eng := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			root:     true,
			basePath: "/",
		},
		router: fasthttprouter.New(),
		Html:   &HtmlTemplate{},
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
	e.Option = option
	url := "http://"

	if string(option.Host[0]) == ":" {
		url += "0.0.0.0" + option.Host
	} else {
		url += option.Host
	}
	log.Println("Server Run " + url)
	log.Printf("Debug: %v", erguotouDebug)

	var err error
	if option.Size == 0 {
		err = fasthttp.ListenAndServe(option.Host, e.router.Handler)
	} else {
		err = fasthttp.ListenAndServeUpSize(option.Host, e.router.Handler, option.Size)
	}

	return err
}

// 文件服务器
func (e *Engine) Static(path, dir string) {
	u := string(path[len(path)-1])
	if u == "/" {
		path = path + "*filepath"
	} else {
		path = path + "/*filepath"
	}
	e.engine.router.ServeFiles(path, dir)
}

// 文件服务器
func (e *Engine) Status(path, dir string) {
	e.Static(path, dir)
}

// 注册模板  ("templates/**/*"),funcMap   // 这里的设计思路貌似错误了    不是一开始就激活  而是 view html才激活这里
// 写入html
func (e *Engine) LoadHTMLGlob(path string) {
	e.Html.Path = path
}

func (e *Engine) Delims(left, right string) {
	e.Html.SetDelims(left, right)
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.Html.SetFuncMap(funcMap)
}

//func (e *Engine) LoadHTMLPath(path string, funcMap template.FuncMap, left string, right string) {
//	Path = path
//	FuncMap = funcMap
//	HtmlPool = NewObjPoll(func() interface{} {
//		return e.LoadHTMLDebug()
//	}, 5)
//
//	// 创建备用临时对象池
//	HtmlTemporary = &sync.Pool{
//		New: func() interface{} {
//			return e.LoadHTMLDebug()
//		},
//	}
//
//	log.Println("Html模板加载完毕")
//}
//
//// 开发默认html热加载
//func (e *Engine) LoadHTMLDebug() *template.Template {
//	funcMap := FuncMap
//	var HtmlGlob *template.Template
//	var err error
//	if funcMap == nil {
//		HtmlGlob, err = template.New("s").ParseGlob(Path)
//	} else {
//		HtmlGlob, err = template.New("s").Delims("{[", "]}").Funcs(funcMap).ParseGlob(Path)
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//	return HtmlGlob
//}
