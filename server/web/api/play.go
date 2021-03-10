package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"server/log"
)

type playReqJS struct {
	Uri string
}

func play(c *gin.Context) {
	var req playReqJS
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.TLogln("try plau url:", req.Uri)
	cmd := exec.Command("./cmd/play", req.Uri)
	err = cmd.Run()
	if err!=nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(200)
}
