package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"qcc-tmp-deloy-commit/app/dao"
	"qcc-tmp-deloy-commit/app/model"
	"qcc-tmp-deloy-commit/library/response"
	"strings"
)


//请求写入数据库

var Acceptmsg = AccService{}
type AccService struct {}

func (s *AccService) Getmsgs(r *model.Statistics) error {
	if r.Project == "" {
		return errors.New(fmt.Sprintf("项目不能为空,请选择"))
	}
	if r.Category == "" {
		return errors.New(fmt.Sprintf("类别不为空"))
	}
	if r.Taskid == "" {
		return errors.New(fmt.Sprintf("任务ID不为空"))
	}
	if r.Describe == "" {
		return errors.New(fmt.Sprintf("请填写发版描述"))
	}
	if r.Persion == "" {
		return errors.New(fmt.Sprintf("请填写申请人"))
	}
	//r.Time = gtime.New(time.Now())
	formats := []string{
		"Y-m-d",
	}
	t := gtime.Now()
	r.Time = t.Format(strings.Join(formats,","))
	if _,err := dao.Statistics.Save(r);err != err{
		return err
	}
	return nil

}

/*
筛选统计项目相关信息：单个项目发布次数
				  每个项目具体信息
				  每个分类下项目发布次数
 				  每个分类下项目信息
 */

var Selectmsg = SelectService{}

type SelectService struct {}


//筛选指向项目所有信息

func (s SelectService) SelectProkect(project string) []byte {
	//定义空接口，接受任意类型
	//var ss = make([]map[string]interface{}, 0,100)
	var d = response.JsonResponse{}
	projectinfo := g.DB().Model("statistics").Safe()
	r, err := projectinfo.Where("project = ? ",project).All()
	if err != nil{
		d.Code = 1
		d.Message = "查询项目出错"
		data, _ := json.Marshal(d)
		return data
	}
	if r.IsEmpty(){
		d.Code =1
		d.Message = "没有查询到项目"
		data,_ := json.Marshal(d)
		return data
	}else {
		d.Code = 0
		d.Message = "查询到项目"
		d.Data = r
		data,_ :=  json.Marshal(d)
		return data
	}

}

//统计指定项目发布次数

func (s SelectService) ProjectSum(project string) int  {
	var d = response.JsonResponse{}
	sumproject := g.DB().Model("statistics").Safe()
	r,err := sumproject.Where("project = ?",project).Count()
	if err != nil{
		d.Code = 1
		d.Message = "查询项目失败"
		data := 0
		return data
	}else {
		d.Code = 0
		d.Message = "查询成功"
		data := r
		return data
	}
}


// 每个项目发布次数

func (s SelectService)ProjectandSum(project string)[]byte  {
	var d = response.JsonResponse{}
	ps := g.DB().Model("statistics").Safe()
	r,err := ps.Fields("COUNT(*) ,project").Group("project").All()
	if err != nil {
		d.Code = 1
		d.Message= "查询出错"
		data,_ := json.Marshal(d)
		return data
	}else {
		d.Code = 0
		d.Message = "success"
		d.Data = r
		data,_ := json.Marshal(d)
		return data
	}
}

//统计发版类别次数

func (s SelectService)CategroySum(categroy string) []byte  {
	var d = response.JsonResponse{}
	ss := g.DB().Model("statistics").Safe()
	r, err := ss.Fields("Count(*), category").Group("category").All()
	if err != nil {
		d.Code =1
		d.Message = "查询出错"
		data, _ := json.Marshal(d)
		return data
	}else {
		d.Code = 0
		d.Message ="success"
		d.Data = r
		data,_ := json.Marshal(d)
		return data
	}
}





