package api

import (
	"github.com/go-playground/validator/v10"
	db "github.com/saintlyly/simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)

	}

	router.POST("accounts", server.createAccount)
	router.GET("accounts/:id", server.getAccount)
	router.GET("accounts", server.listAccount)
	router.POST("account/:id", server.updateAccount)

	router.POST("transfers", server.createTransfer)

	server.router = router
	return server

}
