/**
 * @Author: DollarKiller
 * @Description: session or cookie
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:16 2019-10-12
 */
package erguotou

import (
	"github.com/dollarkillerx/erguotou/fasthttp"
	"time"
)

func (c *Context) GetCookie(key string) string {
	return string(c.Ctx.Request.Header.Cookie(key))
}

func (c *Context) SetCookie(key string, val string) {
	cookie := fasthttp.Cookie{}
	cookie.SetKey(key)
	cookie.SetValue(val)
	cookie.SetHTTPOnly(true)
	cookie.SetPath("/")
	c.Ctx.Response.Header.SetCookie(&cookie)
}

func (c *Context) SetCookieTime(key,val string,ti time.Duration) {
	cookie := fasthttp.Cookie{}
	cookie.SetKey(key)
	cookie.SetValue(val)
	cookie.SetHTTPOnly(true)
	cookie.SetPath("/")
	//duration := time.Second * 1000
	////time := time.Now() + duration
	////ti
	//now := time.Now()
	//now.Hour() =
	//cookie.SetExpire()
	c.Ctx.Response.Header.SetCookie(&cookie)
}