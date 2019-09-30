/**
 * @Author: DollarKiller
 * @Description: 包装上文
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:00 2019-09-29
 */
package erguotou

import (
	"github.com/dollarkillerx/erguotou/clog"
	"github.com/dollarkillerx/erguotou/fasthttp"
	"mime/multipart"
)

type Context struct {
	Ctx      *fasthttp.RequestCtx // ctx
	index    int                  // 中间件计数器
	engine   *Engine
	handlers HandlersChain // 处理函数slice
}

// 处理函数
type HandlerFunc func(ctx *Context)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

// 来到下一级  调用链
func (c *Context) Next() {
	c.index += 1
	if c.index <= len(c.handlers) {
		c.handlers[c.index-1](c)
	}
}

// 返回string
func (c *Context) String(code int, msg string) {
	c.Ctx.SetStatusCode(code)
	c.Ctx.WriteString(msg)
}

// 返回json
func (c *Context) Json(code int, msg interface{}) {
	c.Ctx.SetStatusCode(code)
	c.Ctx.SetContentType("application/json")
	bytes, e := Jsonp.Marshal(msg)
	if e != nil {
		clog.PrintWa(e)
		return
	}
	c.Ctx.Write(bytes)
}

// 返回[]byte
func (c *Context) Write(code int, msg []byte) {
	c.Ctx.SetStatusCode(code)
	c.Ctx.SetContentType("application/json")
	c.Ctx.Write(msg)
}

// 获取path value
func (c *Context) PathValue(val string) (interface{}, bool) {
	value := c.Ctx.UserValue(val)
	if value == nil {
		return nil, false
	}

	return value, true
}

// 获取参数path string
func (c *Context) PathValueString(val string) (string, bool) {
	value := c.Ctx.UserValue(val)
	if value == nil {
		return "", false
	}

	s, ok := value.(string)
	return s, ok
}

// 获取参数path int
func (c *Context) PathValueInt(val string) (int, bool) {
	value := c.Ctx.UserValue(val)
	if value == nil {
		return 0, false
	}

	s, ok := value.(int)
	return s, ok
}

// 返回文件
func (c *Context) SeedFile(path string) {
	c.Ctx.SetStatusCode(200)
	c.Ctx.SendFile(path)
}

// 返回文件bytes
func (c *Context) SeedFileByte(file []byte) {
	c.Ctx.SetStatusCode(200)
	c.Ctx.SendFileBytes(file)
}

// 获取get数据
func (c *Context) GetVal(key string) []byte {
	args := c.Ctx.QueryArgs()

	peek := args.Peek(key)

	return peek
}

// 获取post数据
func (c *Context) PostVal(key string) []byte {
	args := c.Ctx.PostArgs()

	peek := args.Peek(key)

	return peek
}

// 获取body数据
func (c *Context) Body() []byte {
	return c.Ctx.PostBody()
}

// 接受文件
func (c *Context) FormFile(file string) (*multipart.FileHeader,error) {
	return c.Ctx.FormFile(file)
}

