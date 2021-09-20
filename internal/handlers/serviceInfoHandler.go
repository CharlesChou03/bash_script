package handlers

import (
	"net/http"

	"github.com/CharlesChou03/_git/testmodule.git/config"
	"github.com/gin-gonic/gin"
)

// @Summary health checker API
// @Success 200 {string} string "ok"
// @Router /health [get]
func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// @Summary service version API
// @Success 200 {string} string "0.0.1"
// @Router /version [get]
func VersionHandler(c *gin.Context) {
	c.String(http.StatusOK, config.Env.Version)
}
