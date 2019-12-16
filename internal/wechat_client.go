package internal

import (
	log "github.com/sirupsen/logrus"

	"github.com/parnurzeal/gorequest"
)

// WechatClient defines a wechat client
type WechatClient struct {
	Agent *gorequest.SuperAgent
	Host  string
	Token string
}

// NewWechatClient returns a inited wechat
func NewWechatClient(token string) *WechatClient {
	c := gorequest.New()
	return &WechatClient{
		Agent: c,
		Host:  "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=",
		Token: token,
	}
}

// SendToWechat send the corresponding content to wechat
func (c *WechatClient) SendToWechat(content string) {
	type markdown struct {
		Content string `json:"content"`
	}
	req := struct {
		MsgType  string   `json:"msgtype"`
		Markdown markdown `json:"markdown"`
	}{
		MsgType: "markdown",
		Markdown: markdown{
			Content: content,
		},
	}

	resp, body, errs := c.Agent.Post(c.Host + c.Token).Send(req).End()
	if errs != nil {
		log.Println(errs)
	}
	log.Println(resp.StatusCode, body)
}
