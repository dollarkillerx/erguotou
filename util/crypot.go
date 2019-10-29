/**
 * @Author: DollarKiller
 * @Description: 加密
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:19 2019-10-29
 */
package util

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

// 获取md5
func Md5Encode(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 获取sha1
func Sha1Encode(str string) string {
	data := []byte(str)
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

// 获取sha256
func Sha256Encode(str string) string {
	sum256 := sha256.Sum256([]byte(str))
	s := hex.EncodeToString(sum256[:])
	return s
}

// RSA256 公钥密钥对生成
// @params: bits 密钥长度
// @returns: private 密钥
// @returns: public 公钥
func GenRsaKey(bits int) (e error, priKey string, pubKey string) {

	// 生成私钥
	privateKey, e := rsa.GenerateKey(rand.Reader, bits)
	if e != nil {
		return e, "", ""
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RAS PRIVATE KEY",
		Bytes: derStream,
	}
	//fmt.Println("私密钥:",string(pem.EncodeToMemory(priBlock)))
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return e, "", ""
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	//fmt.Printf("=======公钥文件内容=========%v", string(pem.EncodeToMemory(publicBlock)))

	if err != nil {
		return e, "", ""
	}
	return nil, string(pem.EncodeToMemory(priBlock)), string(pem.EncodeToMemory(publicBlock))
}

// Rsa256 加密
// @params: origData 原始数据
// @Params: pubKey 公钥
func RsaEncrypt(origData, pubKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// Rsa256 加密简单
// @params: origData 原始数据
// @Params: pubKey 公钥
func RsaEncryptSimple(origData, pubKey string) (string, error) {
	orgData := []byte(origData)
	pubK := []byte(pubKey)
	bytes, e := RsaEncrypt(orgData, pubK)
	if e != nil {
		return "", e
	}
	encode := Base64Encode(bytes)
	return encode, nil
}

// Rsa256 解密
// @params: ciphertext 加密数据
// @Params: prvKey 私钥
func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// Rsa256 解密简单
func RsaDecryptSimple(ciphertext, privateKey string) (string, error) {
	decode, i := Base64Decode(ciphertext)
	if i != nil {
		return "", i
	}
	pri := []byte(privateKey)
	bytes, e := RsaDecrypt(decode, pri)
	return string(bytes), e
}

// Rsa256 签名
// @params: origData 需要签名的数据
// @Params: prvKey 私钥
func RsaSign(data, prvKey []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(data)
	sum := hash.Sum(nil)

	//解密私钥
	block, _ := pem.Decode(prvKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 签名
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, sum)
}

// Rsa256 验签
// @params: data 原始数据
// @params: signature 签名
// @params: publicKey 公钥
func RsaSignVer(data, signature, publicKey []byte) error {
	hashed := sha256.Sum256(data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

// Rsa256 签名简单
// @params: origData 需要签名的数据
// @Params: prvKey 私钥
func RsaSignSimple(data, prvKey string) (string, error) {
	bytes := []byte(data)
	i := []byte(prvKey)
	sign, e := RsaSign(bytes, i)
	if e != nil {
		return "", e
	}
	return Base64Encode(sign), nil
}

// Rsa256 验签简单
// @params: data 原始数据
// @params: signature 签名
// @params: publicKey 公钥
func RsaSignVerSimple(data, signature, publicKey string) error {
	dat := []byte(data)
	bytes, e := Base64Decode(signature)
	if e != nil {
		return e
	}
	pub := []byte(publicKey)
	return RsaSignVer(dat, bytes, pub)
}

// Base64编码
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64解码
func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// Base64URL编码
func Base64URLEncode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// Base64URL解码
func Base64URLDecode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}

// 对称加密 AES 高级标准加密
func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func AESEncode(key []byte, src []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func AESDecode(key []byte, src []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}
