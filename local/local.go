/**
 * @Author: DollarKiller
 * @Description: local 本地化
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:45 2019-10-14
 */
package local

import (
	"encoding/json"
	"io/ioutil"
)

type Local struct {
	source    []byte                            `json:"source"` // init 数据源
	SourceMap map[string]map[string]interface{} // init 数据源map
}

func LocalInit() *Local {
	return &Local{}
}

// init 配置文件
func (l *Local) Init(path string) error {
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
	l.source = bytes

	source := make(map[string]map[string]interface{})

	e = json.Unmarshal(bytes, &source)
	if e != nil {
		return e
	}
	l.SourceMap = source
	return nil
}

var Source = `
{
  "zh" : {
    "name": "名称"
  },
  "fr" : {
    "name": "fr name"
  }
}
`
