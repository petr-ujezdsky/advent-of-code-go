package maps

// FirstEntry returns first entry as given by `range`
func FirstEntry[K comparable, V any](m map[K]V) (K, V) {
	for key, value := range m {
		return key, value
	}
	panic("Empty set")
}

// FirstKey returns first key as given by `range`
func FirstKey[K comparable, V any](m map[K]V) K {
	for key := range m {
		return key
	}
	panic("Empty set")
}

// FirstValue returns first value as given by `range`
func FirstValue[K comparable, V any](m map[K]V) V {
	for _, value := range m {
		return value
	}
	panic("Empty set")
}

// Values returns slice of map values
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0

	for _, value := range m {
		values[i] = value
		i++
	}

	return values
}

// Keys returns slice of map keys
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0

	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}

// Copy returns shallow copy of the map
func Copy[K comparable, V any](source map[K]V) map[K]V {
	dest := make(map[K]V, len(source))

	for key, value := range source {
		dest[key] = value
	}

	return dest
}

// Intersection finds intersection of keys
func Intersection[K comparable, V any](maps []map[K]V) map[K]V {
	intersection := Copy(maps[0])

	for _, m := range maps[1:] {
		for key := range intersection {
			if _, ok := m[key]; !ok {
				delete(intersection, key)
			}

			if len(intersection) == 0 {
				return intersection
			}
		}
	}

	return intersection
}
