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

func Logger(ctx *Context)  {
	path := ctx.Ctx.Path()

	log.Println(string(path))
}
