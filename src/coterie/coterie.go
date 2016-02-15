package coterie

import (
	"encoding/binary"
	"errors"
	"net"

	"github.com/golang/protobuf/proto"
)

func WriteCoterieMsg(coterieMsg *CoterieMsg, conn net.Conn) error {
	bytes, err := proto.Marshal(coterieMsg)
	if err != nil {
		return err
	}
	lengthBytes := make([]byte, 8)
	binary.PutUvarint(lengthBytes, uint64(len(bytes)))

	if _, err := conn.Write(lengthBytes); err != nil {
		return err
	}

	if _, err := conn.Write(bytes); err != nil {
		return err
	}

	return nil
}

func ReadCoterieMsg(conn net.Conn) (*CoterieMsg, error) {
	buf := make([]byte, 4096)
	_, err := conn.Read(buf[:8])
	if err != nil {
		return nil, err
	}

	length, bytesRead := binary.Uvarint(buf[:8])
	if bytesRead < 0 {
		return nil, errors.New("Unable to parse length of dht protobuf")
	}

	_, err = conn.Read(buf[:length])
	if err != nil {
		return nil, err
	}

	dhtMsg := new(CoterieMsg)
	err = proto.Unmarshal(buf[:length], dhtMsg)
	if err != nil {
		return nil, err
	}

	return dhtMsg, nil
}
