package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

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
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//加密
	o, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(input))
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

	// 分块解密
	var decrypted []byte
	blockSize := priv.PublicKey.N.BitLen() / 8
	for len(i) > 0 {
		var blockToDecrypt []byte
		if len(i) > blockSize {
			blockToDecrypt = i[:blockSize]
			i = i[blockSize:]
		} else {
			blockToDecrypt = i
			i = nil
		}
		decryptedBlock, err := rsa.DecryptPKCS1v15(rand.Reader, priv, blockToDecrypt)
		if err != nil {
			return nil, err
		}
		decrypted = append(decrypted, decryptedBlock...)
	}

	return decrypted, nil
}
func (p *Parser) DecodePkcs8(input []byte) (output []byte, err error) {
	ciphertext, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed: %s", err)
	}

	//解密
	block, _ := pem.Decode([]byte(p.privateKey))
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析 PKCS8 格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 计算每个块的大小
	keySize := priv.(*rsa.PrivateKey).Public().(*rsa.PublicKey).N.BitLen() / 8
	var plaintext []byte

	// 分割数据并解密
	for len(ciphertext) > 0 {
		var chunk []byte
		if len(ciphertext) > keySize {
			chunk = ciphertext[:keySize]
			ciphertext = ciphertext[keySize:]
		} else {
			chunk = ciphertext
			ciphertext = nil
		}

		// 解密
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), chunk)
		if err != nil {
			return nil, err
		}

		plaintext = append(plaintext, decrypted...)
	}

	return plaintext, nil
}

func NewParser(publicKey, privateKey string) *Parser {
	return &Parser{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func NewDefaultParser() *Parser {
	viper.SetDefault("key.private", "./config/rsa_private_key.pem")
	viper.SetDefault("key.public", "./config/rsa_public_key.pem")
	privateKeyPath := viper.GetViper().GetString("key.private")
	publicKeyPath := viper.GetViper().GetString("key.public")
	if !utils.FileExist(privateKeyPath) || !utils.FileExist(publicKeyPath) {
		createKey(publicKeyPath, privateKeyPath)
	}
	privateKey := utils.ReadFile(privateKeyPath)
	publicKey := utils.ReadFile(publicKeyPath)
	return NewParser(string(publicKey), string(privateKey))
}

func createKey(publicKeyPath, privateKeyPath string) {
	if !utils.FileExist("./config") {
		os.Mkdir("./config", os.ModePerm)
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logger.Panicf("create key error: %v", err)
	}

	utils.CreateFile(privateKeyPath, 0700, func(file *os.File) {
		err := pem.Encode(file, &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		})
		if err != nil {
			logger.Panicf("encode private key error: %v", err)
		}
	})

	utils.CreateFile(publicKeyPath, 0755, func(file *os.File) {
		pub := key.Public()
		err := pem.Encode(file, &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		})
		if err != nil {
			logger.Panicf("encode public key error: %v", err)
		}
	})
}
