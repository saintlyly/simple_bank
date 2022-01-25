package api

import (
	db "simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("api/accounts", server.createAccount)
	router.GET("api/accounts/:id", server.getAccount)
	router.GET("api/accounts", server.listAccount)
	router.POST("api/account/:id", server.updateAccount)

	server.router = router
	return server

}
