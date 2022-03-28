package model

import "github.com/blinkbean/dingtalk"

var DingDing = DingTalk{}

type DingTalk struct {}

func (s DingTalk) DingSendMsg(token,title string, dm []string, mobiles []string)  {
	var dingToken = []string{token}
	cli := dingtalk.InitDingTalk(dingToken,".")
	cli.SendMarkDownMessageBySlice(title,dm,dingtalk.WithAtMobiles(mobiles))
}