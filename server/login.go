package server

import (
	"log"
	"net/http"

	"github.com/daoraimi/devlab-backend/domain/value"
	"github.com/gin-gonic/gin"
)

func (s *Server) login(c *gin.Context) {
	var req value.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, value.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, value.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	resp, err := s.user.Login(c, req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, value.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
