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

type Parser struct {
	publicKey  string
	privateKey string
}

func (p *Parser) Encode(input []byte) (output []byte, err error) {
	//解密 pem 格式的公钥
	block, _ := pem.Decode([]byte(p.publicKey))
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
	o, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(input))
	if err != nil {
		return nil, err
	}
	return []byte(hex.EncodeToString(o)), nil
}

func (p *Parser) Decode(input []byte) (output []byte, err error) {
	//解密
	block, _ := pem.Decode([]byte(p.privateKey))
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析 PKCS1 格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	i, err := hex.DecodeString(string(input))
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, i)
}

func NewParser(publicKey, privateKey string) *Parser {
	return &Parser{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func NewDefaultParser() *Parser {
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
	privateKey := utils.ReadFile(privateKeyPath)
	publicKey := utils.ReadFile(publicKeyPath)
	return NewParser(string(publicKey), string(privateKey))
}
