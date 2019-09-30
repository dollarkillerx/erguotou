/**
 * @Author: DollarKiller
 * @Description: 获取git参数
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:10 2019-09-30
 */
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	app.Get("/hello", func(ctx *erguotou.Context) {
		val := ctx.GetVal("hello")

		ctx.Write(200,val)
	})

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}
