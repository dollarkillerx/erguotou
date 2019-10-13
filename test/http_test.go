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

// 测试路由分组&中间件
func TestErguotou(t *testing.T) {
	engine := erguotou.New()

	engine.Use(func(ctx *erguotou.Context) {
		log.Println("1")
		ctx.Next() // 执行下一级   反之不执行  哈哈哈
	})

	engine.Get("/hello/:name", func(ctx *erguotou.Context) {
		value, b := ctx.PathValue("name")
		if b {
			log.Println(value)
		}
		ctx.Next()
	}, func(ctx *erguotou.Context) {
		ctx.String(200, "hello")
	})

	cpp := engine.Group("/cpp")
	{
		cpp.Get("/hh", func(ctx *erguotou.Context) {
			panic("err")
		})
	}

	err := engine.Run(erguotou.SetHost(":8082"))
	if err != nil {
		panic(err)
	}
}

type user struct {
	Name     string `json:"name" `
	Password string `json:"password" `
}

// 测试参数绑定
func TestBandJson(t *testing.T) {
	app := erguotou.New()

	data := user{}
	app.Post("/testjson", func(ctx *erguotou.Context) {
		value := ctx.BindValue(&data)
		if value != nil {
			panic(value)
		}

		ctx.Json(200, data)
	})

	app.Get("/testjson", func(ctx *erguotou.Context) {
		value := ctx.BindValue(&data)
		if value != nil {
			panic(value)
		}

		log.Println(data)
	})

	err := app.Run(erguotou.SetHost(":8082"))
	if err != nil {
		panic(err)
	}
}

// 测试文件服务器
func TestFileServe(t *testing.T) {
	app := erguotou.New()

	app.Status("/hello", ".")

	err := app.Run(erguotou.SetHost(":8082"))
	if err != nil {
		panic(err)
	}
}
