package utils

// Map implements the high order `Map` function found in other programming languages
func Map(items []string, f func(string) string) []string {
	store := make([]string, len(items))

	for _, item := range items {
		store = append(store, f(item))
	}

	return store
}
