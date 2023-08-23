package repositories

import (
	"context"

	"github.com/diazharizky/teleforecaster/pkg/cache"
)

type stateRepository struct {
	cache *cache.Cache
}

func NewStateRepository(cache *cache.Cache) stateRepository {
	return stateRepository{cache}
}

func (r stateRepository) List(country string) ([]string, error) {
	key := country + ":states"
	res, err := r.cache.Client.LRange(context.TODO(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r stateRepository) Save(country string, states []string) error {
	key := country + ":states"
	return r.cache.Client.RPush(context.TODO(), key, states).Err()
}
