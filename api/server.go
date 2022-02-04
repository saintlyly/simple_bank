package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	db "github.com/saintlyly/simple_bank/db/sqlc"
	"github.com/saintlyly/simple_bank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saintlyly/simple_bank/token"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)

	}

	server.setupRouter()
	return server, nil

}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("accounts", server.createAccount)
	router.GET("accounts/:id", server.getAccount)
	router.GET("accounts", server.listAccounts)
	router.POST("account/:id", server.updateAccount)

	router.POST("/users", server.createUser)
	router.POST("/user/login", server.loginUser)

	router.POST("transfers", server.createTransfer)
	server.router = router

}
