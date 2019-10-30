/**
 * @Author: DollarKiller
 * @Description: validate test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:37 2019-10-30
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

type category struct {
	Name string `form:"name" json:"name" `
	Slug string `form:"slug" json:"slug" validate:"required"`
}

func main() {
	app := erguotou.New()

	app.Post("/test", func(ctx *erguotou.Context) {
		data := category{}
		err := ctx.BindValue(&data)
		if err != nil {
			panic(err)
		}
		log.Println(data)
		validate := validator.New()
		err = validate.Struct(&data)
		if err != nil {
			log.Println("验证错误")
			return
		} else {
			log.Println("验证成功")
		}
	})

	app.Run(erguotou.SetHost(":8089"))
}
