package dingtalk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/qinyuanmao/gomao/logger"
)

type Client struct {
	webhook    string
	httpClient *http.Client
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

func NewClient(webhook string) *Client {
	c := &Client{
		webhook:    webhook,
		httpClient: &http.Client{Timeout: 30 * time.Second},
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
	jsonMsg, _ := json.Marshal(dingMsg)
	req, err := http.NewRequest("POST", c.webhook, strings.NewReader(string(jsonMsg)))
	if err != nil {
		logger.Errorf("send dingtalk message %v: %v", dingMsg, err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		logger.Errorf("send dingtalk message %v: %v", dingMsg, err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("request dingtalk webhook fail: %s, message: %v", resp.Status, dingMsg)
		return fmt.Errorf("request http error: %s", resp.Status)
	}

	return nil
}
