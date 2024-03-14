package main

import (
	"github.com/KridtinC/pokedict/config"
	"github.com/KridtinC/pokedict/infrastructure"
	"github.com/KridtinC/pokedict/internal/adapters/handler"
	"github.com/KridtinC/pokedict/internal/adapters/pokeapi"
	"github.com/KridtinC/pokedict/internal/adapters/repository"
	"github.com/KridtinC/pokedict/internal/core/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config.New()
	if config.Get().AppConfig.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	db := infrastructure.NewMongoDB()

	pokeAPI := pokeapi.NewClient()

	txRepository := repository.NewTxRepository(db.Client())
	pokemonRepository := repository.NewPokemonRepository(db)
	abilityRepository := repository.NewAbilityRepository(db)
	pokemonService := services.NewPokemonService(txRepository, pokemonRepository, abilityRepository, pokeAPI)

	pokemonHandler := handler.NewPokemonHandler(pokemonService)

	r := gin.Default()
	r.Use(handler.LoggerMiddleware())
	r.Handle("GET", "/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	pkm := r.Group("/pokemon")
	pkm.GET("", pokemonHandler.FindAll)
	pkm.GET("/:id", pokemonHandler.FindByID)

	r.Run(":" + config.Get().AppConfig.Port)
}
