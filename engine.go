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
	"log"
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
		url += url + "0.0.0.0" + option.Host
	} else {
		url += url + option.Host
	}
	log.Println("Server Run " + url)

	err := fasthttp.ListenAndServe(option.Host, e.fsroot.Handler)

	return err
}
