package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexander-bautista/marvel/pkg/character"
	"github.com/alexander-bautista/marvel/pkg/comic"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type handler struct {
	comicService     comic.ComicService
	characterService character.CharacterService
}

func NewHandler(c comic.ComicService, ch character.CharacterService) http.Handler {
	h := &handler{c, ch}

	router := gin.Default()

	v1 := router.Group("/api/comics")
	{
		v1.GET("/", h.GetAllComics)
		v1.GET("/:id", h.GetOneComic)
		v1.GET("/:id/estimatedTaxes", h.ComicEstimatedTaxes)
	}

	v2 := router.Group("/api/characters")
	{
		v2.GET("/", h.GetAllCharacters)
		v2.GET("/:id", h.GetOneCharacter)
		v2.GET("/:id/scream", h.CharacterScream)
		v2.POST("/", h.CharacterAdd)
	}

	return router
}

func (h *handler) GetOneComic(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	result, err := h.comicService.GetOne(int(id))

	if err != nil {
		if errors.Cause(err) == comic.ErrComicNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find comic with id %s", idParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *handler) GetAllComics(c *gin.Context) {
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

func (h *handler) ComicEstimatedTaxes(c *gin.Context) {
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

func (h *handler) GetOneCharacter(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	result, err := h.characterService.GetOne(int(id))

	if err != nil {
		if errors.Cause(err) == comic.ErrComicNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find comic with id %s", idParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *handler) GetAllCharacters(c *gin.Context) {
	chars, err := h.characterService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, chars)
}

func (h *handler) CharacterScream(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%s is not a valid parameter", idParam)})
		return
	}

	// Get query  parameters.
	what := c.Query("what")

	if len(what) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Must provide a what parameter"})
		return
	}

	scream, err := h.characterService.Scream(int(id), what)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Cannot find item with id %s", idParam)})
	}

	c.JSON(http.StatusOK, gin.H{"message": scream})
}

func (h *handler) CharacterAdd(c *gin.Context) {
	var char character.Character

	err := c.BindJSON(&char)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "body mismatch"})
		return
	}

	result, err := h.characterService.Add(&char)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
