### api测试部分
erguotou
```
import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	// 全局基础中间件
	//app.Use(erguotou.Logger)

	// 注册路由
	app.Get("/hello", func(ctx *erguotou.Context) {
		ctx.String(200, "test")
	})

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false))
	if err != nil {
		panic(err)
	}
}

```
gin 

```
app := gin.New()

app.GET("/hello", func(context *gin.Context) {
    context.String(200,"test")
})

app.Run(":8082")
```

### html 部分
html
``` 
{{define  "/user/hello.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Ok}}</title>
</head>
<body>
{{.Ok}}
</body>
</html>

{{end}}
```
erguotou
``` 
import (
	"github.com/dollarkillerx/erguotou"
)

func main() {
	app := erguotou.New()

	//app.Use(erguotou.Logger)

	// 注册html
	app.LoadHTMLPath("examples/html/view/**/*")

	app.Get("/", testhtml)

	app.Run(erguotou.SetHost(":8081"))
}


func testhtml(ctx *erguotou.Context) {

	ctx.Data("Ok","test")

	ctx.HTML(200,"/user/hello.html")
}
```
gin 
``` 
router := gin.New()
router.LoadHTMLGlob("view/**/*")
router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "/user/hello.html", gin.H{
        "Ok": "test",
    })
})
router.Run(":8082")
```