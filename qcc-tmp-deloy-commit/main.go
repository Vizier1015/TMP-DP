package main

import (
	_ "qcc-tmp-deloy-commit/boot"
	_ "qcc-tmp-deloy-commit/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
