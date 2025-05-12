package handler

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(taskHandler TaskHandler) *Router {
	router := gin.Default()
	router.POST("/tasks", taskHandler.Post)
	return &Router{
		router,
	}
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
