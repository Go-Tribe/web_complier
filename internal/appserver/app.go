package appserver

import (
	"fmt"
	c "web_complier/configs"
	"web_complier/core"
	"web_complier/internal/appserver/router"
)

func NewApp(basename string) {
	core.InitCore()
	r := router.SetRouters()
	err := r.Run(fmt.Sprintf("%s:%d", c.Config.Server.Host, c.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}
