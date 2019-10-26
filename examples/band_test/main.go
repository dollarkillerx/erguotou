/**
 * @Author: DollarKiller
 * @Description: 参数绑定
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:45 2019-10-26
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
)

func main() {
	app := erguotou.New()

	type item struct {
		Id  int    `json:"id"`
		Msg string `json:"msg"`
	}

	app.Get("/", func(ctx *erguotou.Context) {
		item := item{}
		it := ctx.BindGet(&item)
		ctx.Json(200, it)
	})

	app.Post("/", func(ctx *erguotou.Context) {
		item := item{}
		it := ctx.BindFrom(&item)
		ctx.Json(200, it)
	})

	app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(true))
}
