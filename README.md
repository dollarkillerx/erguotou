# Erguotou Web Framework
![](./README/erguotou.jpg)

erguotou 二锅头 基于fasthttp的轻量级web框架

### 入门
- [安装](#安装)
- [快速开始](#快速开始)
- [Api示范](#Api示范)
    - [RestfulApi](#RestfulApi)
    - [路径中参数](#路径中参数)
    - [Get参数](#Get参数)
    - [POST参数](#POST参数)
    - [参数绑定](#参数绑定)
    - [上传文件](#上传文件)
    - [路由分组](#路由分组)
    - [使用中间件](#使用中间件)
    - [文件服务器](#文件服务器)
    - [HTML渲染](#HTML渲染)
### 安装
``` 
go get github.com/dollarkillerx/erguotou
```
### 快速开始
``` 
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	// 注册全局中间件
	app.Use(erguotou.Logger)

	// 注册路由
	app.Get("/hello", func(ctx *erguotou.Context) {
		ctx.String(200,"hello erguotou")
	})

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}
```
## Api示范

### RestfulApi
``` 
func main() {
	app := erguotou.New()

	app.Get("/", func(ctx *erguotou.Context) {
		ctx.String(200,"hello erguotou")
	})

	app.Post("/", func(ctx *erguotou.Context) {
		ctx.String(200,"hello erguotou")
	})

	app.Delete("/", func(ctx *erguotou.Context) {
		ctx.String(200,"hello erguotou")
	})

	app.Put("/", func(ctx *erguotou.Context) {
		ctx.String(200,"hello erguotou")
	})

	app.Head("/", func(ctx *erguotou.Context) {
		ctx.String(200,"hello erguotou")
	})

	app.Options("/", func(ctx *erguotou.Context) {
		ctx.String(200,"hello erguotou")
	})

	app.Patch("/", func(ctx *erguotou.Context) {
		ctx.String(200,"hello erguotou")
	})

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}
```

### 路径中参数
``` 
app.Get("/hello/:hello", func(ctx *erguotou.Context) {
    value, b := ctx.PathValueString("hello")
    if b {
        ctx.String(200,value)
    }
})
```
### Get参数
``` 
app.Get("/hello", func(ctx *erguotou.Context) {
    val := ctx.GetVal("hello")

    ctx.Write(200,val)
})
```

### POST参数
``` 
app.Post("/hello", func(ctx *erguotou.Context) {
    val := ctx.PostVal("hello")

    ctx.Write(200,val)
})
```

### Body内容
``` 
app.Post("/hello", func(ctx *erguotou.Context) {
    body := ctx.Body()

    ctx.Write(200,body)
})
```

### 参数绑定
``` 
type user struct {
	Name string `json:"name" `
	Password string `json:"password" `
}

func TestBandJson(t *testing.T) {
	app := erguotou.New()

	data := user{}
	app.Post("/testjson", func(ctx *erguotou.Context) {
		value := ctx.BandValue(&data)
		if value != nil {
			panic(value)
		}

		ctx.Json(200,data)
	})

	app.Get("/testjson", func(ctx *erguotou.Context) {
		value := ctx.BandValue(&data)
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
```
### 上传文件
``` 
app.Get("/hello", func(ctx *erguotou.Context) {
    // FormFile 读取文件
    header, e := ctx.FormFile("file")
    if e != nil {
        ctx.String(400,"")
    }
    file, e := header.Open()
    if header != nil {
        defer file.Close()
    }

    bytes, e := ioutil.ReadAll(file)
    if e == nil {

        // SeedFile 发送文件
        ctx.SeedFileByte(bytes)
    }
})
```

### 路由分组
``` 
package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	api := app.Group("/api")
	{
		api.Get("/hello", func(ctx *erguotou.Context) {
			ctx.String(200,"hello")
		})

		api.Get("/ppc", func(ctx *erguotou.Context) {
			ctx.String(200,"ppc")
		})
	}

	err := app.Run(erguotou.SetHost(":8082"))
	if err != nil {
		panic(err)
	}
}
```

### 使用中间件
``` 
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
```

### 文件服务器
``` 
func TestFileServe(t *testing.T) {
	app := erguotou.New()

	app.Status("/hello",".")

	err := app.Run(erguotou.SetHost(":8082"))
	if err != nil {
		panic(err)
	}
}
```

### HTML渲染
``` 
func main() {
	app := erguotou.New()

	app.Use(erguotou.Logger)

	// 注册html
	app.LoadHTMLPath("examples/html/view/**/*")

	app.Get("/", testhtml)

	app.Run(erguotou.SetHost(":8081"))
}


func testhtml(ctx *erguotou.Context) {

	ctx.Data("Ok","this is ok!")

	ctx.HTML(200,"/user/hello.html")
}
```