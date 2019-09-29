/**
 * @Author: DollarKiller
 * @Description: http test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:30 2019-09-29
 */
package test

import (
	"github.com/dollarkillerx/erguotou/fasthttp"
	"github.com/dollarkillerx/erguotou/fasthttprouter"
	"testing"
)

func TestHttp(t *testing.T) {
	router := fasthttprouter.New()

	err := fasthttp.ListenAndServe(":8082", router.Handler)
	if err != nil {
		panic(err)
	}
}