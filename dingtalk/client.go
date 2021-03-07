package dingtalk

import (
	"fmt"
	"sync"

	"e.coding.net/tssoft/repository/gomao/logger"
	"e.coding.net/tssoft/repository/gomao/network"
	"github.com/spf13/viper"
)

type Client struct {
	webhook    string
	httpClient *network.Client
	notifyChan chan *DingTalkMsg
}

type DingTalkMsg struct {
	MsgType  string   `json:"msgtype"`
	Text     Text     `json:"text,omitempty"`
	Markdown Markdown `json:"markdown,omitempty"`
	At       At       `json:"at,omitempty"`
}

type Text struct {
	Content string `json:"content"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

var instance *Client
var once sync.Once

func GetInstance() *Client {
	once.Do(func() {
		instance = newClient(viper.GetString("dingding.webhook"))
	})
	return instance
}

func newClient(webhook string) *Client {
	c := &Client{
		webhook:    webhook,
		httpClient: network.NewClient(),
		notifyChan: make(chan *DingTalkMsg, 100),
	}
	go c.consume()
	logger.Infof("dingtalk  webhook: %s", webhook)
	return c
}

func (c *Client) Close() {
	close(c.notifyChan)
}

func (c *Client) Notify(dingMsg *DingTalkMsg) error {
	select {
	case c.notifyChan <- dingMsg:
		return nil
	default:
		return fmt.Errorf("too many errors to send")
	}
}

func (c *Client) BlockNotify(dingMsg *DingTalkMsg) error {
	return c.send(dingMsg)
}

func (c *Client) consume() {
	for dingMsg := range c.notifyChan {
		c.send(dingMsg)
	}
}

func (c *Client) send(dingMsg *DingTalkMsg) error {
	req := c.httpClient.Post(viper.GetString("dingtalk.webhook"))
	req.SetHeader("Content-Type", "application/json")
	req.JSON(dingMsg)
	var resultMap = make(map[string]interface{})
	err := req.ToJSON(&resultMap)

	if err != nil {
		logger.Errorf("send dingtalk message %v: %v", dingMsg, err)
		return err
	}
	if _, exist := resultMap["errmsg"]; exist {
		logger.Errorf("send dingtalk message failed with code %d, message: %s", resultMap["errcode"], resultMap["errmsg"])
		return err
	}

	return nil
}
