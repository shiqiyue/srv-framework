package jsons

// εεΊεε
func Unmarshal(input []byte, v interface{}) error {
	return json.Unmarshal(input, v)
}
