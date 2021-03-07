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
	"regexp"
	"strconv"
	"time"

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
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(engine.publicKey))
	if block == nil {
		return "", errors.New("Public key error!")
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
		return "", errors.New("Private key error!")
	}
	//解析PKCS1格式的私钥
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
	return string(o), nil
}

func EncodeToken(userID int64, expiredAt int64) (token string, err error) {
	originValue := fmt.Sprintf("%d:%d", userID, expiredAt)
	return Encode(originValue)
}

func DecodeToToken(token string) (userID int64, expiredAt int64, err error) {
	output, err := Decode(token)
	if err != nil {
		return
	}
	r := regexp.MustCompile(`(\d+)\:(\d+)`)
	result := r.FindStringSubmatch(output)
	if len(result) < 3 {
		return 0, 0, fmt.Errorf("Token is incorrect!")
	}
	userID, _ = strconv.ParseInt(result[1], 10, 64)
	expiredAt, _ = strconv.ParseInt(result[2], 10, 64)
	if expiredAt < time.Now().Unix() {
		return 0, 0, fmt.Errorf("Token is expired!")
	}
	return
}
