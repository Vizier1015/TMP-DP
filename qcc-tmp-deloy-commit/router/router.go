package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"qcc-tmp-deloy-commit/app/api"
	"qcc-tmp-deloy-commit/app/service"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.CORS,)
		group.ALL("/hello", api.Hello)
		group.ALL("/tmpdp", api.Acceptmsg)
		group.ALL("/sendmsg",api.Sendmsg)
		group.ALL("/getdatas", func(r *ghttp.Request) {
			r.Response.ServeFileDownload("/Users/vizier/Documents/1.txt")
		})

	})
}
