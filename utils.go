/**
 * @Author: DollarKiller
 * @Description: 工具库
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 18:08 2019-09-29
 */
package erguotou

import (
	"fmt"
	"github.com/dollarkillerx/erguotou/fasthttp"
	"strconv"
)

func HttpSplice(h1, h2 string) string {
	u := string(h1[len(h1)-1])
	if u == "/" {
		u = h1[:len(h1)-1]
	} else {
		u = h1
	}

	u2 := string(h2[0])
	if u2 == "/" {
		u += h2
	} else {
		u += "/" + h2
	}

	return u
}

type Utils struct {

}

func (u *Utils) Get(url string) ([]byte,error) {
	statusCode, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil,err
	}
	if statusCode != 200 {
		return nil,fmt.Errorf(strconv.Itoa(statusCode))
	}

	return body,nil
}