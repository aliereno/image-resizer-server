package handlers

import (
	"fmt"
	"github.com/aliereno/image-resizer-server/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
	"net/http"
	"strings"
)

func UploadFileHandler(staticFilePath string) gin.HandlerFunc {
	return func(c *gin.Context) {

		file, err := c.FormFile("data")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf(err.Error()))
			return
		}
		oldFileName := strings.Split(file.Filename, ".")
		if len(oldFileName) != 2 {
			c.String(http.StatusBadRequest, fmt.Sprintf("filename error."))
			return
		}
		newFileName := fmt.Sprintf("%s.%s", shortuuid.New(), oldFileName[1])
		err = c.SaveUploadedFile(file, staticFilePath+"/"+newFileName)
		if err != nil {
			logger.Fatal(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"file_name": newFileName,
		})
		return
	}
}
