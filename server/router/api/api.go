package api

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	api := r.Group("/api")

	RegisterUser(api)

	RegisterOrders(api)

	RegisterPolls(api)

}
