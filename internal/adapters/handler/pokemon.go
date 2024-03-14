package handler

import (
	"net/http"
	"strconv"

	"github.com/KridtinC/pokedict/internal/adapters/handler/model"
	"github.com/KridtinC/pokedict/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type pokemonHandler struct {
	pokemonService ports.PokemonService
}

func NewPokemonHandler(pokemonService ports.PokemonService) *pokemonHandler {
	return &pokemonHandler{
		pokemonService,
	}
}

func (h *pokemonHandler) FindAll(c *gin.Context) {
	c.JSON(200, "ok")
}

func (h *pokemonHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	pkm, err := h.pokemonService.FindByID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(200, &model.FindByIDResponse{
		Data: model.PokemonFromDomain(pkm),
	})
}
