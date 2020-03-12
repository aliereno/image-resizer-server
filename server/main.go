package main

import (
	"github.com/aliereno/image-resizer-server/internal/handlers"
	log "github.com/aliereno/image-resizer-server/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var host, port, projectPath string

func init() {
	host = "localhost"
	port = "7777"
	projectPath = "/home/alieren/Projects/Go/image-resizer-server"
}

func main() {
	//gin.SetMode(gin.ReleaseMode)

	// Logging to a file.
	f, _ := os.OpenFile(projectPath+"/server/gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	e, _ := os.OpenFile(projectPath+"/server/gin-error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(e, os.Stdout)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("files/:width/:height/:filename", handlers.ServeResizedFiles(projectPath+"/files"))
	r.POST("/upload-file", handlers.UploadFileHandler(projectPath+"/files"))
	r.GET("/ping", handlers.Ping())

	log.Fatal(r.Run(host + ":" + port))
}
