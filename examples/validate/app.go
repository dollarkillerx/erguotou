/**
 * @Author: DollarKiller
 * @Description: validate
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:53 2019-10-30
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
	"log"
)

type category struct {
	Name string `form:"name" json:"name" validate:"required"`
	Slug string `form:"slug" json:"slug" validate:"required"`
}

func main() {
	app := erguotou.New()

	app.Post("/test", func(ctx *erguotou.Context) {
		inputData := &category{}
		err := ctx.BindValue(inputData)
		if err != nil {
			log.Fatalln(err)
		}
		// 默认就验证了validate
		ctx.Json(200, inputData)
	})

	app.Run(erguotou.SetHost(":8089"))
}
