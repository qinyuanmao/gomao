package main

import (
	"github.com/qinyuanmao/gomao/logger"
	"github.com/qinyuanmao/gomao/security"
	_ "github.com/qinyuanmao/gomao/security"
)

func main() {
	var info = "2021-01-02 13:33:33"
	resp, err := security.Encode(info)
	if err != nil {
		logger.Error(err)
	}
	logger.Info(resp)
	result, err := security.Decode(resp)
	if err != nil {
		logger.Error(err)
	}
	logger.Info(result)

	resp, err = security.EncodeToken(1, 1613345901)
	if err != nil {
		logger.Error(err)
	}
	logger.Info(resp)

	userID, expiredAt, err := security.DecodeToToken(resp)
	if err != nil {
		logger.Error(err)
	}
	logger.Info(userID, expiredAt)
}
