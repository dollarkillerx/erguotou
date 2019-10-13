/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:03 2019-10-13
 */
package test

import (
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/erguotou/session"
	"log"
	"testing"
)

func TestSession(t *testing.T) {
	engine := erguotou.New()

	engine.Get("/hello", func(ctx *erguotou.Context) {
		cache := session.GetSessionCache()
		cache.Set("name", "hello")
		cache.Set("ppc", "sdasd")
		cache.Save(ctx)
	})

	engine.Get("/ppc", func(ctx *erguotou.Context) {
		cache := session.GetSessionCache()
		get, b := cache.Get(ctx, "name")
		if b {
			log.Println(get)
		}
		get, b = cache.Get(ctx, "ppc")
		if b {
			log.Println(get)
		}
	})

	err := engine.Run(erguotou.SetHost(":8082"))
	if err != nil {
		panic(err)
	}
}
