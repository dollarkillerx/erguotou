/**
 * @Author: DollarKiller
 * @Description: 内部高速缓存
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:56 2019-10-29
 */
package cache

import (
	"github.com/bluele/gcache"
	"github.com/dollarkillerx/erguotou/util"
	"time"
)

/**
内部高速缓存 给客户两种设置和获取方式
1. 标准数据存储 (推荐小的数据采用此方式)
2. 较大数据存储 (针对较大数据,我们会采用压缩存储,编码和解码会产生消耗)
*/

var (
	Cache gcache.Cache
)

func init() {
	Cache = gcache.New(20).LRU().Build()
}

func SetWithExpire(key, value interface{}, expiration time.Duration) error {
	return Cache.SetWithExpire(key, value, expiration)
}

func Set(key, value interface{}) error {
	return Cache.Set(key, value)
}

func Get(key interface{}) (interface{}, error) {
	return Cache.Get(key)
}

func GetEx(key interface{}) (interface{}, bool) {
	get, e := Cache.Get(key)
	if get == nil || e != nil {
		return get, false
	} else {
		return get, true
	}
}

func Exis(key interface{}) bool {
	get, e := Cache.Get(key)
	if get == nil || e != nil {
		return false
	} else {
		return true
	}
}

func GetAll(checkExpired bool) map[interface{}]interface{} {
	return Cache.GetALL(checkExpired)
}

// 存储较大数据
func SetBig(key, val string) error {
	zip := util.StrZip{}
	data := zip.Zip(val)
	return Cache.Set(key, data)
}

func SetBigWithExpire(key, value string, expiration time.Duration) error {
	zip := util.StrZip{}
	data := zip.Zip(value)
	return Cache.SetWithExpire(key, data, expiration)
}

// 获取较大数据
func GetBig(key string) (string, error) {
	get, e := Cache.Get(key)
	if e != nil {
		return "", e
	}
	das := get.(string)

	zip := util.StrZip{}
	unzip := zip.Unzip(das)
	return unzip, nil
}
