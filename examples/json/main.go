/**
 * @Author: DollarKiller
 * @Description: json test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 19:07 2019-10-05
 */
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	app.Get("/", func(ctx *erguotou.Context) {
		ctx.Json(200,"ok")
	})

	app.Run(erguotou.SetHost(":8081"))
}
