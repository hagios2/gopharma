package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hagios2/gopharma/util"
)

func main() {
	server := gin.Default()

	//setup db connection
	util.Connection()

	Routes(server)

	err := server.Run()
	if err != nil {
		return
	}
}
