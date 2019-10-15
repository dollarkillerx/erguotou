/**
 * @Author: DollarKiller
 * @Description: session or cookie
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:16 2019-10-12
 */
package erguotou

import (
	"github.com/dollarkillerx/erguotou/fasthttp"
)

func (c *Context) GetCookie(key string) string {
	return string(c.Ctx.Request.Header.Cookie(key))
}

func (c *Context) SetCookie(key string, val string) {
	cookie := fasthttp.Cookie{}
	cookie.SetKey(key)
	cookie.SetValue(val)
	cookie.SetHTTPOnly(true)
	c.Ctx.Response.Header.SetCookie(&cookie)
}
