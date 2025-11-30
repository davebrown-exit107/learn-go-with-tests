package dictionary

func Search(search_term string, dictionary map[string]string) string {
	return dictionary[search_term]
}
