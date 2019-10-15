/**
 * @Author: DollarKiller
 * @Description: 基础中间件
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:55 2019-09-30
 */
package erguotou

import (
	"github.com/dollarkillerx/erguotou/clog"
	"github.com/dollarkillerx/erguotou/local"
	"io/ioutil"
	"log"
	"sync"
)

var localOnce sync.Once
var localOnce2 sync.Once

func init() {
	localOnce = sync.Once{}
	localOnce2 = sync.Once{}
}

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

func Local(language string) func(ctx *Context) {
	return func(ctx *Context) {
		localOnce2.Do(func() {
			ctx.SetCookie("language", language)
		})
		language = ctx.GetCookie("language")
		if ctx.engine.Option.Debug {
			init := local.LocalInit()
			e := init.Init("local/local.json")
			if e != nil {
				e := ioutil.WriteFile("local/local.json", []byte(local.Source), 00755)
				if e != nil {
					log.Fatal("Localization is initialized")
				}
			}
			i,ok := init.SourceMap[language]
			if ok {
				ctx.Data("local", i)
			}else {
				clog.PrintWa("No language pack")
			}
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
				i,ok := init.SourceMap[language]
				if ok {
					ctx.Data("local", i)
				}else {
					clog.PrintWa("No language pack")
				}
			})
		}

		ctx.Next()
	}

}
