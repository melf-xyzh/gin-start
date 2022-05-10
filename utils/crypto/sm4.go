package crypto

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/tjfoc/gmsm/sm4"
)

var SM4 = new(mySm4)

type mySm4 struct{}

// byteXor
/**
 *  @Description: byte异或
 *  @param byte1
 *  @param byte2
 *  @return byteAnswer
 *  @return err
 */
func byteXor(byte1, byte2 []byte) (byteAnswer []byte, err error) {
	if len(byte1) != len(byte2) {
		err = errors.New("两个[]byte长度不相等")
		return
	}
	byteAnswer = make([]byte, len(byte1), len(byte1))
	for i := 0; i < len(byte1); i++ {
		byteAnswer[i] = byte1[i] ^ byte2[i]
	}
	return
}

// SecretText
/**
 *  @Description: 银联联机UtvtSm4Mac算法
 *  @param mab MAC ELEMEMENT BLOCK
 *  @param key 秘钥
 *  @return secretText 加密因子
 *  @return err
 */
func (s *mySm2) SecretText(mab string, key []byte) (secretText string, err error) {
	// 创建一个Sm4Cipher
	c, err := sm4.NewCipher(key)
	if err != nil {
		return
	}
	//// 由[硬件序列号+支付授权码后6位]构成MAC ELEMEMENT BLOCK （MAB）。
	//mab := sn + authCode[len(authCode)-6:]
	// SM4算法的MAB，按每16个字节做异或（不管信息中的字符格式），如果最后不满16个字节，则添加“0X00”。
	mabByte := []byte(mab)
	padding := 16 - len(mab)%16
	// Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	paddingText := bytes.Repeat([]byte{byte(0x00)}, padding)
	mabByte = append(mabByte, paddingText...)
	// 进行异或运算
	resultBlock := make([]byte, 16, 16)
	for i := 0; i < len(mabByte)/16-1; i++ {
		var byteXorErr error
		if i == 0 {
			resultBlock, byteXorErr = byteXor(mabByte[:16], mabByte[16:32])
		} else {
			resultBlock, byteXorErr = byteXor(mabByte[i*16:i*16+16], resultBlock)
		}
		if byteXorErr != nil {
			err = byteXorErr
			return
		}
	}
	// 将异或运算后的最后16个字节（RESULT BLOCK）转换成32 个HEXDECIMAL
	resultBlockStr := hex.EncodeToString(resultBlock)
	resultBlockByte := []byte(resultBlockStr)
	// 取前16个字节用SM4加密
	encBlock1 := make([]byte, 16)
	c.Encrypt(encBlock1, resultBlockByte[:16])
	// 将加密后的结果与后16个字节异或
	tempBlock, byteXorErr := byteXor(encBlock1, resultBlockByte[16:])
	if byteXorErr != nil {
		err = byteXorErr
		return
	}
	// 用异或的结果TEMP BLOCK 再进行一次SM4密钥算法运算。
	encBlock2 := make([]byte, 16)
	c.Encrypt(encBlock2, tempBlock)
	// 将运算后的结果（ENC BLOCK2）转换成32 个HEXDECIMAL
	// 取前8个字节作为硬件序列号加密数据
	//encResult := hex.EncodeToString(encBlock2)
	secretText = fmt.Sprintf("%X", encBlock2)[:8]
	return
}

// ReadPwdFromFile
/**
 *  @Description: 从文件读取sm4 key
 *  @receiver m
 *  @param path key.pem文件路径
 *  @param pwd 文件密码
 *  @return key sm4 key
 *  @return err
 */
func (m *mySm4) ReadPwdFromFile(path string, pwd []byte) (key []byte, err error) {
	key, err = sm4.ReadKeyFromPemFile(path, pwd)
	return
}

// name
/**
 *  @Description: 将sm4 key写入pem文件
 *  @receiver m
 *  @param path pem文件路径
 *  @param key sm4 key
 *  @param pwd 文件密码
 *  @return err
 */
func (m *mySm4) name(path string, key, pwd []byte) (err error) {
	err = sm4.WriteKeyToPemFile(path, key, pwd)
	return
}
