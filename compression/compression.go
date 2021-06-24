// Package compression provides service that provides operations for
// image compression.
package compression

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"

	"github.com/vrubleuskii/image-compression/cache"
)

var (
	ErrCompression = errors.New("unable to compress the image")
	ErrNotFound    = errors.New("image not found")
)

type Service struct {
	cache *cache.Cache
}

func NewService(c *cache.Cache) *Service {
	return &Service{
		cache: c,
	}
}

// Compress compresses the image with imgName stored in working directory via jpeg algorithm
// using specified quality parameter.
// If an error occurs during compression, ErrCompression will be returned.
// If the image is not found, ErrNotFound will be returned.
func (s *Service) Compress(imgName string, quality int) (image.Image, error) {
	key := cache.Key{Name: imgName, Parameter: quality}

	if compressed := s.cache.Get(key); compressed != nil {
		return compressed, nil
	}

	f, err := os.Open(imgName)
	if err != nil {
		return nil, ErrNotFound
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
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
