package weixin

import (
	"github.com/qinyuanmao/gomao/logger"
	"github.com/xlstudio/wxbizdatacrypt"
)

func DecryptWXOpenData(appID, sessionKey, encryptData, iv string) (string, error) {
	pc := wxbizdatacrypt.WxBizDataCrypt{AppId: appID, SessionKey: sessionKey}
	result, err := pc.Decrypt(encryptData, iv, true)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	} else {
		decodeStr := result.(string)
		return decodeStr, nil
	}
}
