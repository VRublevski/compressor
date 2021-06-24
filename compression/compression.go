// Package compression provides service that provides operations for
// image compression.
package compression

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"

	"github.com/vrubleuskii/image-compression/cache"
)

var (
	ErrCompression = errors.New("unable to compress the image")
)

type Service struct {
	cache *cache.Cache
}

func NewService(c *cache.Cache) *Service {
	return &Service{
		cache: c,
	}
}

// Compress compresses the img via jpeg algorithm using specified quality.
// If an error occurs, ErrCompression will be returned.
func (s *Service) Compress(imgName string, img image.Image, quality int) (image.Image, error) {
	key := cache.Key{Name: imgName, Parameter: quality}

	if compressed := s.cache.Get(key); compressed != nil {
		return compressed, nil
	}

	compressed, err := compress(img, quality)
	if err != nil {
		return nil, ErrCompression
	}

	s.cache.Put(key, compressed)
	return compressed, nil
}

func compress(img image.Image, quality int) (image.Image, error) {
	options := jpeg.Options{Quality: quality}

	var b bytes.Buffer
	if err := jpeg.Encode(&b, img, &options); err != nil {
		return nil, err
	}

	return jpeg.Decode(&b)
}
