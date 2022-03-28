package response


import "github.com/gogf/gf/net/ghttp"

//数据返回JSON数据结构

type JsonResponse struct {
	Code int	`json:"code"` // 错误码（0:success;1 false >1 错误码）
	Message string `json:"message"`
	Data interface{} `json:"data"` // 返回数据

}

// 标准返回数据结构封装

func Json(r *ghttp.Request,code int,message string, data ...interface{})  {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code: code,
		Message: message,
		Data: responseData,
	})
}

//返回json数据并退出当前http执行函数

func JsonExit(r *ghttp.Request,err int,msg string, data ...interface{})  {
	Json(r,err,msg,data...)
	r.Exit()
}