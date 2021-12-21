package api

import (
	"os"

	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	var err error
	var filename string
	if filename = c.Param("filename"); filename == "" {
		FailWithMsg(c, "no filename provided")
		return
	}
	filename = "files/" + filename
	if _, err = os.Stat(filename); err != nil {
		FailWithMsg(c, err.Error())
		return
	}
	c.File(filename)
}
