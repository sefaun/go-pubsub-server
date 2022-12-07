package utils

func Contains(source []string, element string) bool {
	set := make(map[string]struct{}, len(source))
	for _, s := range source {
		set[s] = struct{}{}
	}

	_, ok := set[element]
	return ok
}
