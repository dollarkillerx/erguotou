# Erguotou Web Framework
![](./README/erguotou.jpg)

erguotou 二锅头 基于fasthttp的轻量级web框架

### 入门
- [安装](#安装)
- [快速开始](#快速开始)
- [性能测试](#性能测试)  测试性能优于gin
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

### 性能测试
- erguotou
``` 
➜  Test wrk -t12 -c400 -d30s http://0.0.0.0:8081/hello
Running 30s test @ http://0.0.0.0:8081/hello
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.69ms    1.03ms 117.82ms   97.54%
    Req/Sec     8.21k     4.53k   13.52k    68.48%
  2701628 requests in 30.10s, 349.44MB read
  Socket errors: connect 155, read 70, write 0, timeout 0
Requests/sec:  89752.96
Transfer/sec:     11.61MB
```

``` 
➜  Test wrk -t12 -c400 -d30s --latency http://0.0.0.0:8081/hello
Running 30s test @ http://0.0.0.0:8081/hello
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.89ms    1.90ms 126.43ms   96.70%
    Req/Sec     7.16k     4.39k   13.00k    46.44%
  Latency Distribution
     50%    2.73ms
     75%    2.96ms
     90%    3.35ms
     99%    8.07ms
  2571271 requests in 30.10s, 333.30MB read
  Socket errors: connect 155, read 74, write 0, timeout 0
Requests/sec:  85414.79
Transfer/sec:     11.07MB
```
- gin
``` 
➜  Test wrk -t12 -c400 -d30s http://0.0.0.0:8082/hello
Running 30s test @ http://0.0.0.0:8082/hello
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.10ms   24.11ms   1.00s    99.49%
    Req/Sec     5.31k     3.64k   22.98k    48.79%
  1905011 requests in 30.10s, 218.01MB read
  Socket errors: connect 155, read 84, write 0, timeout 0
Requests/sec:  63282.30
Transfer/sec:      7.24MB
```

``` 
➜  Test wrk -t12 -c400 -d30s --latency http://0.0.0.0:8082/hello
Running 30s test @ http://0.0.0.0:8082/hello
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.24ms   26.01ms   1.02s    99.53%
    Req/Sec     5.32k     3.47k   12.66k    53.40%
  Latency Distribution
     50%    3.42ms
     75%    4.18ms
     90%    5.41ms
     99%   17.25ms
  1901550 requests in 30.07s, 217.62MB read
  Socket errors: connect 155, read 59, write 0, timeout 0
Requests/sec:  63238.53
Transfer/sec:      7.24MB
```

- erguotou html
``` 
➜  Test wrk -t12 -c400 -d30s http://0.0.0.0:8081
Running 30s test @ http://0.0.0.0:8081
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.25ms    8.40ms 404.41ms   98.48%
    Req/Sec     9.64k     4.29k   34.86k    77.26%
  2591623 requests in 30.10s, 602.23MB read
  Socket errors: connect 155, read 74, write 0, timeout 0
Requests/sec:  86101.85
Transfer/sec:     20.01MB
```

``` 
➜  Test wrk -t12 -c400 -d30s --latency http://0.0.0.0:8081      
Running 30s test @ http://0.0.0.0:8081
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.14ms    7.33ms 355.55ms   98.36%
    Req/Sec     7.92k     4.08k   16.39k    56.52%
  Latency Distribution
     50%    2.31ms
     75%    3.13ms
     90%    4.89ms
     99%   12.44ms
  2607365 requests in 30.10s, 605.77MB read
  Socket errors: connect 155, read 67, write 0, timeout 0
  Non-2xx or 3xx responses: 1
Requests/sec:  86617.57
Transfer/sec:     20.12MB
```
- gin html
``` 
➜  Test wrk -t12 -c400 -d30s http://0.0.0.0:8082      
Running 30s test @ http://0.0.0.0:8082
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    23.39ms   14.52ms 527.53ms   97.52%
    Req/Sec     0.87k   449.67     2.35k    61.40%
  309982 requests in 30.06s, 72.72MB read
  Socket errors: connect 155, read 187, write 0, timeout 0
Requests/sec:  10311.24
Transfer/sec:      2.42MB
```

``` 
➜  Test wrk -t12 -c400 -d30s --latency http://0.0.0.0:8082
Running 30s test @ http://0.0.0.0:8082
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    22.08ms   13.85ms 534.41ms   98.69%
    Req/Sec     0.92k   450.04     1.97k    53.33%
  Latency Distribution
     50%   21.20ms
     75%   22.91ms
     90%   24.87ms
     99%   36.05ms
  329183 requests in 30.05s, 77.23MB read
  Socket errors: connect 155, read 180, write 0, timeout 0
Requests/sec:  10955.90
Transfer/sec:      2.57MB
```

- 测试环境
```
MacBook Pro (Retina, 15-inch, Mid 2015)
majave10.14.6
2.2 GHz Intel Core i7
16G DDR3 1600
测试代码详见examples test.md
```
报告案例
``` 
Running 30s test @ http://www.baidu.com （压测时间30s）
  12 threads and 400 connections （共12个测试线程，400个连接）
              （平均值） （标准差）  （最大值）（正负一个标准差所占比例）
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    （延迟）
    Latency   386.32ms  380.75ms   2.00s    86.66%
    (每秒请求数)
    Req/Sec    17.06     13.91   252.00     87.89%
  Latency Distribution （延迟分布）
     50%  218.31ms
     75%  520.60ms
     90%  955.08ms
     99%    1.93s 
  4922 requests in 30.06s, 73.86MB read (30.06s内处理了4922个请求，耗费流量73.86MB)
  Socket errors: connect 0, read 0, write 0, timeout 311 (发生错误数)
Requests/sec:    163.76 (QPS 163.76,即平均每秒处理请求数为163.76)
Transfer/sec:      2.46MB (平均每秒流量2.46MB)
```