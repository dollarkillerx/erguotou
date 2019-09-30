/**
 * @Author: DollarKiller
 * @Description: 路由分组
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:25 2019-09-30
 */
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	api := app.Group("/api")
	{
		api.Get("/hello", func(ctx *erguotou.Context) {
			ctx.String(200,"hello")
		})

		api.Get("/ppc", func(ctx *erguotou.Context) {
			ctx.String(200,"ppc")
		})
	}

	err := app.Run(erguotou.SetHost(":8082"))
	if err != nil {
		panic(err)
	}
}
