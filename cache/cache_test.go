package cache

import (
	"image"
	"testing"
)

func TestPut(t *testing.T) {
	const (
		size    = 32
		imgName = "bar.png"
	)

	t.Run("test filling of the cache", func(t *testing.T) {
		img := image.NRGBA{}
		cache := New(size)

		for i := 0; i < size; i++ {
			cache.Put(Key{Name: imgName, Parameter: i}, &img)
		}

		if cache.list.Len() != size {
			t.Errorf("expected list size to be equal to %d actual %d", size, cache.list.Len())
		}
	})

	t.Run("test ordering of elements in list", func(t *testing.T) {
		img := image.NRGBA{}
		cache := New(size)

		want := pair{key: Key{Name: imgName, Parameter: 0}, image: &img}

		for i := 0; i < size; i++ {
			cache.Put(Key{Name: imgName, Parameter: i}, &img)
		}
		got := cache.list.Back()

		if got.Value != want {
			t.Errorf("expected last element to be %v actual %v", want, got)
		}
	})
}
