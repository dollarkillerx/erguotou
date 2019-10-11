/**
 * @Author: DollarKiller
 * @Description: options 插件设计模式 做设计
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 15:08 2019-09-29
 */
package erguotou

var (
	erguotouDebug bool
)

func init() {
	erguotouDebug = false
}

type Options struct {
	Host string
	Size int
}

type Option func(*Options)

// 设置host
func SetHost(host string) Option {
	return func(options *Options) {
		options.Host = host
	}
}

// 设置debug
func SetDebug(debug bool) Option {
	erguotouDebug = debug
	return func(options *Options) {
	}
}

// 设置上传大小
func SetUploadSize(size int) Option {
	return func(options *Options) {
		options.Size = size
	}
}
