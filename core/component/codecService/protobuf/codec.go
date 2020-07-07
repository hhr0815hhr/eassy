package protobuf

import (
	"eassy/core/component/msgService"
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"hash/crc32"
	"reflect"
)

type ProtoCodec struct {
	littleEndian bool
}

func NewCodec() *ProtoCodec {
	p := new(ProtoCodec)
	p.littleEndian = false
	return p
}

func (p *ProtoCodec) SetEndian(isLittleEndian bool) {
	p.littleEndian = isLittleEndian
}

func (p *ProtoCodec) Encode(msg interface{}) ([]byte, error) {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return nil, errors.New("pb message pointer required")
	}
	msgID := msgType.Elem().Name()
	_id := crc32.ChecksumIEEE([]byte(msgID))

	id := make([]byte, 4)
	if p.littleEndian {
		binary.LittleEndian.PutUint32(id, _id)
	} else {
		binary.BigEndian.PutUint32(id, _id)
	}

	// data
	data, err := proto.Marshal(msg.(proto.Message))
	return data, err
}

func (p *ProtoCodec) Decode(route int, data []byte) (res interface{}, err error) {
	msg, ok := msgService.GetMsgService().GetMsgByRouteId(route)
	if !ok {
		return
	}
	res = reflect.New(msg.ReqType.Elem()).Interface()
	err = proto.UnmarshalMerge(data, res.(proto.Message))
	return
}
