/**
 * @Author: DollarKiller
 * @Description:  渲染html
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 19:58 2019-10-02
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
)

func main() {
	app := erguotou.New()

	//app.Use(erguotou.Logger)

	// 注册html
	app.LoadHTMLPath("examples/html/view/**/*")

	app.Get("/", testhtml)

	app.Run(erguotou.SetHost(":8081"),erguotou.SetDebug(true))
}


func testhtml(ctx *erguotou.Context) {

	ctx.Data("Ok","test")

	ctx.HTML(200,"/user/hello.html")
}