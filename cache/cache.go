// Package cache provides types and values for working with
// cache that stores image processing results.
package cache

import (
	"container/list"
	"image"
)

type Key struct {
	Name      string
	Parameter int
}

type Cache struct {
	list *list.List
	size int
}

type pair struct {
	key   Key
	image image.Image
}

// New creates a new Cache with capacity equal to size.
func New(size int) *Cache {
	return &Cache{
		list: list.New(),
		size: size,
	}
}

// Get returns the Image associated with the key, or nil if the key is not
// present.
func (c *Cache) Get(key Key) image.Image {
	return nil
}

// Put stores val under key in the cache replacing old entry if necessary.
func (c *Cache) Put(key Key, val image.Image) {
	if c.list.Len() == c.size {
		c.list.Remove(c.list.Back())
	}

	p := pair{key: key, image: val}
	c.list.PushFront(&p)
}
