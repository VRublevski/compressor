// Package metrics provides service that processes requests
// for server metrics.
package metrics

import "github.com/vrubleuskii/image-compression/cache"

type Service struct {
	cache *cache.Cache
}

func NewService(c *cache.Cache) *Service {
	return &Service{
		cache: c,
	}
}

func (s *Service) CacheSize() int {
	return s.cache.Size()
}
