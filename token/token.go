/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:08 2019-10-29
 */
package token

import (
	"errors"
	"github.com/dollarkillerx/erguotou"
	"github.com/dollarkillerx/erguotou/cache"
	"github.com/dollarkillerx/erguotou/util"
	"time"
)

var (
	Token *token
)

// token在cache中的标识符
const tokenKey = "erguotou-token-ccm01"

type token struct {
	key string `json:"salt"`
}

func init() {
	Token = &token{
		key: util.SuperRand()[:32],
	}
}

type JwtHeader struct {
	Alg  string `json:"alg"`  // 算法名称
	Type string `json:"type"` // 类型
}

type JwtPayload struct {
	User    string        `json:"user"`  // 用户 名
	Email   string        `json:"email"` // 用户邮件
	Data    interface{}   `json:"data"`  // 需要存的数据
	TimeOut time.Duration `json:"-"`
	Iss     string        `json:"-"` // 用户唯一标识
}

func NewJwt() *JwtPayload {
	return &JwtPayload{
		TimeOut: time.Hour * 12,
	}
}

// 生成token
func (t *token) GeneraJwtToken(payload *JwtPayload) (string, error) {
	header := &JwtHeader{
		Alg:  "ase128",
		Type: "Rijndael",
	}
	head, e := erguotou.Jsonp.Marshal(header)
	if e != nil {
		return "", e
	}
	body, e := erguotou.Jsonp.Marshal(payload)
	if e != nil {
		return "", e
	}

	head64 := util.Base64URLEncode(head)
	body64 := util.Base64URLEncode(body)

	footer := util.Base64URLEncode(util.AESEncode([]byte(t.key), []byte(head64+body64)))

	token := head64 + "." + body64 + "." + footer

	// 存储到内存
	cacheKey := t.getCacheKey(token)
	e = cache.SetWithExpire(cacheKey, payload, payload.TimeOut)
	return token, e
}

func (t *token) getCacheKey(token string) string {
	return util.Md5Encode(token + tokenKey)
}

// 验证token
func (t *token) CheckToken(token string) (*JwtPayload, error) {
	key := t.getCacheKey(token)
	// 先cache中获取
	get, bool := cache.GetEx(key)
	if !bool {
		return nil, errors.New("401")
	}

	payload, ok := get.(*JwtPayload)
	if !ok {
		return nil, errors.New("Data type conversion failed")
	}

	return payload, nil
}
