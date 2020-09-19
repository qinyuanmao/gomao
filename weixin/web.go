package weixin

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WeixinWebPerson struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	Errcode      int    `json:"errcode"` //错误码
	ErrMsg       string `json:"errMsg"`  //错误信息
}

func WebLogin(appID, appSecret, code string) (person WeixinWebPerson, err error) {
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
	return
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
	return
}

type WeixinWebUser struct {
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json::"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
	Errcode    int      `json:"errcode"` //错误码
	ErrMsg     string   `json:"errMsg"`  //错误信息
}

func (person *WeixinWebPerson) GetWebUserInfo() (user WeixinWebUser, err error) {
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	resp, err := http.Get(fmt.Sprintf(url, person.AccessToken, person.OpenID))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		return
	}
	return
}

type Ticket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
	Errcode   int    `json:"errcode"` //错误码
	ErrMsg    string `json:"errMsg"`  //错误信息
}

func (person *WeixinWebPerson) GetTicket() (ticket Ticket, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=wx_card"
	resp, err := http.Get(fmt.Sprintf(url, person.AccessToken))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ticket)
	if err != nil {
		return
	}
	return
}

func (t Ticket) GetSignature(nonceStr, url string) (timestamp int64, signature string) {
	timestamp = time.Now().Unix()
	signature = fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", t.Ticket, nonceStr, timestamp, url)
	h := sha1.New()
	h.Write([]byte(signature))
	bs := h.Sum(nil)
	signature = strings.ToLower(string(bs))
	return
}
