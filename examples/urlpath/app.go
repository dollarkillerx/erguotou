/**
 * @Author: DollarKiller
 * @Description: urlpath  反向解析
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:00 2019-10-30
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
	"html/template"
)

func main() {
	app := erguotou.New()

	app.Get("/cs", func(ctx *erguotou.Context) {

	}).Name("cs")

	app.Get("/pc/:id/:cs", func(ctx *erguotou.Context) {

	}).Name("pc")

	app.SetFuncMap(template.FuncMap{"add": func(a int) int { return a + 1 }})

	app.LoadHTMLGlob("examples/urlpath/view/**/*")

	app.Get("/", func(ctx *erguotou.Context) {
		ctx.Data("c", 2)
		ctx.HTML(200, "/user/hello.html")
	})

	app.Run(erguotou.SetHost(":8085"), erguotou.SetDebug(true))
}
