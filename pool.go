/**
 * @Author: DollarKiller
 * @Description: 对象池
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:02 2019-10-02
 */
package erguotou

import (
	"errors"
	"time"
)

// 对象池
type ObjPool struct {
	bufChan chan interface{}
}

type PoolGenerateItem func() interface{}

// 创建对象池
func NewObjPoll(obj PoolGenerateItem, num int) *ObjPool {
	pool := ObjPool{}
	pool.bufChan = make(chan interface{}, num)
	for i := 0; i < num; i++ {
		pool.bufChan <- obj()
	}

	return &pool
}

// 获取对象
func (p *ObjPool) GetObj(timeout time.Duration) (interface{}, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

// 放回对象
func (p *ObjPool) Release(obj interface{}) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("pool overflow")
	}
}
