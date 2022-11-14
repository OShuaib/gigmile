package controller

import (
	"github.com/gin-gonic/gin"
)

func (h *NewHttp) Routes(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.Use(h.Serve())
	//mux.HandleFunc("/", Handle)
	r.GET("/countries", h.GetCountries())
	r.GET("/countries/:id", h.GetCountry())
	r.PATCH("/countries/:id", h.UpdateCountry())
	r.POST("/country", h.CreateCountry())

}
