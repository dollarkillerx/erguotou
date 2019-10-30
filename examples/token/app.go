/**
 * @Author: DollarKiller
 * @Description: token
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 09:01 2019-10-30
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/erguotou/token"
	"log"
	"time"
)

func main() {
	app := erguotou.New()

	app.Get("/", func(ctx *erguotou.Context) {
		jwt := token.NewJwt()
		jwt.User = "dollarkiller"
		jwt.TimeOut = time.Hour * 6

		s, e := token.Token.GeneraJwtToken(jwt)
		if e != nil {
			log.Fatal("token 生成失败")
		}
		ctx.Json(200, map[string]interface{}{
			"token": s,
		})
	})

	app.Get("/user", func(ctx *erguotou.Context) {
		token := ctx.Header("token")
		ctx.String(200, token)
	})

	app.Run(erguotou.SetHost(":8086"))
}
