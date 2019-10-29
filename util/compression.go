/**
 * @Author: DollarKiller
 * @Description: str 压缩
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:30 2019-10-29
 */
package util

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

type StrZip struct {
}

func (s *StrZip) Zip(str string) string {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(str)); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	strc := base64.StdEncoding.EncodeToString(b.Bytes())
	return strc
}

func (s *StrZip) Unzip(str string) string {
	data, _ := base64.StdEncoding.DecodeString(str)
	rdata := bytes.NewReader(data)
	rc, _ := gzip.NewReader(rdata)
	all, _ := ioutil.ReadAll(rc)
	return string(all)
}
