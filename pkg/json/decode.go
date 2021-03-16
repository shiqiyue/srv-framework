package json

// 反序列化
func Unmarshal(input []byte, v interface{}) error {
	return json.Unmarshal(input, v)
}
