package weixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeixinPerson struct {
	OpenID     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	UnionID    string `json:"unionid"`     //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
	Errcode    int    `json:"errcode"`     //错误码
	ErrMsg     string `json:"errMsg"`      //错误信息
}

func Login(appID, appSecret, code string) (weixinPerson WeixinPerson, err error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	resp, err := http.Get(fmt.Sprintf(url, appID, appSecret, code))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &weixinPerson)
	if err != nil {
		return
	}
	if weixinPerson.Errcode != 0 {
		return weixinPerson, errors.New(fmt.Sprintf("code: %d, errmsg: %s", weixinPerson.Errcode, weixinPerson.ErrMsg))
	}
	return
}
