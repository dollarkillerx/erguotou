/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:35 2019-10-12
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
	"log"
)

func main() {
	app := erguotou.New()

	app.Get("/setcookie", func(ctx *erguotou.Context) {
		val := ctx.GetVal("name")

		ctx.SetCookie("name", string(val))

		ctx.Write(200, val)
	})

	app.Get("/getcookie", func(ctx *erguotou.Context) {
		val := ctx.GetVal("name")

		cookie := ctx.GetCookie(string(val))
		log.Println(cookie)

		ctx.Write(200, []byte(cookie))
	})

	err := app.Run(erguotou.SetHost(":8082"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}
