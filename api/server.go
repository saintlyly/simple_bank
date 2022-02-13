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
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
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
	router.POST("/users", server.createUser)
	router.POST("/user/login", server.loginUser)
	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRouter.POST("accounts", server.createAccount)
	authRouter.GET("accounts/:id", server.getAccount)
	authRouter.GET("accounts", server.listAccounts)

	authRouter.POST("transfers", server.createTransfer)
	server.router = router

}
