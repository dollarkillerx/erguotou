/**
 * @Author: DollarKiller
 * @Description: 基础中间件
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:55 2019-09-30
 */
package erguotou

import (
	"github.com/dollarkillerx/erguotou/local"
	"io/ioutil"
	"log"
	"sync"
)

var localOnce = sync.Once{}

func Logger(ctx *Context) {
	path := ctx.Ctx.Path()

	var head string

	switch {
	case ctx.Ctx.IsGet():
		head = "Get: "
	case ctx.Ctx.IsPost():
		head = "Post: "
	case ctx.Ctx.IsDelete():
		head = "Delete: "
	case ctx.Ctx.IsPut():
		head = "Put: "
	}

	log.Println(head + string(path))
	ctx.Next()
}

func Local(ctx *Context) {
	if ctx.engine.Option.Debug {
		init := local.LocalInit()
		e := init.Init("local/local.json")
		if e != nil {
			e := ioutil.WriteFile("local/local.json", []byte(local.Source), 00755)
			if e != nil {
				log.Fatal("Localization is initialized")
			}
		}
		ctx.Data("local", init.SourceMap)
	} else {
		localOnce.Do(func() {
			init := local.LocalInit()
			e := init.Init("local/local.json")
			if e != nil {
				e := ioutil.WriteFile("local/local.json", []byte(local.Source), 00755)
				if e != nil {
					log.Fatal("Localization is initialized")
				}
			}
			ctx.Data("local", init.SourceMap)
		})
	}

	ctx.Next()
}
