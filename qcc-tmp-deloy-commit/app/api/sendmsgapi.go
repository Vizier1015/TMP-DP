package api

import (
	"github.com/gogf/gf/net/ghttp"
	"qcc-tmp-deloy-commit/app/model"
	"qcc-tmp-deloy-commit/app/service"
	"qcc-tmp-deloy-commit/library/response"
)

var Sendmsg =  SendmsgApi{}
type SendmsgApi struct {

}


func (o SendmsgApi)Sendmsg(r *ghttp.Request)  {
	var (
		sendmsgReq *model.Statistics
	)
	if err := r.ParseForm(&sendmsgReq); err != nil{
		response.JsonExit(r,1,err.Error())
	}

	if err := service.TmpDp.DingSend(sendmsgReq);err != nil{
		response.JsonExit(r,1,err.Error())

	}else {
		response.JsonExit(r,0,"ok")
	}
}