/**
 * @Author: DollarKiller
 * @Description:  渲染html
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 19:58 2019-10-02
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
	"html/template"
	"math/rand"
	"time"
)

func main() {
	app := erguotou.New()

	//app.Use(erguotou.Logger)

	// 注册html
	app.LoadHTMLGlob("examples/html/view/**/*")     // 设置html目录
	app.SetFuncMap(template.FuncMap{"Add": pageAA}) // 这里设置自义定模板函数
	app.Delims("{%", "%}")                          // 这里设置模板
	app.Get("/", testhtml)

	app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(true))
}

func pageAA(page int) int {
	return page + 1
}

func testhtml(ctx *erguotou.Context) {

	data := []string{"ok", "hello", "pc", "sp", "ppr", "ssr"}
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(len(data))

	ctx.Data("Ok", data[intn])

	ctx.Data("zc", 11)

	ctx.HTML(200, "/user/hello.html")
}
