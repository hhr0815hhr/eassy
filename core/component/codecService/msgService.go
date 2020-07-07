package codecService

import (
	"eassy/core/component/codecService/json"
	"eassy/core/component/codecService/protobuf"
	"sync"
)

var (
	msgType = MsgTypeProto
	once    sync.Once
	codec   IMsgCodec
)

func SetMsgType(t int) {
	msgType = t
}

type IMsgCodec interface {
	Encode(msg interface{}) ([]byte, error)
	Decode(route int, data []byte) (res interface{}, err error)
}

func GetCodecService() IMsgCodec {
	once.Do(func() {
		switch msgType {
		case MsgTypeJson:
			codec = json.NewCodec()
		case MsgTypeProto:
			codec = protobuf.NewCodec()
		}
	})
	return codec
}
