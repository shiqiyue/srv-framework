package jsons

import "io"
import jsoniter "github.com/json-iterator/go"

func NewEncoder(w io.Writer) *jsoniter.Encoder {
	return jsoniter.NewEncoder(w)
}

func NewDecoder(r io.Reader) *jsoniter.Decoder {
	return jsoniter.NewDecoder(r)
}
