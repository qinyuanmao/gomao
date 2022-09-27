package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"sync"

	"e.coding.net/tssoft/repository/gomao/logger"
	"github.com/spf13/viper"
)

type encryptParser struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

var engine *encryptParser
var encryptOnce sync.Once

func GetEncryptEngine() *encryptParser {
	if engine == nil {
		encryptOnce.Do(func() {
			engine = newEngine()
		})
	}
	return engine
}

func newEngine() *encryptParser {
	privateData := viper.GetString("key.private")
	if privateData == "" {
		logger.Error("Private key is empty")
		return nil
	}
	publicData := viper.GetString("key.public")
	if privateData == "" {
		logger.Error("Public key is empty")
		return nil
	}
	engine = new(encryptParser)
	var err error
	engine.PrivateKey, err = privateKey(privateData)
	if err != nil {
		logger.Errorf("set private key failed: %s", err)
	}
	engine.PublicKey, err = publicKey(publicData)
	if err != nil {
		logger.Errorf("set public key failed: %s", err)
	}
	return engine
}

func publicKey(data string) (*rsa.PublicKey, error) {
	//解密 pem 格式的公钥
	block, _ := pem.Decode([]byte(data))
	if block == nil {
		return nil, errors.New("Public key error!")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	return pub, nil
}

func privateKey(data string) (*rsa.PrivateKey, error) {
	//解密
	block, _ := pem.Decode([]byte(data))
	if block == nil {
		return nil, errors.New("Private key error!")
	}
	//解析 PKCS1 格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func (p *encryptParser) Encode(input string) (output string, err error) {
	o, err := rsa.EncryptPKCS1v15(rand.Reader, p.PublicKey, []byte(input))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(o), nil
}

func (p *encryptParser) Decode(input string) (output string, err error) {
	i, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	o, err := rsa.DecryptPKCS1v15(rand.Reader, p.PrivateKey, i)
	return string(o), nil
}
