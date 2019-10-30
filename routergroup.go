/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 16:04 2019-09-29
 */
package erguotou

import (
	"github.com/dollarkillerx/erguotou/fasthttp"
	"log"
)

// 内部路由器
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine // 注册组路由的时候传入
	root     bool    // 注册中间件的时候判断是否是根
}

// 注册全局中间件
func (r *RouterGroup) Use(middleware ...HandlerFunc) {
	r.Handlers = append(r.Handlers, middleware...)
}

// 注册组路由
func (r *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	heads := append(r.Handlers, handlers...)

	return &RouterGroup{
		Handlers: heads,
		basePath: HttpSplice(r.basePath, relativePath),
		engine:   r.engine,
		root:     false,
	}
}

func (r *RouterGroup) Get(relativePath string, handlers ...HandlerFunc) *urlp {
	url := HttpSplice(r.basePath, relativePath)
	handle := append(r.Handlers, handlers...)
	r.handle("GET", url, handle)
	return &urlp{url: url}
}

func (r *RouterGroup) Post(relativePath string, handlers ...HandlerFunc) *urlp {
	url := HttpSplice(r.basePath, relativePath)
	handle := append(r.Handlers, handlers...)
	r.handle("POST", url, handle)
	return &urlp{url: url}

}

func (r *RouterGroup) Delete(relativePath string, handlers ...HandlerFunc) *urlp {
	url := HttpSplice(r.basePath, relativePath)
	handle := append(r.Handlers, handlers...)
	r.handle("DELETE", url, handle)
	return &urlp{url: url}
}

func (r *RouterGroup) Put(relativePath string, handlers ...HandlerFunc) *urlp {
	url := HttpSplice(r.basePath, relativePath)
	handle := append(r.Handlers, handlers...)
	r.handle("PUT", url, handle)
	return &urlp{url: url}
}

func (r *RouterGroup) Patch(relativePath string, handlers ...HandlerFunc) *urlp {
	url := HttpSplice(r.basePath, relativePath)
	handle := append(r.Handlers, handlers...)
	r.handle("PATCH", url, handle)
	return &urlp{url: url}
}

func (r *RouterGroup) Head(relativePath string, handlers ...HandlerFunc) *urlp {
	url := HttpSplice(r.basePath, relativePath)
	handle := append(r.Handlers, handlers...)
	r.handle("HEAD", url, handle)
	return &urlp{url: url}
}

func (r *RouterGroup) Options(relativePath string, handlers ...HandlerFunc) *urlp {
	url := HttpSplice(r.basePath, relativePath)
	handle := append(r.Handlers, handlers...)
	r.handle("OPTIONS", url, handle)
	return &urlp{url: url}
}

// 完全注册路由
func (r *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) {
	ctx := Context{
		index:    0,
		engine:   r.engine,
		handlers: handlers,
	}

	var err error

	switch httpMethod {
	case "POST":
		r.engine.router.POST(relativePath, func(ctxF *fasthttp.RequestCtx) {
			defer func() {
				if re := recover(); re != nil {
					ctx.Ctx.SetStatusCode(500)

					// 打印错误信息
					log.Println(re)
					// 打印堆栈信息
					utils := Utils{}
					trace := utils.PanicTrace(2048)
					log.Println(string(trace))
					_, err = ctx.Ctx.WriteString("server error")
				}
			}()
			ctx.Ctx = ctxF
			ctx.Next()
		})
	case "GET":
		r.engine.router.GET(relativePath, func(ctxF *fasthttp.RequestCtx) {
			defer func() {
				if re := recover(); re != nil {
					ctx.Ctx.SetStatusCode(500)
					// 打印错误信息
					log.Println(re)
					// 打印堆栈信息
					utils := Utils{}
					trace := utils.PanicTrace(2048)
					log.Println(string(trace))
					_, err = ctx.Ctx.WriteString("server error")
				}
			}()
			ctx.Ctx = ctxF
			ctx.Next()
		})
	case "DELETE":
		r.engine.router.DELETE(relativePath, func(ctxF *fasthttp.RequestCtx) {
			defer func() {
				if re := recover(); re != nil {
					ctx.Ctx.SetStatusCode(500)
					// 打印错误信息
					log.Println(re)
					// 打印堆栈信息
					utils := Utils{}
					trace := utils.PanicTrace(2048)
					log.Println(string(trace))
					_, err = ctx.Ctx.WriteString("server error")
				}
			}()
			ctx.Ctx = ctxF
			ctx.Next()
		})
	case "PUT":
		r.engine.router.PUT(relativePath, func(ctxF *fasthttp.RequestCtx) {
			defer func() {
				if re := recover(); re != nil {
					ctx.Ctx.SetStatusCode(500)
					// 打印错误信息
					log.Println(re)
					// 打印堆栈信息
					utils := Utils{}
					trace := utils.PanicTrace(2048)
					log.Println(string(trace))
					_, err = ctx.Ctx.WriteString("server error")
				}
			}()
			ctx.Ctx = ctxF
			ctx.Next()
		})
	case "PATCH":
		r.engine.router.PATCH(relativePath, func(ctxF *fasthttp.RequestCtx) {
			defer func() {
				if re := recover(); re != nil {
					ctx.Ctx.SetStatusCode(500)
					// 打印错误信息
					log.Println(re)
					// 打印堆栈信息
					utils := Utils{}
					trace := utils.PanicTrace(2048)
					log.Println(string(trace))
					_, err = ctx.Ctx.WriteString("server error")
				}
			}()
			ctx.Ctx = ctxF
			ctx.Next()
		})
	case "HEAD":
		r.engine.router.HEAD(relativePath, func(ctxF *fasthttp.RequestCtx) {
			defer func() {
				if re := recover(); re != nil {
					ctx.Ctx.SetStatusCode(500)
					// 打印错误信息
					log.Println(re)
					// 打印堆栈信息
					utils := Utils{}
					trace := utils.PanicTrace(2048)
					log.Println(string(trace))
					_, err = ctx.Ctx.WriteString("server error")
				}
			}()
			ctx.Ctx = ctxF
			ctx.Next()
		})
	case "OPTIONS":
		r.engine.router.OPTIONS(relativePath, func(ctxF *fasthttp.RequestCtx) {
			defer func() {
				if re := recover(); re != nil {
					ctx.Ctx.SetStatusCode(500)
					// 打印错误信息
					log.Println(re)
					// 打印堆栈信息
					utils := Utils{}
					trace := utils.PanicTrace(2048)
					log.Println(string(trace))
					_, err = ctx.Ctx.WriteString("server error")
				}
			}()
			ctx.Ctx = ctxF
			ctx.Next()
		})
	}

	if err != nil {
		log.Println(err)
	}
}
