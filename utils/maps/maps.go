package maps

// FirstEntry returns first entry as given by `range`
func FirstEntry[K comparable, V any](m map[K]V) (K, V) {
	for key, value := range m {
		return key, value
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
