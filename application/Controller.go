package application

import (
	"net/http"
	"ws_amaris/domain/dto"
	"ws_amaris/domain/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.Service
}

func InitController(router *gin.Engine) {
	controller := Controller{
		service: service.InitServiceImpl(),
	}
	router.GET("/splitter/:string", controller.Splitter)
	router.GET("/pokemon/:id", controller.FindPokemonById)
	router.POST("/friendly-chains", controller.FriendlyChains)
}

func (r *Controller) Splitter(c *gin.Context) {

	split := c.Param("string")

	splitResponse := r.service.Splitter(split)

	c.JSON(http.StatusOK, splitResponse)
}

func (r *Controller) FindPokemonById(c *gin.Context) {

	id := c.Param("id")

	splitResponse, response := r.service.PokenById(id)

	if response != nil {
		c.JSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, splitResponse)
}

func (r *Controller) FriendlyChains(c *gin.Context) {

	var (
		friendly dto.FriendlyChainsDto
	)

	if err := c.ShouldBindJSON(&friendly); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.Response{
			Status:      http.StatusUnprocessableEntity,
			Description: "Invalid body",
		})
		return
	}

	message := r.service.FriendlyChains(friendly)

	c.JSON(http.StatusOK, message)
}
