package repositories

import (
	"context"
	"fmt"

	"github.com/diazharizky/teleforecaster/pkg/cache"
)

type cityRepository struct {
	cache *cache.Cache
}

func NewCityRepository(cache *cache.Cache) cityRepository {
	return cityRepository{cache}
}

func (r cityRepository) List(country, state string) ([]string, error) {
	key := fmt.Sprintf("%s:states:%s", country, state)
	res, err := r.cache.Client.LRange(context.TODO(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r cityRepository) Save(country, state string, cities []string) error {
	key := fmt.Sprintf("%s:states:%s", country, state)
	return r.cache.Client.RPush(context.TODO(), key, cities).Err()
}
