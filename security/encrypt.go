package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"os/exec"

	"e.coding.net/tssoft/repository/gomao/logger"
	"e.coding.net/tssoft/repository/gomao/utils"
	"github.com/spf13/viper"
)

type parser struct {
	publicKey  string
	privateKey string
}

var engine *parser

func init() {
	privateKeyPath := viper.GetViper().GetString("key.private")
	if privateKeyPath == "" {
		privateKeyPath = "./config/rsa_private_key.pem"
	}
	if !utils.FileExist(privateKeyPath) {
		err := exec.Command("/bin/bash", "-c", fmt.Sprintf("openssl genrsa -out %s 256", privateKeyPath)).Run()
		if err != nil {
			logger.Errorf("Create %s file error: %s", privateKeyPath, err)
		}
	}
	publicKeyPath := viper.GetViper().GetString("key.public")
	if publicKeyPath == "" {
		publicKeyPath = "./config/rsa_public_key.pem"
	}
	if !utils.FileExist(publicKeyPath) {
		err := exec.Command("/bin/bash", "-c", fmt.Sprintf("openssl rsa -in %s -pubout -out %s", privateKeyPath, publicKeyPath)).Run()
		if err != nil {
			logger.Errorf("Create %s file error: %s", publicKeyPath, err)
		}
	}
	engine = new(parser)
	engine.privateKey = utils.ReadFile(privateKeyPath)
	engine.publicKey = utils.ReadFile(publicKeyPath)
}

func Encode(input string) (output string, err error) {
	//解密 pem 格式的公钥
	block, _ := pem.Decode([]byte(engine.publicKey))
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	o, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(input))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(o), nil
}

func Decode(input string) (output string, err error) {
	//解密
	block, _ := pem.Decode([]byte(engine.privateKey))
	if block == nil {
		return "", fmt.Errorf("private key error")
	}
	//解析 PKCS1 格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	i, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	// 解密
	o, err := rsa.DecryptPKCS1v15(rand.Reader, priv, i)
	return string(o), err
}
