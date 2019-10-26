/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:57 2019-10-26
 */
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	app.Get("/hello", func(ctx *erguotou.Context) {
		ctx.Write(200, []byte("Hello World"))
	})

	app.Get("/hello2", func(ctx *erguotou.Context) {
		ctx.String(200, "Hello World")
	})

	app.Get("/hello3", func(ctx *erguotou.Context) {
		ctx.Json(200, erguotou.H{"code": 200, "msg": "Hello World"})
	})

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}
