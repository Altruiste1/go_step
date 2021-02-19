package utools

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

//全局变量
var privateKey, publicKey []byte

func init() {
	var err error
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
}

/**
 * @brief  获取RSA公钥长度
 * @param[in]       PubKey				    RSA公钥
 * @return   成功返回 RSA公钥长度，失败返回error	错误信息
 */
func GetPubKeyLen(PubKey []byte) (int, error) {
	if PubKey == nil {
		return 0, errors.New("input arguments error")
	}

	block, _ := pem.Decode(PubKey)
	if block == nil {
		return 0, errors.New("public rsaKey error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return 0, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return pub.N.BitLen(), nil
}

/**
 * @brief  获取RSA私钥长度
 * @param[in]       PriKey				    RSA私钥
 * @return   成功返回 RSA私钥长度，失败返回error	错误信息
 */
func GetPriKeyLen(PriKey []byte) (int, error) {
	if PriKey == nil {
		return 0, errors.New("input arguments error")
	}

	block, _ := pem.Decode(PriKey)
	if block == nil {
		return 0, errors.New("private rsaKey error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return 0, err
	}

	return priv.N.BitLen(), nil
}

//func main() {
//	//获取rsa 公钥长度
//	PubKeyLen, _ := GetPubKeyLen(publicKey)
//	fmt.Println("pbulic key len is ", PubKeyLen)
//
//	//获取rsa 私钥长度
//	PriKeyLen, _ := GetPriKeyLen(privateKey)
//	fmt.Println("private key len is ", PriKeyLen)
//}