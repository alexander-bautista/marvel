package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexander-bautista/marvel/pkg/comic"
	"github.com/gin-gonic/gin"
)

type handler struct {
	comicService comic.ComicService
}

func NewHandler(c comic.ComicService) http.Handler {
	h := &handler{c}

	router := gin.Default()

	router.GET("/", h.GetAll)
	router.GET("/:id", h.GetOne)
	router.GET("/:id/estimatedTaxes", h.EstimatedTaxes)

	return router
}

func (h *handler) GetOne(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	comic, err := h.comicService.GetOne(int(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find comic with id %s", idParam)})
	}

	c.JSON(http.StatusOK, comic)
}

func (h *handler) GetAll(c *gin.Context) {
	// Get query  parameters.
	dateRange := c.Query("dateRange")
	titleStartsWith := c.Query("titleStartsWith")

	if len(dateRange) == 0 && len(titleStartsWith) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Must provide a filter dateRange or titleStartsWith"})
		return
	}

	comics, err := h.comicService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comics)
}

func (h *handler) EstimatedTaxes(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	taxes, err := h.comicService.CalculateTaxes(int(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find item with id %s", idParam)})
	}

	c.JSON(http.StatusOK, taxes)
}
