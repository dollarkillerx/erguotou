/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 18:57 2019-10-13
 */
package session

import (
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/erguotou/clog"
	"time"
)

type SessionEngine struct {
	db SessionInterface
	se *Session
}

var (
	session *SessionEngine
)

const sessionKey = "erguotou-session"

func GetSessionCache() *SessionEngine {
	if session == nil {
		return &SessionEngine{
			db: getCacheSession(),
			se: &Session{
				Data: map[string]interface{}{},
			},
		}
	} else {
		return session
	}
}

func (s *SessionEngine) Get(ctx *erguotou.Context, key string) (interface{}, bool) {
	cookie := ctx.GetCookie(sessionKey)
	if cookie == "" {
		return nil, false
	}
	sess, e := s.db.Get(cookie)
	if e != nil {
		clog.PrintEr(e)
		return nil, false
	}
	return sess.GetVal(key)
}

func (s *SessionEngine) Set(key string, val interface{}) {
	s.se.SetKey(key, val)
}

func (s *SessionEngine) Save(ctx *erguotou.Context) error {
	set, e := s.db.Set(s.se)
	if e != nil {
		return e
	}
	ctx.SetCookie(sessionKey, set)
	return nil
}

func (s *SessionEngine) SaveTime(ctx *erguotou.Context, time time.Duration) error {
	set, e := s.db.SetTime(s.se, time)
	if e != nil {
		return e
	}
	ctx.SetCookie(sessionKey, set)
	return nil
}

func (s *SessionEngine) Del (ctx *erguotou.Context,id string) error {
	err := s.db.Del(id)
	if err != nil {
		return err
	}
	ctx.SetCookie(sessionKey, "")
	return nil
}