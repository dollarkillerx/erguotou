/**
 * @Author: DollarKiller
 * @Description: 随机
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:24 2019-10-29
 */
package util

import (
	"math/rand"
	"strconv"
	"time"
)

func SuperRand() string {
	head := int(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	body := rand.Intn(999999)
	footer := int(time.Now().UnixNano())

	encode := Sha256Encode(strconv.Itoa(head + body + footer))

	return encode
}
