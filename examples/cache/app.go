/**
 * @Author: DollarKiller
 * @Description: 高速缓存
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 08:52 2019-10-30
 */
package main

import (
	"github.com/dollarkillerx/erguotou/cache"
	"log"
	"time"
)

func main() {
	err := cache.Set("key", "value this is test")
	if err != nil {
		log.Fatal(err)
	}
	val, err := cache.Get("key")
	if err != nil {
		log.Fatal(val)
	}
	log.Println(val.(string))

	// 设置过期时间
	err = cache.SetWithExpire("name", "dollarkiller", time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	val, err = cache.Get("name")
	if err != nil {
		log.Fatal(val)
	}
	log.Println(val.(string))

	// 存储较大的data
	data := `jaisdjaisjiejwinieowniewnfioixomxmaihfcnioani21u321437483yr843yfhfaoshadioqjiqmx9032032`

	err = cache.SetBig("big", data)
	if err != nil {
		log.Println(err)
	}

	s, err := cache.GetBig("big")
	if err != nil {
		log.Println(err)
	}

	log.Println(s)
}
