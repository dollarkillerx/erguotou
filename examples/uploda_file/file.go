/**
 * @Author: DollarKiller
 * @Description: upload file
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:18 2019-09-30
 */
package main

import (
	"github.com/dollarkillerx/erguotou"
	"io/ioutil"
)

func main() {
	app := erguotou.New()

	app.Get("/hello", func(ctx *erguotou.Context) {
		// FormFile 读取文件
		header, e := ctx.FormFile("file")
		if e != nil {
			ctx.String(400, "")
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

	err := app.Run(erguotou.SetHost(":8081"), erguotou.SetDebug(false), erguotou.SetUploadSize(8<<20)) // 8M
	if err != nil {
		panic(err)
	}
}
