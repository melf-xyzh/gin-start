/**
 * @Time    :2022/4/17 16:17
 * @Author  :MELF晓宇
 * @Email   :xyzh.melf@petalmail.com
 * @FileName:aes.go
 * @Project :gin-start
 * @Blog    :https://blog.csdn.net/qq_29537269
 * @Guide   :https://guide.melf.space
 * @Information:
 *
 */

package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/melf-xyzh/gin-start/global"
)

// 高级加密标准（Advanced Encryption Standard，AES）

type AesCrypto struct{}

var AES = new(AesCrypto)

// EnPwdCode
/**
 * @Description: 加密base64
 * @param pwd
 * @return string
 * @return error
 */
func (aesCrypto *AesCrypto) EnPwdCode(pwd []byte) (string, error) {
	aesPwd := global.V.GetString("AES.Password")
	if aesPwd == "" {
		panic("AES.Password未初始化")
	}
	// pwdKey 加密钥匙，16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
	pwdKey := []byte(aesPwd)
	result, err := aesEncrypt(pwd, pwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

// DePwdCode
/**
 * @Description: 解密
 * @param pwd
 * @return []byte
 * @return error
 */
func (aesCrypto *AesCrypto) DePwdCode(pwd string) ([]byte, error) {
	aesPwd := global.V.GetString("AES.Password")
	if aesPwd == "" {
		panic("AES.Password未初始化")
	}
	// pwdKey 加密钥匙，16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
	pwdKey := []byte(aesPwd)
	//解密base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	//执行AES解密
	return aesDecrypt(pwdByte, pwdKey)
}

// pkcs7Padding
/**
 * @Description: PKCS7 填充模式
 * @param ciphertext
 * @param blockSize
 * @return []byte
 */
func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	// Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

// PKCS7UnPadding
/**
 * @Description:填充的反向操作，删除填充字符串
 * @param origData
 * @return []byte
 * @return error
 */
func pkcs7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充字符串长度
	unPadding := int(origData[length-1])
	//截取切片，删除填充字节，并且返回明文
	return origData[:(length - unPadding)], nil
}

// aesEncrypt
/**
 * @Description: 实现加密
 * @param origData
 * @param key
 * @return []byte
 * @return error
 */
func aesEncrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = pkcs7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// aesDecrypt
/**
 * @Description: 实现解密
 * @param cypted
 * @param key
 * @return []byte
 * @return error
 */
func aesDecrypt(cypted []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = pkcs7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}
