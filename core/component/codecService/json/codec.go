package json

type JsonCodec struct {
}

func NewCodec() *JsonCodec {
	p := new(JsonCodec)
	return p
}

// goroutine safe
func (p *JsonCodec) Decode(route int, data []byte) (interface{}, error) {

	return nil, nil
}

// goroutine safe
func (p *JsonCodec) Encode(msg interface{}) ([]byte, error) {
	return nil, nil
}
