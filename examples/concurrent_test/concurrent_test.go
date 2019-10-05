/**
 * @Author: DollarKiller
 * @Description: 并发测试
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:51 2019-10-02
 */
package concurrent_test

import (
	"github.com/dollarkillerx/erguotou"
	"log"
	"testing"
)

func TestHtmlCurrent(t *testing.T) {
	utils := erguotou.Utils{}

	for i := 0; i < 10000; i++ {
		bytes, e := utils.Get("http://0.0.0.0:8081")
		if e != nil {
			log.Println(string(bytes))
		}
	}
}
