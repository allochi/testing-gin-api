package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	SetupAPI(r)
	r.Run(":9600")
}

func SetupAPI(r *gin.Engine) {
	r.GET("/accounts", GetAccounts)
}
