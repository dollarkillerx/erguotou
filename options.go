/**
 * @Author: DollarKiller
 * @Description: options 插件设计模式 做设计
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:08 2019-09-29
 */
package erguotou

var (
	erguotou_debug bool
)

func init() {
	erguotou_debug = false
}

type Options struct {
	Host string
}

type Option func(*Options)

// 设置host
func SetHost(host string) Option {
	return func(options *Options) {
		options.Host = host
	}
}

// 设置debug
func SetDebug(debug bool) {
	erguotou_debug = debug
}
