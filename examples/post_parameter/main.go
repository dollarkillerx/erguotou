/**
 * @Author: DollarKiller
 * @Description: 获取post参数
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:13 2019-09-30
 */
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	app.Get("/hello", func(ctx *erguotou.Context) {
		val := ctx.PostVal("hello")

		ctx.Write(200, val)
	})

	app.Post("/hello", func(ctx *erguotou.Context) {
		body := ctx.Body()

		ctx.Write(200, body)
	})

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}
