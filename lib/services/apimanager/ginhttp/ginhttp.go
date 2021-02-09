package ginhttp

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinFamily struct {
	Engine *gin.Engine
}

func StartGinEngine() GinFamily {
	var ginFamily GinFamily

	r := gin.Default()

	ginFamily.Engine = r

	return ginFamily

}

func (g GinFamily) StartServer(port string) {
	g.Engine.Run(fmt.Sprintf(":%s", port))
}
