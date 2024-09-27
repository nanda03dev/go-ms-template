package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go-ms-template/src/application/service"
)

type HealthController struct {
	healthService service.HealthService
}

func NewHealthController(healthService service.HealthService) *HealthController {
	return &HealthController{
		healthService: healthService,
	}
}

func (c *HealthController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, c.healthService.Health())
}
