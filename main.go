package main

import (
	_ "noticeservices/boot"
	_ "noticeservices/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
