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
			t.Errorf("value was not added: got %q want %q", dictionary.items[key], value)
		}
	})

	t.Run("saves value to key when map is not initialized", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "key1"
		value := "value1"
		dictionary.Add(key, value)

		if dictionary.items[key] != value {
			t.Errorf("value was not added: got %q want %q", dictionary.items[key], value)
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("removes a value", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "key1"
		value := "value1"
		dictionary.Add(key, value)

		if dictionary.items[key] != value {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key], value)
		}

		ok := dictionary.Remove(key)

		if !ok {
			t.Error("the value was not removed")
		}
	})
}
