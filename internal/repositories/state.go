package repositories

import (
	"fmt"

	"github.com/diazharizky/teleforecaster/pkg/cache"
)

type stateRepository struct {
	cache *cache.Cache
}

const stateKeyFormat = "%s:states"

var stateCache = map[string][]string{}

func NewStateRepository(cache *cache.Cache) stateRepository {
	return stateRepository{cache}
}

func (r stateRepository) List(country string) []string {
	key := fmt.Sprintf(stateKeyFormat, country)
	return stateCache[key]
}

func (r stateRepository) Save(country string, states []string) {
	if len(states) <= 0 {
		return
	}

	key := fmt.Sprintf(stateKeyFormat, country)
	stateCache[key] = states
}
