package server

import (
	"Go-Rampup/config"
	"Go-Rampup/server/base"
	"fmt"
	"net/http"
)

var App *base.App
var Srv *http.Server

func StartServer() {
	conf := config.GetConfig()
	App := base.App{}
	App.Initialize()
	App.Server(fmt.Sprintf("%v:%v", conf.Server.IPAddress, conf.Server.Port))
}
