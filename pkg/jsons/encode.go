package jsons

// εΊεε
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
