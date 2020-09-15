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

func TestGet(t *testing.T) {
	t.Run("gets a value", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "key1"
		value := "value1"
		dictionary.Add(key, value)

		if dictionary.items[key] != value {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key], value)
		}

		got := dictionary.Get(key)

		if got != value {
			t.Errorf("the wrong value was returned: got %q want %q", got, value)
		}
	})
}

func TestExists(t *testing.T) {
	t.Run("returns true if key exists", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "key1"
		value := "value1"
		dictionary.Add(key, value)

		if dictionary.items[key] != value {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key], value)
		}

		if !dictionary.Exists(key) {
			t.Error("key was reported as nonexistent")
		}
	})

	t.Run("returns false if key does not exists", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "key1"
		value := "value1"
		dictionary.Add(key, value)

		if dictionary.items[key] != value {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key], value)
		}

		if dictionary.Exists("butts") {
			t.Error("key was reported as existent")
		}
	})
}

func TestClear(t *testing.T) {
	t.Run("clears the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "key1"
		value := "value1"
		dictionary.Add(key, value)

		if dictionary.items[key] != value {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key], value)
		}

		dictionary.Clear()

		if dictionary.Exists(key) {
			t.Error("the dictionary was not cleared")
		}
	})
}

func TestLength(t *testing.T) {
	t.Run("returns the number of items", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "key1"
		value := "value1"
		dictionary.Add(key, value)

		if dictionary.items[key] != value {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key], value)
		}

		if dictionary.Length() != 1 {
			t.Errorf("length was not reported correctly: got %d want %d", dictionary.Length(), 1)
		}
	})
}

func TestKeys(t *testing.T) {
	t.Run("returns the dictionary's keys", func(t *testing.T) {
		dictionary := Dictionary{}

		key1 := "key1"
		value1 := "value1"
		dictionary.Add(key1, value1)

		if dictionary.items[key1] != value1 {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key1], value1)
		}

		key2 := "key2"
		value2 := "value2"
		dictionary.Add(key2, value2)

		if dictionary.items[key2] != value2 {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key2], value2)
		}

		keys := dictionary.Keys()

		if keys[0] != key1 {
			t.Errorf("key does not match: got %q want %q", keys[0], key1)
		}

		if keys[1] != key2 {
			t.Errorf("key does not match: got %q want %q", keys[1], key2)
		}
	})
}

func TestValues(t *testing.T) {
	t.Run("returns the dictionary's values", func(t *testing.T) {
		dictionary := Dictionary{}

		key1 := "key1"
		value1 := "value1"
		dictionary.Add(key1, value1)

		if dictionary.items[key1] != value1 {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key1], value1)
		}

		key2 := "key2"
		value2 := "value2"
		dictionary.Add(key2, value2)

		if dictionary.items[key2] != value2 {
			t.Fatalf("value was not added: got %q want %q", dictionary.items[key2], value2)
		}

		values := dictionary.Values()

		if values[0] != value1 {
			t.Errorf("key does not match: got %q want %q", values[0], value1)
		}

		if values[1] != value2 {
			t.Errorf("key does not match: got %q want %q", values[1], value2)
		}
	})
}
