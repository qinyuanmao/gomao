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

type WeixinWebPerson struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"open_id"`
	Scope        string `json:"scope"`
	Errcode      int    `json:"errcode"` //错误码
	ErrMsg       string `json:"errMsg"`  //错误信息
}

func WebLogin(appID, appSecret, code, refreshToken string) (person WeixinWebPerson, err error) {
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	resp, err := http.Get(fmt.Sprintf(url, appID, appSecret, code))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &person)
	if err != nil {
		return
	}
	switch person.Errcode {
	case 0:
		return
	case 42001:
		return RefreshWebToken(appID, refreshToken)
	default:
		return person, fmt.Errorf("code: %d, errmsg: %s", person.Errcode, person.ErrMsg)
	}
}

func RefreshWebToken(appID, refreshToken string) (person WeixinWebPerson, err error) {
	url := "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	resp, err := http.Get(fmt.Sprintf(url, appID, refreshToken))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &person)
	if err != nil {
		return
	}
	switch person.Errcode {
	case 0:
		return
	default:
		return person, fmt.Errorf("code: %d, errmsg: %s", person.Errcode, person.ErrMsg)
	}
}

type WeixinWebUser struct {
	OpenID     string   `json:"open_id"`
	Nickname   string   `json:"nickname"`
	Sex        string   `json:"sex"`
	Province   string   `json::"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
	Errcode    int      `json:"errcode"` //错误码
	ErrMsg     string   `json:"errMsg"`  //错误信息
}

func GetWebUserInfo(openID, accessToken string) (user WeixinWebUser, err error) {
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	resp, err := http.Get(fmt.Sprintf(url, accessToken, openID))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		return
	}
	switch user.Errcode {
	case 0:
		return
	default:
		return user, fmt.Errorf("code: %d, errmsg: %s", user.Errcode, user.ErrMsg)
	}
}
