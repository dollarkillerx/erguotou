/**
 * @Author: DollarKiller
 * @Description: 工具库
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 18:08 2019-09-29
 */
package erguotou

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
