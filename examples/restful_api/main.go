/**
 * @Author: DollarKiller
 * @Description:restful api
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:02 2019-09-30
 */
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	app.Get("/", func(ctx *erguotou.Context) {
		ctx.String(200, "hello erguotou")
	})

	app.Post("/", func(ctx *erguotou.Context) {
		ctx.String(200, "hello erguotou")
	})

	app.Delete("/", func(ctx *erguotou.Context) {
		ctx.String(200, "hello erguotou")
	})

	app.Put("/", func(ctx *erguotou.Context) {
		ctx.String(200, "hello erguotou")
	})

	app.Head("/", func(ctx *erguotou.Context) {
		ctx.String(200, "hello erguotou")
	})

	app.Options("/", func(ctx *erguotou.Context) {
		ctx.String(200, "hello erguotou")
	})

	app.Patch("/", func(ctx *erguotou.Context) {
		ctx.String(200, "hello erguotou")
	})

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}
