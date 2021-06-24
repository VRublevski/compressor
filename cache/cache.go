// Package cache provides types and values for working with
// cache that stores image processing results.
package cache

import (
	"container/list"
	"image"
	"sync"
)

type Key struct {
	Name      string
	Parameter int
}

type Cache struct {
	list *list.List
	size int
	mx   sync.Mutex
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
	c.mx.Lock()
	defer c.mx.Unlock()
	var e *list.Element
	for e = c.list.Front(); e != nil; e = e.Next() {
		p := e.Value.(*pair)
		if p.key == key {
			break
		}
	}

	if e == nil {
		return nil
	}

	img := e.Value.(*pair).image
	c.list.MoveToFront(e)
	return img
}

// Put stores val under key in the cache replacing old entry if necessary.
func (c *Cache) Put(key Key, val image.Image) {
	c.mx.Lock()
	defer c.mx.Unlock()
	// check for presence of an element with the same key
	for e := c.list.Front(); e != nil; e = e.Next() {
		p := e.Value.(*pair)
		if p.key == key {
			c.list.Remove(e)
			break
		}
	}

	if c.list.Len() == c.size {
		c.list.Remove(c.list.Back())
	}

	p := pair{key: key, image: val}
	c.list.PushFront(&p)
}
