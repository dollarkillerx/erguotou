/**
 * @Author: DollarKiller
 * @Description: 路径中参数
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:06 2019-09-30
 */
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	app.Get("/hello/:hello", func(ctx *erguotou.Context) {
		value, b := ctx.PathValueString("hello")
		if b {
			ctx.String(200, value)
		}
	})

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}
