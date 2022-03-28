package api

import (
	"github.com/gogf/gf/net/ghttp"
	"qcc-tmp-deloy-commit/library/response"

	"qcc-tmp-deloy-commit/app/model"
	"qcc-tmp-deloy-commit/app/service"
)

var Acceptmsg = AccmsgApi{}

type AccmsgApi struct {}

// @summary

func (a *AccmsgApi) Getmsg(r *ghttp.Request)  {
	var (
		msgReq *model.Statistics
	)
	if err := r.ParseForm(&msgReq);err != nil{
		response.JsonExit(r,1,err.Error())
	}

	if err := service.Acceptmsg.Getmsgs(msgReq);err != nil{
		response.JsonExit(r,1,err.Error())
	}else {
		response.JsonExit(r,0,"OK")
	}
}

//项目查询

func (a AccmsgApi ) Searchproject(r *ghttp.Request)  {
	project := r.Get("project")
	i, _:= project.(string)
	info := service.Selectmsg.SelectProkect(i)
	r.Response.Writeln(info)
}

//项目数量查询

func (a AccmsgApi)  Sumproject(r *ghttp.Request){
	sumproject := r.Get("project")
	i, _:= sumproject.(string)
	info := service.Selectmsg.ProjectSum(i)
	r.Response.Writeln(info)
}


// 每个项目发布次数

func (a AccmsgApi) Everyps(r *ghttp.Request)  {
	ps := r.Get("project")
	i,_ := ps.(string)
	info := service.Selectmsg.ProjectandSum(i)
	r.Response.Writeln(info)
}
// 统计每个分类发布次数

func (a AccmsgApi) Everyss(r *ghttp.Request)  {
	ss := r.Get("category")
	i,_:= ss.(string)
	info := service.Selectmsg.CategroySum(i)
	r.Response.Writeln(info)
}