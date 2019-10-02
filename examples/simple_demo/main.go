/**
 * @Author: DollarKiller
 * @Description: 快速入门教程
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:49 2019-09-30
 */
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	// 全局基础中间件
	//app.Use(erguotou.Logger)

	// 注册路由
	app.Get("/hello", func(ctx *erguotou.Context) {
		ctx.String(200, "test")
	})

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}
