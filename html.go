/**
 * @Author: DollarKiller
 * @Description: view
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:01 2019-10-21
 */
package erguotou

import (
	"html/template"
	"log"
	"sync"
)

type delims struct {
	Left  string
	Right string
}

type HtmlTemplate struct {
	mu            sync.Mutex
	delims        *delims          // 左右的left
	FuncMap       template.FuncMap // func map
	Path          string           // path 路径
	HtmlPool      *ObjPool         // 对象池
	HtmlTemporary *sync.Pool       // 临时对象池
}

func NewHtmlTemplate(path string) *HtmlTemplate {
	return &HtmlTemplate{
		Path: path,
	}
}

func (h *HtmlTemplate) SetDelims(left, right string) {
	h.delims = &delims{}
	h.delims.Left = left
	h.delims.Right = right
}

func (h *HtmlTemplate) SetFuncMap(funcMap template.FuncMap) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.FuncMap == nil {
		h.FuncMap = funcMap
	} else {
		for k, v := range funcMap {
			h.FuncMap[k] = v
		}
	}
}

// 显示html
func (h *HtmlTemplate) Show() {
	if h.HtmlPool == nil {
		h.HtmlPool = NewObjPoll(func() interface{} {
			return h.LoadHTMLDebug()
		}, 5)
	}
	// 创建备用临时对象池
	if h.HtmlTemporary == nil {
		h.HtmlTemporary = &sync.Pool{
			New: func() interface{} {
				return h.LoadHTMLDebug()
			},
		}
	}
}

// 开发默认html热加载
func (e *HtmlTemplate) LoadHTMLDebug() *template.Template {
	HtmlGlob := &template.Template{}
	var err error
	if e.FuncMap == nil && e.delims == nil {
		HtmlGlob, err = template.New("").ParseGlob(e.Path)
	} else if e.FuncMap != nil && e.delims == nil {
		HtmlGlob, err = template.New("").Funcs(e.FuncMap).ParseGlob(e.Path)
	} else if e.FuncMap == nil && e.delims != nil {
		HtmlGlob, err = template.New("").Delims(e.delims.Left, e.delims.Right).ParseGlob(e.Path)
	} else if e.FuncMap != nil && e.delims != nil {
		HtmlGlob, err = template.New("").Funcs(e.FuncMap).Delims(e.delims.Left, e.delims.Right).ParseGlob(e.Path)
	}

	if err != nil {
		log.Fatal(err)
	}
	return HtmlGlob
}
