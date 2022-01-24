package handler

import (
	"cargo-rest-api/config"
	"cargo-rest-api/infrastructure/message/success"
	"cargo-rest-api/infrastructure/persistence"
	"cargo-rest-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler is a struct.
type PingHandler struct {
	cf *config.Config
}

// PingResponse is a struct.
type PingResponse struct {
	DB    string `json:"db"`
	Redis string `json:"redis"`
}

// NewPingHandler will initialize Ping handler.
func NewPingHandler(config *config.Config) *PingHandler {
	return &PingHandler{cf: config}
}

// @Summary Ping server
// @Description Check server response.
// @Tags development
// @Accept  json
// @Produce  json
// @Success 200 {object} response.successOutput
// @Failure 400 {string} response.errorOutput
// @Failure 404 {object} response.errorOutput
// @Failure 500 {object} response.errorOutput
// @Router /api/ping [get]
// Ping will handle ping request.
func (p *PingHandler) Ping(c *gin.Context) {
	var pingData PingResponse

	_, errDBConn := persistence.NewDBConnection(p.cf.DBConfig)
	if errDBConn != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errDBConn)
		pingData.DB = "Not OK"
	} else {
		pingData.DB = "OK"
	}

	redisConnection, errRedisConn := persistence.NewRedisConnection(p.cf.RedisConfig)
	if errRedisConn != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errRedisConn)
		pingData.Redis = "Not OK"
	} else {
		_ = redisConnection.Close()
		pingData.Redis = "OK"
	}

	response.NewSuccess(c, pingData, success.PingSuccessfull).JSON()
}
