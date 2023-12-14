package main

import (
	"todolist/config"
	"todolist/db/dao"
	"todolist/pkg/util"
	"todolist/routes"
)

func main() {
	loading()
	r := routes.Newrouter()
	_ = r.Run(config.HttpPort)
}
func loading() {
	config.Init()
	util.InitLog()
	dao.MYSQLInit()
}
