package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	srv = gin.Default()
)

func main() {
	port := os.Getenv("PORT")
	s := &http.Server{
		Addr:    ":" + port,
		Handler: srv,
	}
	srv.NoRoute(unknown)
	srv.NoMethod(unknown)
	srv.POST("/login", login)
	srv.POST("/register", register)
	srv.GET("/check", check)
	if err := s.ListenAndServe(); err != nil {
		log.Println("failed to start server")
	}
}

func login(ctx *gin.Context) {}

func register(ctx *gin.Context) {}

func check(ctx *gin.Context) {}

func unknown(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, &Error{
		Code:    http.StatusMethodNotAllowed,
		Message: "invalid request",
	})
}

type User struct {
	ID string `json:"id"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
