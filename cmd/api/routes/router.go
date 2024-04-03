package routes

import (
	"github.com/gin-gonic/gin"
	"pixel-pay/database"
	"pixel-pay/internal/controllers"
)

type Router struct {
	gin *gin.Engine
	db  database.Pgx
}

func NewRouter(gin *gin.Engine, db database.Pgx) *Router {
	return &Router{gin: gin, db: db}
}

func (r *Router) SetupRoutes() {
	r.gin.POST("/transaction", func(context *gin.Context) {
		controllers.PostTransaction(context, r.db)
	})
}
func (r *Router) Run(addr string) error {
	return r.gin.Run(addr)
}
