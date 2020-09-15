package dictionary

import "sync"

// Key holds the key of a dictionary item.
// It is defined as an interface so it can be a generic type.
type Key interface{}

// Value holds the value of a dictionary item.
// It is defined as an interface so it can be a generic type.
type Value interface{}

// Dictionary is an object with key of type Key and value of type Value.
type Dictionary struct {
	items map[Key]Value
	lock  sync.RWMutex
}
