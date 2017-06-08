package godingtalk

import (
	"net/url"
)

//SendRobotTextMessage can send a text message to a group chat
func (c *DingTalkClient) SendRobotTextMessage(
	accessToken string,
	msg string,
	atMobiles []string,
	atAll bool) error {
	var data OAPIResponse
	params := url.Values{}
	params.Add("access_token", accessToken)
	request := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"isAtAll":   atAll,
		},
	}
	err := c.httpRPC("robot/send", params, request, &data)
	return err
}
