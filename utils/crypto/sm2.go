package crypto

import (
	"crypto"
	"crypto/rand"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"log"
	"os"
	"path/filepath"
	"unsafe"
)

var SM2 = new(mySm2)

type mySm2 struct{}

// CreateSM2Key
/**
 *  @Description: 随机生成公私钥
 *  @return privateKey
 *  @return publicKey
 *  @return err
 */
func (s *mySm2) CreateSM2Key() (privateKey *sm2.PrivateKey, publicKey *sm2.PublicKey, err error) {
	// 生成sm2秘钥对
	privateKey, err = sm2.GenerateKey(rand.Reader)
	if err != nil {
		return
	}

	// 进行sm2公钥断言
	publicKey = privateKey.Public().(*sm2.PublicKey)
	return
}

// CreatePrivatePem
/**
 *  @Description: 创建私钥文件
 *  @param privateKey 私钥
 *  @param pwd 私钥密码
 *  @param path 生成的私钥文件路径
 *  @return err
 */
func (s *mySm2) CreatePrivatePem(privateKey *sm2.PrivateKey, pwd []byte, path string) (err error) {
	// 将私钥反序列化并进行pem编码
	var privateKeyToPem []byte
	privateKeyToPem, err = x509.WritePrivateKeyToPem(privateKey, pwd)
	if err != nil {
		return err
	}
	// 将私钥写入磁盘
	if path == "" {
		path = "cert/sm2Private.Pem"
	}
	// 获取文件中的路径
	paths, _ := filepath.Split(path)
	err = os.MkdirAll(paths, os.ModePerm)
	if err != nil {
		return err
	}
	var file *os.File
	file, err = os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(privateKeyToPem)
	if err != nil {
		return err
	}
	return
}

// CreatePublicPem
/**
 *  @Description: 创建公钥文件
 *  @param publicKey 公钥
 *  @param path 生成的公钥文件路径
 *  @return err
 */
func (s *mySm2) CreatePublicPem(publicKey *sm2.PublicKey, path string) (err error) {
	// 将私钥反序列化并进行pem编码
	var publicKeyToPem []byte
	publicKeyToPem, err = x509.WritePublicKeyToPem(publicKey)
	if err != nil {
		return err
	}
	// 将私钥写入磁盘
	if path == "" {
		path = "cert/sm2Public.Pem"
	}
	// 获取文件中的路径
	paths, _ := filepath.Split(path)
	err = os.MkdirAll(paths, os.ModePerm)
	if err != nil {
		return err
	}
	var file *os.File
	file, err = os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(publicKeyToPem)
	if err != nil {
		return err
	}
	return
}

// ReadPrivatePem
/**
 *  @Description: 读取私钥文件
 *  @param path 私钥文件路径
 *  @param pwd 私钥文件密码，无则填nil
 *  @return privateKey 私钥
 *  @return err
 */
func (s *mySm2) ReadPrivatePem(path string, pwd []byte) (privateKey *sm2.PrivateKey, err error) {
	// 打开文件读取私钥
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var fileInfo os.FileInfo
	fileInfo, err = file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, fileInfo.Size(), fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	// 将pem格式私钥文件进行反序列化
	privateKey, err = x509.ReadPrivateKeyFromPem(buf, pwd)
	if err != nil {
		return nil, err
	}
	return
}

// ReadPublicPem
/**
 *  @Description: 读取公钥文件
 *  @param path 公钥文件路径
 *  @return publicKey 公钥文件
 *  @return err
 */
func (s *mySm2) ReadPublicPem(path string) (publicKey *sm2.PublicKey, err error) {
	// 打开文件读取私钥
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var fileInfo os.FileInfo
	fileInfo, err = file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, fileInfo.Size(), fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	// 将pem格式私钥文件进行反序列化
	publicKey, err = x509.ReadPublicKeyFromPem(buf)
	if err != nil {
		return nil, err
	}
	return
}

// Encrypt
/**
 *  @Description: SM2加密（公钥加密）
 *  @param data 需要加密的数据
 *  @param publicKey 公钥
 *  @return cipherStr 加密后的字符串
 */
func (s *mySm2) Encrypt(data string, publicKey *sm2.PublicKey) (cipherStr string) {
	// 将字符串转为[]byte
	dataByte := []byte(data)
	// sm2加密
	cipherTxt, err := publicKey.EncryptAsn1(dataByte, rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	// 转为16进制字符串输出
	//cipherStr = fmt.Sprintf("%x", cipherTxt)
	cipherStr = hex.EncodeToString(cipherTxt)
	return
}

// Decode
/**
 *  @Description: 解密（私钥解密）
 *  @param cipherStr 加密后的字符串
 *  @param privateKey 私钥
 *  @return data 解密后的数据
 *  @return err
 */
func (s *mySm2) Decode(cipherStr string, privateKey *sm2.PrivateKey) (data string, err error) {
	// 16进制字符串转[]byte
	bytes, _ := hex.DecodeString(cipherStr)
	// sm2解密
	var dataByte []byte
	dataByte, err = privateKey.DecryptAsn1(bytes)
	if err != nil {
		return data, err
	}
	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&dataByte))
	return *str, err
}

// Sign
/**
 *  @Description: 签名
 *  @param msg 需要签名的内容
 *  @param privateKey 私钥
 *  @param signer
 *  @return sign
 *  @return err
 */
func (s *mySm2) Sign(msg string, privateKey *sm2.PrivateKey, signer crypto.SignerOpts) (sign string, err error) {
	if signer == nil {
		signer = crypto.SHA256
	}
	dataByte := []byte(msg)
	var signByte []byte
	// sm2签名
	signByte, err = privateKey.Sign(rand.Reader, dataByte, signer)
	if err != nil {
		return "", err
	}
	// 转为16进制字符串输出
	sign = hex.EncodeToString(signByte)
	return sign, nil
}

// Verify
/**
 *  @Description: 验签
 *  @param msg 需要验签的内容
 *  @param sign 验签
 *  @param publicKey 公钥
 *  @return verify
 */
func (s *mySm2) Verify(msg, sign string, publicKey *sm2.PublicKey) (verify bool) {
	// 16进制字符串转[]byte
	msgBytes := []byte(msg)
	signBytes, _ := hex.DecodeString(sign)
	// sm2验签
	verify = publicKey.Verify(msgBytes, signBytes)
	return
}
