package handlers

import (
	"github.com/disintegration/imaging"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Ping is simple keep-alive/ping handler
func ServeResizedFiles(staticFilePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		width := c.Param("width")
		height := c.Param("height")
		filename := c.Param("filename")

		// string width and height to int
		nWidth, err := strconv.ParseInt(width, 10, 32)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		nHeight, err := strconv.ParseInt(height, 10, 32)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		path := staticFilePath + "/" + "resized" + "/" + width + "/" + height
		if _, err := os.Stat(path + "/" + filename); err == nil {
			// file already exist, so serve it
			file, err := os.Open(path + "/" + filename)
			if err != nil {
				c.Status(http.StatusNotFound)
			}
			splittedFilename := strings.Split(filename, ".")
			c.DataFromReader(http.StatusOK, -1, "image/"+splittedFilename[1], file, nil)
			_ = file.Close()

		} else {
			// file doesn't exist, so create resized file
			os.MkdirAll(path, 0700)
			basePath := staticFilePath + "/" + filename
			baseFile, err := os.Open(basePath)
			newFile, errr := os.OpenFile(path+"/"+filename, os.O_CREATE|os.O_WRONLY, 0644)
			defer baseFile.Close()
			defer newFile.Close()
			if err != nil {
				c.Status(http.StatusNotFound)
				return
			}
			if errr != nil {
				c.String(http.StatusServiceUnavailable, errr.Error())
				return
			}
			var contentLength int64 = -1
			img, err := imaging.Decode(baseFile)
			m := imaging.Resize(img, int(nWidth), int(nHeight), imaging.NearestNeighbor)
			format, _ := imaging.FormatFromFilename(filename)
			_ = imaging.Encode(newFile, m, format)
			contentType := "image/" + strings.ToLower(format.String())
			c.DataFromReader(http.StatusOK, contentLength, contentType, newFile, nil)
		}
	}
}
