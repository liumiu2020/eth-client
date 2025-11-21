package route

import "github.com/gin-gonic/gin"

var R *gin.Engine

func RouteReg() {
	R.POST("/bid", Bid)
}
