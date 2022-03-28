package service

import (
	"errors"
	"fmt"
	"github.com/blinkbean/dingtalk"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"qcc-tmp-deloy-commit/app/dao"
	"qcc-tmp-deloy-commit/app/model"
	"strings"
)

// 发版通知
var TmpDp = TmpDeployMsg{}

type TmpDeployMsg struct {}

func GetPhone() []string {
	p2 := []string{}
	p1 := []string{"15618040081","15002130603"}
	for _,v := range p1{
		p2 = append(p2,v)
	}
	return p2
}

func (TmpDeployMsg)DingSend(r *model.Statistics)  error {
	if r.Project == "" {
		return errors.New(fmt.Sprintf("项目不能为空,请选择"))
	}
	if r.Category == "" {
		return errors.New(fmt.Sprintf("类别不为空"))
	}
	if r.Describe == "" {
		return errors.New(fmt.Sprintf("请填写发版描述"))
	}
	if r.Persion == "" {
		return errors.New(fmt.Sprintf("请填写申请人"))
	}
	if r.Taskid == ""{
		return errors.New(fmt.Sprintf("请填写taskid"))
	}
	formats := []string{
		"Y-m-d",
	}
	t := gtime.Now()
	r.Time = t.Format(strings.Join(formats,","))

	//r.Time = gtime.New(time.Now())
	if _,err := dao.Statistics.Save(r);err != err{
		return err
	}
	//p1 := []string{"15618040081","15002130603"}
	//for _,v := range p1{
	//	fmt.Println(v)
	//}
	//var p = GetPhone()
	//fmt.Println(p)
	p:="15618040081"
	var mobiles = make([]string,0,10)
	dm := dingtalk.DingMap()
	dm.Set("临时发布", dingtalk.H2)
	dm.Set("---",dingtalk.N)
	dm.Set("项目名称:"+r.Project,dingtalk.N)
	dm.Set("发版类别:"+r.Category,dingtalk.N)
	dm.Set("任务ID:"+r.Taskid,dingtalk.N)
	dm.Set("发版描述:"+r.Describe,dingtalk.N)
	dm.Set("发版时间:"+gconv.String(r.Time),dingtalk.N)
	dm.Set("申请人:" +gconv.String(r.Persion),dingtalk.N)

	//token := g.Cfg().GetString("")
	token := "46833f68e1a56604b62082b792a4d73c28c8c1daaf7fad712bb543ba454b8170"

	model.DingDing.DingSendMsg(token, "临时发布",dm.Slice(),append(mobiles,gconv.String(p)))

	return nil


}