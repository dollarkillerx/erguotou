/**
 * @Author: DollarKiller
 * @Description: 基础中间件
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:55 2019-09-30
 */
package erguotou

import (
	"log"
)

func Logger(ctx *Context) {
	path := ctx.Ctx.Path()

	var head string

	switch {
	case ctx.Ctx.IsGet():
		head = "Get: "
	case ctx.Ctx.IsPost():
		head = "Post: "
	case ctx.Ctx.IsDelete():
		head = "Delete: "
	case ctx.Ctx.IsPut():
		head = "Put: "
	}

	log.Println(head + string(path))
	ctx.Next()
}
