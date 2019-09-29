/**
 * @Author: DollarKiller
 * @Description: 包装上文
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:00 2019-09-29
 */
package erguotou

import (
	"github.com/dollarkillerx/erguotou/fasthttp"
)

type Context struct {
	Ctx      *fasthttp.RequestCtx // ctx
	index    int                  // 中间件计数器
	engine   *Engine
	handlers HandlersChain // 处理函数slice
}

// 处理函数
type HandlerFunc func(*Context)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

// 来到下一级  调用链
func (c *Context) Next() {
	c.index += 1
	if c.index <= len(c.handlers) {
		c.handlers[c.index-1](c)
	}
}
