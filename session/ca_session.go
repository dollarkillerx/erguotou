/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 18:48 2019-10-13
 */
package session

import (
	"errors"
	"github.com/bluele/gcache"
	"github.com/dollarkillerx/erguotou"
	"sync"
	"time"
)

// session node
type Session struct {
	Data           map[string]interface{} // 存储器
	mt             sync.Mutex
	CreationTime   int64 // 创建时间
	ExpirationTime int64 // 过期时间
}

func (s *Session) SetKey(key string, value interface{}) {
	s.mt.Lock()
	defer s.mt.Unlock()
	s.Data[key] = value
}

func (s *Session) GetVal(key string) (interface{}, bool) {
	s.mt.Lock()
	defer s.mt.Unlock()
	val, ok := s.Data[key]
	if ok {
		return val, true
	} else {
		return nil, false
	}
}

type SessionInterface interface {
	Get(id string) (*Session, error)                          // 拥有通过id获取session
	Set(data *Session) (string, error)                        // 返回id and 错误信息
	SetTime(data *Session, tim time.Duration) (string, error) // 返回id and 错误信息
	Del(id string) error
	Expired(id string) bool // 检测是否过期  过期返回false 反之true
}

// ================================不同数据源的实现=====================================
// 系统自带 gocache 存储
type GoSessionNode struct {
	gcache gcache.Cache
}

var (
	gache gcache.Cache
)

func init() {
	gache = gcache.New(20).
		LRU().
		Build()
}

func getCacheSession() SessionInterface {
	node := GoSessionNode{}
	return &node
}

func (g *GoSessionNode) Get(id string) (*Session, error) {

	// 存储器
	get, err := gache.Get(id)
	if err != nil {
		return nil, err
	}

	s, ok := get.(*Session)
	if ok {
		return s, nil
	}

	return nil, errors.New("data not ex")
}

func (g *GoSessionNode) Set(data *Session) (string, error) {

	// 生成随机key
	utils := erguotou.Utils{}
	key := utils.SuperRand()
	err := gache.SetWithExpire(key, data, 6*time.Hour)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (g *GoSessionNode) SetTime(data *Session, tim time.Duration) (string, error) {
	// 生成随机key
	utils := erguotou.Utils{}
	key := utils.SuperRand()
	err := gache.SetWithExpire(key, data, tim)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (g *GoSessionNode) Expired(id string) bool {
	_, e := gache.Get(id)
	if e != nil {
		return false
	} else {
		return true
	}
}

func (g *GoSessionNode) Del(id string) error {
	err := gache.SetWithExpire(id, "", time.Nanosecond)
	return err
}
