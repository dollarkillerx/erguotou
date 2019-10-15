/**
 * @Author: DollarKiller
 * @Description: 本地化测试
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:05 2019-10-14
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
)

func main() {
	app := erguotou.New()

	app.Use(erguotou.Local("en"))

	app.Get("/", func(ctx *erguotou.Context) {
		// 使用 html 上 {{.local.zh.name}}
	})

	app.Run(erguotou.SetHost(":8081"))
}
