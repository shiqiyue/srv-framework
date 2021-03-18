package jsons

// 序列化
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
