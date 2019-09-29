/**
 * @Author: DollarKiller
 * @Description: 引擎
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:01 2019-09-29
 */
package erguotou

import (
	"github.com/dollarkillerx/erguotou/fasthttprouter"
	"github.com/dollarkillerx/erguotou/fasthttp"
)

type Engine struct {
	RouterGroup
	fsroot *fasthttprouter.Router
}

func New() *Engine {
	eng := &Engine{
		RouterGroup:RouterGroup{
			Handlers:nil,
			root:true,
			basePath:"/",
		},
		fsroot:fasthttprouter.New(),
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
	err := fasthttp.ListenAndServe(option.Host, e.fsroot.Handler)
	return err
}