/**
 * @Author: DollarKiller
 * @Description: 默认的funcMap
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:15 2019-10-30
 */
package erguotou

import "log"

// 反向解析
func urlpath(args ...string) string {
	log.Println(args)
	i := len(args)
	if i == 1 {
		return mainUrls.getPath(string(args[0]))
	} else if i > 1 {
		return mainUrls.getPaths(args...)
	}
	return ""
}
