package controllers

import (
	"net/http"

	"github.com/myomyintko/gin_gorm_mux/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
