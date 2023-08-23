package repositories

import (
	"fmt"

	"github.com/diazharizky/teleforecaster/pkg/cache"
)

type cityRepository struct {
	cache *cache.Cache
}

const cityKeyFormat = "%s:states:%s"

var cityCache = map[string][]string{}

func NewCityRepository(cache *cache.Cache) cityRepository {
	return cityRepository{cache}
}

func (r cityRepository) List(country, state string) []string {
	key := fmt.Sprintf(cityKeyFormat, country, state)
	return cityCache[key]
}

func (r cityRepository) Save(country, state string, cities []string) {
	if len(cities) <= 0 {
		return
	}

	key := fmt.Sprintf(cityKeyFormat, country, state)
	cityCache[key] = cities
}
