/**
 * @Author: DollarKiller
 * @Description: http test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:30 2019-09-29
 */
package test

import (
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/erguotou/fasthttp"
	"github.com/dollarkillerx/erguotou/fasthttprouter"
	"log"
	"testing"
)

func TestHttp(t *testing.T) {
	router := fasthttprouter.New()

	err := fasthttp.ListenAndServe(":8082", router.Handler)
	if err != nil {
		panic(err)
	}
}


func TestErguotou(t *testing.T) {
	engine := erguotou.New()
	
	engine.Use(func(ctx *erguotou.Context) {
		log.Println("1")
		ctx.Next()   // 执行下一级   反之不执行  哈哈哈
	})

	engine.Get("/hello/:name", func(ctx *erguotou.Context) {
		value, b := ctx.Value("name")
		if b {
			log.Println(value)
		}
		ctx.Next()
	},func(ctx *erguotou.Context) {
		ctx.String(200,"hello")
	})

	cpp := engine.Group("/cpp")
	{
		cpp.Get("/hh", func(ctx *erguotou.Context) {
			ctx.String(200,"我是你大爷")
		})
	}

	err := engine.Run(erguotou.SetHost(":8082"))
	if err != nil {
		panic(err)
	}
}