package dictionary

import "testing"

func TestAdd(t *testing.T) {
	t.Run("saves value to key", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.items = make(map[Key]Value)

		key := "key1"
		value := "value1"
		dictionary.Add(key, value)

		if dictionary.items[key] != value {
			t.Errorf("value was not correctly saved: got %q want %q", dictionary.items[key], value)
		}
	})

	t.Run("saves value to key when map is not initialized", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "key1"
		value := "value1"
		dictionary.Add(key, value)

		if dictionary.items[key] != value {
			t.Errorf("value was not correctly saved: got %q want %q", dictionary.items[key], value)
		}
	})
}
