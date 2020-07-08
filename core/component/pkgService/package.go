package pkgService

const (
	PkgHeadBytes = 5 //包头字节数
)

/**
 * Package protocol encode.
 *
 * package format:
 * +-------+-------------+------------------+
 * | route | body length |       body       |
 * +-------+-------------+------------------+
 *
 * Head: 5bytes
 *   0-1: route id,
 *      0 - heartbeat,
 *   2 - 4: big-endian body length
 * Body: body length bytes
 */

func PkgEncode(route int, body []byte) []byte {
	length := len(body)
	head := make([]byte, PkgHeadBytes)
	head[0] = byte(route)
	head[1] = byte(route >> 8)
	head[2] = byte(length >> 16)
	head[3] = byte(length >> 8)
	head[4] = byte(length)
	return append(head, body...)
}

func PkgDecode(buffer []byte) (route int, body []byte) {
	if len(buffer) < PkgHeadBytes {
		return
	}
	route = GetPkgRoute(buffer)
	length := GetPkgBodyLen(buffer)
	body = make([]byte, length)
	copy(body, buffer[PkgHeadBytes:])
	return
}

func GetPkgRoute(buffer []byte) int {
	return int(buffer[1]<<8) | int(buffer[0])
}

func GetPkgBodyLen(buffer []byte) int {
	return int(uint32(buffer[4]) | uint32(buffer[3])<<8 | uint32(buffer[2])<<16)
}
