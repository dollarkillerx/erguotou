/**
 * @Author: DollarKiller
 * @Description: 初始二锅头
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:36 2019-09-29
 */
package erguotou

import (
	"fmt"
	"log"
)

func init() {
	data := fmt.Sprintf("%c[1;31;40m%c[0m %v", 0x1B, 0x1B, "Ergoutou 二锅头 初始化完毕")

	log.Println(data)
}
