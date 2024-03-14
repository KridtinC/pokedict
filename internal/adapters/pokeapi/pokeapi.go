package pokeapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/KridtinC/pokedict/internal/adapters/pokeapi/model"
	"github.com/KridtinC/pokedict/internal/core/domain"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) GetPokemonByID(ctx context.Context, id int) (domain.Pokemon, error) {
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+strconv.Itoa(id), nil)
	if err != nil {
		return domain.Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return domain.Pokemon{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Pokemon{}, err
	}
	var pkmResp model.Pokemon
	if err := json.Unmarshal(body, &pkmResp); err != nil {
		return domain.Pokemon{}, err
	}
	return pkmResp.ToDomain(), nil
}
